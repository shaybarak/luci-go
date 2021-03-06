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

package span

import (
	"bytes"
	"context"
	"reflect"
	"sort"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	tspb "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/klauspost/compress/zstd"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"

	"go.chromium.org/luci/common/errors"

	internalpb "go.chromium.org/luci/resultdb/internal/proto"
	"go.chromium.org/luci/resultdb/pbutil"
	pb "go.chromium.org/luci/resultdb/proto/rpc/v1"
	typepb "go.chromium.org/luci/resultdb/proto/type"
)

// This file implements utility functions that make spanner API slightly easier
// to use.

// Txn is implemented by all spanner transactions.
type Txn interface {
	// ReadRow reads a single row from the database.
	ReadRow(ctx context.Context, table string, key spanner.Key, columns []string) (*spanner.Row, error)

	// Read reads multiple rows from the database.
	Read(ctx context.Context, table string, key spanner.KeySet, columns []string) *spanner.RowIterator

	// ReadOptions reads multiple rows from the database, and allows customizing
	// options.
	ReadWithOptions(ctx context.Context, table string, keys spanner.KeySet, columns []string, opts *spanner.ReadOptions) (ri *spanner.RowIterator)

	// Query reads multiple rows returned by SQL statement.
	Query(ctx context.Context, statement spanner.Statement) *spanner.RowIterator
}

// Compressed instructs ToSpanner and FromSpanner functions to compress the
// content with https://godoc.org/github.com/klauspost/compress/zstd encoding.
type Compressed []byte

// CompressedProto is like Compressed, but for protobuf messages.
//
// Although not strictly necessary, FromSpanner expects a *CompressedProto,
// for consistency.
//
// Note that `type CompressedProto proto.Message` and
// `type CompressedProto interface{proto.Message}` do not work well with
// type-switches.
type CompressedProto struct{ Message proto.Message }

// ErrNoResults is an error returned when a query unexpectedly has no results.
var ErrNoResults = iterator.Done
var zstdHeader = []byte("ztd\n")

// Globally shared zstd encoder and decoder. We use only their EncodeAll and
// DecodeAll methods which are allowed to be used concurrently. Internally, both
// the encode and the decoder have worker pools (limited by GOMAXPROCS) and each
// concurrent EncodeAll/DecodeAll call temporary consumes one worker (so overall
// we do not run more jobs that we have cores for).
var (
	zstdEncoder *zstd.Encoder
	zstdDecoder *zstd.Decoder
)

func init() {
	var err error
	if zstdEncoder, err = zstd.NewWriter(nil); err != nil {
		panic(err) // this is impossible
	}
	if zstdDecoder, err = zstd.NewReader(nil); err != nil {
		panic(err) // this is impossible
	}
}

func slices(m map[string]interface{}) (keys []string, values []interface{}) {
	keys = make([]string, 0, len(m))
	values = make([]interface{}, 0, len(m))
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	return
}

// ReadRow reads a single row from the database and reads its values.
// ptrMap must map from column names to pointers where the values will be
// written.
func ReadRow(ctx context.Context, txn Txn, table string, key spanner.Key, ptrMap map[string]interface{}) error {
	columns, ptrs := slices(ptrMap)
	row, err := txn.ReadRow(ctx, table, key, columns)
	if err != nil {
		return err
	}

	return FromSpanner(row, ptrs...)
}

// Buffer can convert a value from a Spanner type to a Go type.
// Supported types:
//   - string
//   - InvocationID
//   - InvocationIDSet
//   - tspb.Timestamp
//   - pb.InvocationState
//   - pb.TestStatus
//   - pb.Artifact
//   - typepb.Variant
//   - typepb.StringPair
//   - Compressed
//   - CompressedProto
//   - proto.Message
type Buffer struct {
	nullStr               spanner.NullString
	nullTime              spanner.NullTime
	i64                   int64
	strSlice              []string
	byteSlice, byteSlice2 []byte
}

// FromSpanner is a shortcut for (&Buffer{}).FromSpanner.
// Appropriate when FromSpanner is called only once, whereas Buffer is reusable
// throughout function.
func FromSpanner(row *spanner.Row, ptrs ...interface{}) error {
	return (&Buffer{}).FromSpanner(row, ptrs...)
}

// FromSpanner reads values from row to ptrs, converting types from Spanner
// to Go along the way.
func (b *Buffer) FromSpanner(row *spanner.Row, ptrs ...interface{}) error {
	if len(ptrs) != row.Size() {
		panic("len(ptrs) != row.Size()")
	}

	for i, goPtr := range ptrs {
		if err := b.fromSpanner(row, i, goPtr); err != nil {
			return err
		}
	}
	return nil
}

func (b *Buffer) fromSpanner(row *spanner.Row, col int, goPtr interface{}) error {
	b.strSlice = b.strSlice[:0]
	b.byteSlice = b.byteSlice[:0]

	var spanPtr interface{}
	switch goPtr.(type) {
	case *string:
		spanPtr = &b.nullStr
	case *InvocationID:
		spanPtr = &b.nullStr
	case *InvocationIDSet:
		spanPtr = &b.strSlice
	case **tspb.Timestamp:
		spanPtr = &b.nullTime
	case *pb.TestStatus:
		spanPtr = &b.i64
	case *pb.Invocation_State:
		spanPtr = &b.i64
	case **typepb.Variant:
		spanPtr = &b.strSlice
	case *[]*typepb.StringPair:
		spanPtr = &b.strSlice
	case *[]*pb.Artifact:
		spanPtr = &b.byteSlice
	case proto.Message:
		spanPtr = &b.byteSlice
	case *Compressed:
		spanPtr = &b.byteSlice
	case *CompressedProto:
		spanPtr = &b.byteSlice
	default:
		spanPtr = goPtr
	}

	if err := row.Column(col, spanPtr); err != nil {
		return errors.Annotate(err, "failed to read column %q", row.ColumnName(col)).Err()
	}

	if spanPtr == goPtr {
		return nil
	}

	var err error
	switch goPtr := goPtr.(type) {
	case *string:
		*goPtr = ""
		if b.nullStr.Valid {
			*goPtr = b.nullStr.StringVal
		}

	case *InvocationID:
		*goPtr = ""
		if b.nullStr.Valid {
			*goPtr = InvocationIDFromRowID(b.nullStr.StringVal)
		}

	case *InvocationIDSet:
		*goPtr = make(InvocationIDSet, len(b.strSlice))
		for _, rowID := range b.strSlice {
			goPtr.Add(InvocationIDFromRowID(rowID))
		}

	case **tspb.Timestamp:
		*goPtr = nil
		if b.nullTime.Valid {
			*goPtr = pbutil.MustTimestampProto(b.nullTime.Time)
		}

	case *pb.Invocation_State:
		*goPtr = pb.Invocation_State(b.i64)

	case *pb.TestStatus:
		*goPtr = pb.TestStatus(b.i64)

	case **typepb.Variant:
		if *goPtr, err = pbutil.VariantFromStrings(b.strSlice); err != nil {
			// If it was written to Spanner, it should have been validated.
			panic(err)
		}

	case *[]*typepb.StringPair:
		*goPtr = make([]*typepb.StringPair, len(b.strSlice))
		for i, p := range b.strSlice {
			if (*goPtr)[i], err = pbutil.StringPairFromString(p); err != nil {
				// If it was written to Spanner, it should have been validated.
				panic(err)
			}
		}

	case *[]*pb.Artifact:
		container := &internalpb.Artifacts{}
		if err := proto.Unmarshal(b.byteSlice, container); err != nil {
			// If it was written to Spanner, it should have been validated.
			panic(err)
		}
		*goPtr = container.ArtifactsV1

	case proto.Message:
		if reflect.ValueOf(goPtr).IsNil() {
			return errors.Reason("nil pointer encountered").Err()
		}
		if err := proto.Unmarshal(b.byteSlice, goPtr); err != nil {
			// If it was written to Spanner, it should have been validated.
			panic(err)
		}

	case *Compressed:
		if len(b.byteSlice) == 0 {
			// do not set to nil; otherwise we lose the buffer.
			*goPtr = (*goPtr)[:0]
		} else {
			// *goPtr might be pointing to an existing memory buffer.
			// Try to reuse it for decoding.
			if *goPtr, err = decompress(b.byteSlice, *goPtr); err != nil {
				return err
			}
		}

	case *CompressedProto:
		goPtr.Message.Reset()
		if len(b.byteSlice) != 0 {
			if b.byteSlice2, err = decompress(b.byteSlice, b.byteSlice2); err != nil {
				return err
			}
			if err := proto.Unmarshal(b.byteSlice2, goPtr.Message); err != nil {
				return err
			}
		}

	default:
		panic("impossible")
	}
	return nil
}

// ToSpanner converts values from Go types to Spanner types. In addition to
// supported types in FromSpanner, also supports []interface{} and
// map[string]interface{}.
func ToSpanner(v interface{}) interface{} {
	switch v := v.(type) {
	case InvocationID:
		return v.RowID()

	case InvocationIDSet:
		ret := make([]string, 0, len(v))
		for id := range v {
			ret = append(ret, id.RowID())
		}
		sort.Strings(ret)
		return ret

	case *tspb.Timestamp:
		if v == nil {
			return spanner.NullTime{}
		}
		ret := spanner.NullTime{Valid: true}
		ret.Time, _ = ptypes.Timestamp(v)
		// ptypes.Timestamp always returns a timestamp, even
		// if the returned err is non-nil, see its documentation.
		// The error is returned only if the timestamp violates its
		// own invariants, which will be caught on the attempt to
		// insert this value into Spanner.
		// This is consistent with the behavior of spanner.Insert() and
		// other mutating functions that ignore invalid time.Time
		// until it is time to apply the mutation.
		// Not returning an error here significantly simplifies usage
		// of this function and functions based on this one.
		return ret

	case pb.Invocation_State:
		return int64(v)

	case pb.TestStatus:
		return int64(v)

	case *typepb.Variant:
		return pbutil.VariantToStrings(v)

	case []*typepb.StringPair:
		return pbutil.StringPairsToStrings(v...)

	case []*pb.Artifact:
		ret, err := proto.Marshal(&internalpb.Artifacts{ArtifactsV1: v})
		if err != nil {
			panic(err)
		}
		return ret

	case proto.Message:
		if isMessageNil(v) {
			// Do not store empty messages.
			return []byte(nil)
		}

		ret, err := proto.Marshal(v)
		if err != nil {
			panic(err)
		}
		return ret

	case []interface{}:
		ret := make([]interface{}, len(v))
		for i, el := range v {
			ret[i] = ToSpanner(el)
		}
		return ret

	case map[string]interface{}:
		ret := make(map[string]interface{}, len(v))
		for key, el := range v {
			ret[key] = ToSpanner(el)
		}
		return ret

	case Compressed:
		if len(v) == 0 {
			// Do not store empty bytes.
			return []byte(nil)
		}
		return compress(v)

		// It must go before case proto.Message
	case CompressedProto:
		if isMessageNil(v.Message) {
			// Do not store empty messages.
			return []byte(nil)
		}

		// This might be optimized by reusing the memory buffer.
		raw, err := proto.Marshal(v.Message)
		if err != nil {
			panic(err)
		}
		return compress(raw)

	default:
		return v
	}
}

// ToSpannerSlice converts a slice of Go values to a slice of Spanner values.
// See also ToSpanner.
func ToSpannerSlice(values ...interface{}) []interface{} {
	return ToSpanner(values).([]interface{})
}

// ToSpannerMap converts a map of Go values to a map of Spanner values.
// See also ToSpanner.
func ToSpannerMap(values map[string]interface{}) map[string]interface{} {
	return ToSpanner(values).(map[string]interface{})
}

// UpdateMap is a shortcut for spanner.UpdateMap with ToSpannerMap applied to
// in.
func UpdateMap(table string, in map[string]interface{}) *spanner.Mutation {
	return spanner.UpdateMap(table, ToSpannerMap(in))
}

// InsertMap is a shortcut for spanner.InsertMap with ToSpannerMap applied to
// in.
func InsertMap(table string, in map[string]interface{}) *spanner.Mutation {
	return spanner.InsertMap(table, ToSpannerMap(in))
}

// InsertOrUpdateMap is a shortcut for spanner.InsertOrUpdateMap with ToSpannerMap applied to
// in.
func InsertOrUpdateMap(table string, in map[string]interface{}) *spanner.Mutation {
	return spanner.InsertOrUpdateMap(table, ToSpannerMap(in))
}

// ReadWriteTransaction calls Client(ctx).ReadWriteTransaction and unwraps
// a "transaction is aborted" errors such that the spanner client properly
// retries the function.
func ReadWriteTransaction(ctx context.Context, f func(context.Context, *spanner.ReadWriteTransaction) error) (commitTimestamp time.Time, err error) {
	return Client(ctx).ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		err := f(ctx, txn)
		if unwrapped := errors.Unwrap(err); spanner.ErrCode(unwrapped) == codes.Aborted {
			err = unwrapped
		}
		return err
	})
}

// QueryFirstRow executes a query, reads the first row into ptrs and stops the
// iterator. Returns ErrNoResults if the query does not return at least one row.
func QueryFirstRow(ctx context.Context, txn Txn, st spanner.Statement, ptrs ...interface{}) error {
	st.Params = ToSpannerMap(st.Params)
	it := txn.Query(ctx, st)
	defer it.Stop()
	row, err := it.Next()
	if err != nil {
		return err
	}
	return FromSpanner(row, ptrs...)
}

// Query executes a query.
// Ensures st.Params are Spanner-compatible by modifying st.Params in place.
func Query(ctx context.Context, txn Txn, st spanner.Statement, fn func(row *spanner.Row) error) error {
	st.Params = ToSpannerMap(st.Params)
	return txn.Query(ctx, st).Do(fn)
}

func compress(data []byte) []byte {
	out := make([]byte, 0, len(data)/2+len(zstdHeader)) // hope for at least 2x compression
	out = append(out, zstdHeader...)
	return zstdEncoder.EncodeAll(data, out)
}

func decompress(src, dest []byte) ([]byte, error) {
	if !bytes.HasPrefix(src, zstdHeader) {
		return nil, errors.Reason("expected ztd header").Err()
	}

	dest, err := zstdDecoder.DecodeAll(src[len(zstdHeader):], dest[:0])
	if err != nil {
		return nil, errors.Annotate(err, "failed to decode from zstd").Err()
	}
	return dest, nil
}

func isMessageNil(m proto.Message) bool {
	return reflect.ValueOf(m).IsNil()
}
