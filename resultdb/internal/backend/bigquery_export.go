// Copyright 2019 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package backend

import (
	"context"
	"net/http"
	"time"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/spanner"
	"github.com/golang/protobuf/proto"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/sync/parallel"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/caching"

	"go.chromium.org/luci/resultdb/internal"
	"go.chromium.org/luci/resultdb/internal/span"
	pb "go.chromium.org/luci/resultdb/proto/rpc/v1"
)

const (
	maxInvocationGraphSize = 1000
	maxBatchRowCount       = 500

	// HTTP request size limit is 10 MiB according to
	// https://cloud.google.com/bigquery/quotas#streaming_inserts
	// Use a smaller size as the limit since we are only using the size of
	// test results to estimate the whole payload size.
	maxBatchSize = 6000000
)

var bqTableCache = caching.RegisterLRUCache(50)

// inserter is implemented by bigquery.Inserter.
type inserter interface {
	// Put uploads one or more rows to the BigQuery service.
	Put(ctx context.Context, src interface{}) error
}

type table interface {
	// Metadata fetches the metadata for the table.
	Metadata(ctx context.Context) (md *bigquery.TableMetadata, err error)
}

func getLUCIProject(ctx context.Context, invID span.InvocationID) (string, error) {
	realm, err := span.ReadInvocationRealm(ctx, span.Client(ctx).Single(), invID)
	if err != nil {
		return "", err
	}

	project, _, err := internal.ParseRealm(realm)
	if err != nil {
		return "", errors.Annotate(err, "invocation %q", invID.Name()).Err()
	}
	return project, nil
}

func getBQClient(ctx context.Context, luciProject string, bqExport *pb.BigQueryExport) (*bigquery.Client, error) {
	tr, err := auth.GetRPCTransport(ctx, auth.AsProject, auth.WithProject(luciProject), auth.WithScopes(bigquery.Scope))
	if err != nil {
		return nil, err
	}

	return bigquery.NewClient(ctx, bqExport.Project, option.WithHTTPClient(&http.Client{
		Transport: tr,
	}))
}

// checkBqTable returns true if the table should be created.
func checkBqTable(ctx context.Context, bqExport *pb.BigQueryExport, t table) (bool, error) {
	type cacheKey struct {
		project string
		dataset string
		table   string
	}

	shouldCreateTable := false
	key := cacheKey{
		project: bqExport.Project,
		dataset: bqExport.Dataset,
		table:   bqExport.Table,
	}

	v, err := bqTableCache.LRU(ctx).GetOrCreate(ctx, key, func() (interface{}, time.Duration, error) {
		_, err := t.Metadata(ctx)
		apiErr, ok := err.(*googleapi.Error)
		switch {
		case ok && apiErr.Code == http.StatusNotFound:
			// Table doesn't exist, no need to cache this.
			shouldCreateTable = true
			return nil, 0, err
		case ok && apiErr.Code == http.StatusForbidden:
			// No read table permission.
			return permanentInvocationTaskErrTag.Apply(err), time.Minute, nil
		case err != nil:
			// The err is not a special case above, no need to cache this.
			return nil, 0, err
		default:
			// Table exists and is accessible.
			return nil, 5 * time.Minute, nil
		}
	})

	switch {
	case shouldCreateTable:
		return true, nil
	case err != nil:
		return false, err
	case v != nil:
		return false, v.(error)
	default:
		return false, nil
	}
}

// ensureBQTable creates a BQ table if it doesn't exist.
func ensureBQTable(ctx context.Context, client *bigquery.Client, bqExport *pb.BigQueryExport) error {
	t := client.Dataset(bqExport.Dataset).Table(bqExport.Table)

	// Check the existence of table.
	switch shouldCreateTable, err := checkBqTable(ctx, bqExport, t); {
	case err != nil:
		return err
	case !shouldCreateTable:
		return nil
	}

	// Table doesn't exist. Create one.
	schema, err := bigquery.InferSchema(&TestResultRow{})
	if err != nil {
		panic(err)
	}
	err = t.Create(ctx, &bigquery.TableMetadata{
		Schema: schema,
	})
	apiErr, ok := err.(*googleapi.Error)
	switch {
	case err == nil:
		logging.Infof(ctx, "Created BigQuery table %s.%s.%s", bqExport.Project, bqExport.Dataset, bqExport.Table)
		return nil
	case ok && apiErr.Code == http.StatusConflict:
		// Table just got created. This is fine.
		return nil
	case ok && apiErr.Code == http.StatusForbidden:
		// No create table permission.
		return permanentInvocationTaskErrTag.Apply(err)
	default:
		return err
	}
}

type testVariantKey struct {
	testID      string
	variantHash string
}

// queryExoneratedTestVariants reads exonerated test variants matching the predicate.
func queryExoneratedTestVariants(ctx context.Context, txn *spanner.ReadOnlyTransaction, invIDs span.InvocationIDSet) (map[testVariantKey]struct{}, error) {
	st := spanner.NewStatement(`
		SELECT DISTINCT TestId, VariantHash,
		FROM TestExonerations
		WHERE InvocationId IN UNNEST(@invIDs)
	`)
	st.Params["invIDs"] = invIDs

	tvs := map[testVariantKey]struct{}{}
	var b span.Buffer
	err := span.Query(ctx, txn, st, func(row *spanner.Row) error {
		var key testVariantKey
		if err := b.FromSpanner(row, &key.testID, &key.variantHash); err != nil {
			return err
		}
		tvs[key] = struct{}{}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return tvs, nil
}

func queryTestResultsStreaming(
	ctx context.Context,
	txn *spanner.ReadOnlyTransaction,
	exportedID span.InvocationID,
	q span.TestResultQuery,
	exoneratedTestVariants map[testVariantKey]struct{},
	maxBatchRowCount int,
	maxBatchSize int,
	batchC chan []*bigquery.StructSaver) error {
	invs, err := span.ReadInvocationsFull(ctx, txn, q.InvocationIDs)
	if err != nil {
		return err
	}

	rows := make([]*bigquery.StructSaver, 0, maxBatchRowCount)
	batchSize := 0 // Estimated size of rows in bytes.
	err = span.QueryTestResultsStreaming(ctx, txn, q, func(tr *pb.TestResult, variantHash string) error {
		_, exonerated := exoneratedTestVariants[testVariantKey{testID: tr.TestId, variantHash: variantHash}]
		parentID, _, _ := span.MustParseTestResultName(tr.Name)
		rows = append(rows, generateBQRow(invs[exportedID], invs[parentID], tr, exonerated))
		batchSize += proto.Size(tr)
		if len(rows) >= maxBatchRowCount || batchSize >= maxBatchSize {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case batchC <- rows:
			}
			rows = make([]*bigquery.StructSaver, 0, maxBatchRowCount)
		}
		return nil
	})

	if err != nil {
		return err
	}

	if len(rows) > 0 {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case batchC <- rows:
		}
	}

	return nil
}

func batchExportRows(ctx context.Context, ins inserter, batchC chan []*bigquery.StructSaver) error {
	return parallel.WorkPool(10, func(workC chan<- func() error) {
		for rows := range batchC {
			rows := rows
			workC <- func() error {
				err := ins.Put(ctx, rows)
				if apiErr, ok := err.(*googleapi.Error); ok && apiErr.Code == http.StatusForbidden {
					err = permanentInvocationTaskErrTag.Apply(err)
				}
				return err
			}
		}
	})
}

// exportTestResultsToBigQuery queries test results in Spanner then exports them to BigQuery.
func exportTestResultsToBigQuery(ctx context.Context, ins inserter, invID span.InvocationID, bqExport *pb.BigQueryExport, maxBatchRowCount int, maxBatchSize int) error {
	txn := span.Client(ctx).ReadOnlyTransaction()
	defer txn.Close()

	inv, err := span.ReadInvocationFull(ctx, txn, invID)
	if err != nil {
		return err
	}
	if inv.State != pb.Invocation_FINALIZED {
		return errors.Reason("%s is not finalized yet", invID.Name()).Err()
	}

	// Get the invocation set.
	invIDs, err := span.ReadReachableInvocations(ctx, txn, maxInvocationGraphSize, span.NewInvocationIDSet(invID))
	if err != nil {
		if span.TooManyInvocationsTag.In(err) {
			err = permanentInvocationTaskErrTag.Apply(err)
		}
		return err
	}

	exoneratedTestVariants, err := queryExoneratedTestVariants(ctx, txn, invIDs)
	if err != nil {
		return err
	}

	// Query test results and export to BigQuery.
	batchC := make(chan []*bigquery.StructSaver)

	// Batch exports rows to BigQuery.
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return batchExportRows(ctx, ins, batchC)
	})

	q := span.TestResultQuery{
		Predicate:         bqExport.GetTestResults().GetPredicate(),
		InvocationIDs:     invIDs,
		SelectVariantHash: true,
	}
	eg.Go(func() error {
		defer close(batchC)
		return queryTestResultsStreaming(ctx, txn, invID, q, exoneratedTestVariants, maxBatchRowCount, maxBatchSize, batchC)
	})

	return eg.Wait()
}

// exportResultsToBigQuery exports results of an invocation to a BigQuery table.
func exportResultsToBigQuery(ctx context.Context, invID span.InvocationID, payload []byte) error {
	bqExport := &pb.BigQueryExport{}
	if err := proto.Unmarshal(payload, bqExport); err != nil {
		return err
	}

	luciProject, err := getLUCIProject(ctx, invID)
	if err != nil {
		return err
	}

	client, err := getBQClient(ctx, luciProject, bqExport)
	if err != nil {
		return err
	}
	defer client.Close()

	if err := ensureBQTable(ctx, client, bqExport); err != nil {
		return err
	}

	ins := client.Dataset(bqExport.Dataset).Table(bqExport.Table).Inserter()
	return exportTestResultsToBigQuery(ctx, ins, invID, bqExport, maxBatchRowCount, maxBatchSize)
}
