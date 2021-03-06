// Copyright 2020 The LUCI Authors.
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
	"cloud.google.com/go/bigquery"

	"go.chromium.org/luci/resultdb/internal/span"
	"go.chromium.org/luci/resultdb/pbutil"
	pb "go.chromium.org/luci/resultdb/proto/rpc/v1"
	typepb "go.chromium.org/luci/resultdb/proto/type"
)

// StringPair is a copy of typepb.StringPair, suitable for representing a
// key:value pair in a BQ table.
// Inferred to be a field of type RECORD with Key and Value string fields.
type StringPair struct {
	Key   string `bigquery:"key"`
	Value string `bigquery:"value"`
}

// Invocation is a subset of pb.Invocation for the invocation fields that need
// to be saved in a BQ table.
type Invocation struct {
	// ID is the ID of the invocation.
	ID string `bigquery:"id"`

	// Interrupted is a flag indicating whether the invocation is interrupted or not.
	// For more details, refer to pb.Invocation.Interrupted.
	Interrupted bool `bigquery:"interrupted"`

	// Tags represents Invocation-level string key-value pairs.
	// A key can be repeated.
	Tags []StringPair `bigquery:"tags"`
}

// TestResultRow represents a row in a BigQuery table for result of a functional
// test case.
type TestResultRow struct {
	// ExportedInvocation contains info of the exported invocation.
	// Note that it's possible that this invocation is not the result's
	// immediate parent invocation, but the including invocation.
	ExportedInvocation Invocation `bigquery:"exported"`

	// ParentInvocation contains info of the result's immediate parent
	// invocation.
	ParentInvocation Invocation `bigquery:"parent"`

	// TestID is a unique identifier of the test in a LUCI project.
	// Refer to pb.TestResult.TestId for details.
	TestID string `bigquery:"test_id"`

	// ResultID identifies a test result in a given invocation and test id.
	ResultID string `bigquery:"result_id"`

	// Variant describes one specific way of running the test,
	//  e.g. a specific bucket, builder and a test suite.
	Variant []StringPair `bigquery:"variant"`

	// Expected is a flag indicating whether the result of test case execution is expected.
	// Refer to pb.TestResult.Expected for details.
	Expected bool `bigquery:"expected"`

	// Status of the test result.
	Status string `bigquery:"status"`

	// SummaryHTML is a human-readable explanation of the result, in HTML.
	SummaryHTML string `bigquery:"summary_html"`

	// StartTime is the point in time when the test case started to execute.
	StartTime bigquery.NullTimestamp `bigquery:"start_time"`

	// Duration of the test case execution in seconds.
	Duration bigquery.NullFloat64 `bigquery:"duration"`

	// Tags contains metadata for this test result.
	// It might describe this particular execution or the test case.
	Tags []StringPair `bigquery:"tags"`

	// If the failures of the test variant are exonerated.
	// Note: the exoneration is at the test variant level, not result level.
	Exonerated bool `bigquery:"exonerated"`
}

// stringPairProtosToStringPairs returns a slice of StringPair derived from *typepb.StringPair.
func stringPairProtosToStringPairs(pairs []*typepb.StringPair) []StringPair {
	if len(pairs) == 0 {
		return nil
	}

	sp := make([]StringPair, len(pairs))
	for i, p := range pairs {
		sp[i] = StringPair{
			Key:   p.Key,
			Value: p.Value,
		}
	}
	return sp
}

// variantToStringPairs returns a slice of StringPair derived from *typepb.Variant.
func variantToStringPairs(vr *typepb.Variant) []StringPair {
	defMap := vr.GetDef()
	if len(defMap) == 0 {
		return nil
	}

	keys := pbutil.SortedVariantKeys(vr)
	sp := make([]StringPair, len(keys))
	for i, k := range keys {
		sp[i] = StringPair{
			Key:   k,
			Value: defMap[k],
		}
	}
	return sp
}

func invocationProtoToInvocation(inv *pb.Invocation) Invocation {
	return Invocation{
		ID:          string(span.MustParseInvocationName(inv.Name)),
		Interrupted: inv.Interrupted,
		Tags:        stringPairProtosToStringPairs(inv.Tags),
	}
}

// generateBQRow returns a *bigquery.StructSaver to be inserted into BQ.
func generateBQRow(exported, parent *pb.Invocation, tr *pb.TestResult, exonerated bool) *bigquery.StructSaver {
	trr := &TestResultRow{
		ExportedInvocation: invocationProtoToInvocation(exported),
		ParentInvocation:   invocationProtoToInvocation(parent),
		TestID:             tr.TestId,
		ResultID:           tr.ResultId,
		Variant:            variantToStringPairs(tr.Variant),
		Expected:           tr.Expected,
		Status:             tr.Status.String(),
		SummaryHTML:        tr.SummaryHtml,
		Tags:               stringPairProtosToStringPairs(tr.Tags),
		Exonerated:         exonerated,
	}

	if tr.StartTime != nil {
		trr.StartTime = bigquery.NullTimestamp{
			Timestamp: pbutil.MustTimestamp(tr.StartTime),
			Valid: true,
		}
	}

	if tr.Duration != nil {
		trr.Duration = bigquery.NullFloat64{
			Float64: pbutil.MustDuration(tr.Duration).Seconds(),
			Valid: true,
		}
	}

	return &bigquery.StructSaver{
		InsertID: tr.Name,
		Struct:   trr,
	}
}
