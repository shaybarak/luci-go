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

package recorder

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	tspb "github.com/golang/protobuf/ptypes/timestamp"

	"go.chromium.org/luci/common/isolated"
	"go.chromium.org/luci/common/isolatedclient/isolatedfake"

	"go.chromium.org/luci/resultdb/internal"
	"go.chromium.org/luci/resultdb/internal/recorder/chromium/formats"
	"go.chromium.org/luci/resultdb/internal/span"
	"go.chromium.org/luci/resultdb/internal/tasks"
	"go.chromium.org/luci/resultdb/internal/testutil"
	"go.chromium.org/luci/resultdb/pbutil"
	pb "go.chromium.org/luci/resultdb/proto/rpc/v1"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestValidateDeriveInvocationRequest(t *testing.T) {
	Convey(`TestValidateDeriveInvocationRequest`, t, func() {
		Convey(`Valid`, func() {
			req := &pb.DeriveInvocationRequest{
				SwarmingTask: &pb.DeriveInvocationRequest_SwarmingTask{
					Hostname: "swarming-host.appspot.com",
					Id:       "beeff00d",
				},
			}
			So(validateDeriveInvocationRequest(req), ShouldBeNil)
		})

		Convey(`Invalid swarming_task`, func() {
			Convey(`missing`, func() {
				req := &pb.DeriveInvocationRequest{}
				So(validateDeriveInvocationRequest(req), ShouldErrLike, "swarming_task missing")
			})

			Convey(`missing hostname`, func() {
				req := &pb.DeriveInvocationRequest{
					SwarmingTask: &pb.DeriveInvocationRequest_SwarmingTask{
						Id: "beeff00d",
					},
				}
				So(validateDeriveInvocationRequest(req), ShouldErrLike, "swarming_task.hostname missing")
			})

			Convey(`bad hostname`, func() {
				req := &pb.DeriveInvocationRequest{
					SwarmingTask: &pb.DeriveInvocationRequest_SwarmingTask{
						Hostname: "https://swarming-host.appspot.com",
						Id:       "beeff00d",
					},
				}
				So(validateDeriveInvocationRequest(req), ShouldErrLike,
					"swarming_task.hostname should not have prefix")
			})

			Convey(`missing id`, func() {
				req := &pb.DeriveInvocationRequest{
					SwarmingTask: &pb.DeriveInvocationRequest_SwarmingTask{
						Hostname: "swarming-host.appspot.com",
					},
				}
				So(validateDeriveInvocationRequest(req), ShouldErrLike, "swarming_task.id missing")
			})
		})
	})
}

func TestDeriveInvocation(t *testing.T) {
	Convey(`TestDeriveInvocation`, t, func() {
		ctx := testutil.SpannerTestContext(t)

		testutil.MustApply(ctx, testutil.InsertInvocation("inserted", pb.Invocation_FINALIZED, nil))

		Convey(`calling to shouldWriteInvocation works`, func() {
			Convey(`if we already have the invocation written`, func() {
				doWrite, err := shouldWriteInvocation(ctx, span.Client(ctx).Single(), "inserted")
				So(err, ShouldBeNil)
				So(doWrite, ShouldBeFalse)
			})

			Convey(`if we don't yet have the invocation written`, func() {
				doWrite, err := shouldWriteInvocation(ctx, span.Client(ctx).Single(), "another")
				So(err, ShouldBeNil)
				So(doWrite, ShouldBeTrue)
			})
		})

		ctx = internal.WithHTTPClient(ctx, &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		})

		// Set up fake isolatedserver.
		isoFake := isolatedfake.New()
		isoServer := httptest.NewServer(isoFake)
		defer isoServer.Close()

		// Inject isolated objects.
		testJSON := []byte(`{
			"all_tests": [
				"FooTest.TestDoBar",
				"FooTest.DoesBar"
			],
			"per_iteration_data": [{
				"MyInstantiation/FooTest.DoesBar/1": [
					{
						"status": "SUCCESS",
						"elapsed_time_ms": 100
					}
				],
				"FooTest.TestDoBar": [
					{
						"status": "CRASH",
						"elapsed_time_ms": 200
					},
					{
						"status": "FAILURE",
						"elapsed_time_ms": 300
					}
				]
			}]
		}`)
		size := int64(len(testJSON))
		fileDigest := isoFake.Inject("ns", testJSON)
		isoOut := isolated.Isolated{
			Files: map[string]isolated.File{"output.json": {Digest: fileDigest, Size: &size}},
		}
		isoOutBytes, err := json.Marshal(isoOut)
		So(err, ShouldBeNil)
		outputsDigest := isoFake.Inject("ns", isoOutBytes)

		// Set up fake swarming service.
		// This service only supports one task and returns the same one regardless of input.
		swarmingFake := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, fmt.Sprintf(`{
				"state": "COMPLETED",
				"outputs_ref": {
					"isolatedserver": "%s",
					"namespace": "ns",
					"isolated": "%s"
				},
				"tags": [
					"bucket:bkt",
					"buildername:blder",
					"test_suite:foo_unittests",
					"ninja_target://tests:tests"
				],
				"created_ts": "2019-10-14T13:49:16.01",
				"completed_ts": "2019-10-14T14:49:16.01"
			}`, isoServer.URL, outputsDigest))
		}))
		defer swarmingFake.Close()

		// Define base request we'll be using.
		swarmingHostname := strings.TrimPrefix(swarmingFake.URL, "https://")
		req := &pb.DeriveInvocationRequest{
			SwarmingTask: &pb.DeriveInvocationRequest_SwarmingTask{
				Hostname: swarmingHostname,
			},
		}

		recorder := newTestRecorderServer()
		recorder.DerivedInvBQTable = &pb.BigQueryExport{
			Project:     "project",
			Dataset:     "dataset",
			Table:       "table",
			TestResults: &pb.BigQueryExport_TestResults{},
		}

		Convey(`inserts a new invocation`, func() {
			req.SwarmingTask.Id = "completed-task"
			inv, err := recorder.DeriveInvocation(ctx, req)
			So(err, ShouldBeNil)

			So(inv, ShouldResembleProto, &pb.Invocation{
				Name:                inv.Name, // inv.Name is non-determinisic in this test
				State:               pb.Invocation_FINALIZED,
				CreateTime:          &tspb.Timestamp{Seconds: 1571060956, Nanos: 1e7},
				Tags:                pbutil.StringPairs(formats.OriginalFormatTagKey, formats.FormatGTest),
				FinalizeTime:        &tspb.Timestamp{Seconds: 1571064556, Nanos: 1e7},
				Deadline:            &tspb.Timestamp{Seconds: 1571064556, Nanos: 1e7},
				IncludedInvocations: []string{inv.Name + "::batch::0"},
			})

			// Assert we wrote correct test results.
			txn := span.Client(ctx).ReadOnlyTransaction()
			defer txn.Close()

			invIDs, err := span.ReadReachableInvocations(ctx, txn, 100, span.NewInvocationIDSet(span.MustParseInvocationName(inv.Name)))
			So(err, ShouldBeNil)
			trs, _, err := span.QueryTestResults(ctx, txn, span.TestResultQuery{
				InvocationIDs: invIDs,
				PageSize:      100,
			})
			So(err, ShouldBeNil)
			So(trs, ShouldHaveLength, 3)
			So(trs[0].TestId, ShouldEqual, "ninja://tests:tests/FooTest.DoesBar")
			So(trs[0].Status, ShouldEqual, pb.TestStatus_PASS)
			So(trs[0].Variant, ShouldResembleProto, pbutil.Variant(
				"bucket", "bkt",
				"builder", "blder",
				"test_suite", "foo_unittests",
				"param/instantiation", "MyInstantiation",
				"param/id", "1",
			))
			So(trs[1].TestId, ShouldEqual, "ninja://tests:tests/FooTest.TestDoBar")
			So(trs[1].Status, ShouldEqual, pb.TestStatus_CRASH)
			So(trs[2].TestId, ShouldEqual, "ninja://tests:tests/FooTest.TestDoBar")
			So(trs[2].Status, ShouldEqual, pb.TestStatus_FAIL)

			// Read InvocationTask to confirm it's added.
			taskKey := tasks.BQExport.Key(fmt.Sprintf("%s:0", span.MustParseInvocationName(inv.Name)))
			var payload []byte
			testutil.MustReadRow(ctx, "InvocationTasks", taskKey, map[string]interface{}{
				"Payload": &payload,
			})
			bqExports := &pb.BigQueryExport{}
			err = proto.Unmarshal(payload, bqExports)
			So(err, ShouldBeNil)
			So(bqExports, ShouldResembleProto, recorder.DerivedInvBQTable)
		})
	})
}

func TestBatchInsertTestResults(t *testing.T) {
	Convey(`TestBatchInsertTestResults`, t, func() {
		ctx := testutil.SpannerTestContext(t)
		startTS := ptypes.TimestampNow()

		inv := &pb.Invocation{
			State:        pb.Invocation_FINALIZED,
			CreateTime:   startTS,
			FinalizeTime: startTS,
			Deadline:     startTS,
		}
		results := []*pb.TestResult{
			{TestId: "Foo.DoBar", ResultId: "0", Status: pb.TestStatus_PASS},
			{TestId: "Foo.DoBar", ResultId: "1", Status: pb.TestStatus_FAIL},
			{TestId: "Foo.DoBar", ResultId: "2", Status: pb.TestStatus_CRASH},
		}
		recorder := newTestRecorderServer()

		checkBatches := func(baseID span.InvocationID, actualInclusions span.InvocationIDSet, expectedBatches [][]*pb.TestResult) {
			// Check included Invocations.
			expectedInclusions := make(span.InvocationIDSet, len(expectedBatches))
			for i := range expectedBatches {
				expectedInclusions.Add(batchInvocationID(baseID, i))
			}
			So(actualInclusions, ShouldResemble, expectedInclusions)

			txn := span.Client(ctx).ReadOnlyTransaction()
			defer txn.Close()

			// Check that the TestResults are batched as expected.
			for i, expectedBatch := range expectedBatches {
				actualResults, _, err := span.QueryTestResults(ctx, txn, span.TestResultQuery{
					InvocationIDs: span.NewInvocationIDSet(batchInvocationID(baseID, i)),
					PageSize:      100,
				})
				So(err, ShouldBeNil)
				So(actualResults, ShouldHaveLength, len(expectedBatch))

				for k, expectedResult := range expectedBatch {
					So(actualResults[k].TestId, ShouldEqual, expectedResult.TestId)
					So(actualResults[k].ResultId, ShouldEqual, strconv.Itoa(k))
					So(actualResults[k].Status, ShouldEqual, expectedResult.Status)
				}
			}
		}

		Convey(`for one batch`, func() {
			baseID := span.InvocationID("one_batch")
			inv.Name = baseID.Name()
			actualInclusions, err := recorder.batchInsertTestResults(ctx, inv, results, 5)
			So(err, ShouldBeNil)
			checkBatches(baseID, actualInclusions, [][]*pb.TestResult{results})
		})

		Convey(`for multiple batches`, func() {
			baseID := span.InvocationID("multiple_batches")
			inv.Name = baseID.Name()
			actualInclusions, err := recorder.batchInsertTestResults(ctx, inv, results, 2)
			So(err, ShouldBeNil)
			checkBatches(baseID, actualInclusions, [][]*pb.TestResult{results[:2], results[2:]})

			Convey(`with batch size a factor of number of TestResults`, func() {
				baseID := span.InvocationID("multiple_batches_factor")
				inv.Name = baseID.Name()
				actualInclusions, err := recorder.batchInsertTestResults(ctx, inv, results, 1)
				So(err, ShouldBeNil)
				checkBatches(baseID, actualInclusions,
					[][]*pb.TestResult{results[:1], results[1:2], results[2:]})
			})
		})
	})
}
