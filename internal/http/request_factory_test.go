package http

import (
	"testing"

	"fmt"
	"github.com/moadqassem/elaxes/internal"
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"net/http"
	"reflect"
	"time"
)

var (
	serverAddress         = "localhost:8080"
	expectedClusterHealth = &internal.ClusterHealth{
		ClusterName:                 "testing-cluster",
		Status:                      "green",
		TimedOut:                    false,
		NumberOfNodes:               1,
		NumberOfDataNodes:           1,
		ActivePrimaryShards:         5,
		ActiveShards:                5,
		RelocatingShards:            0,
		InitializingShards:          0,
		UnassignedShards:            0,
		DelayedUnassignedShards:     0,
		NumberOfPendingTasks:        0,
		NumberOfInFlightFetch:       0,
		TaskMaxWaitingInQueueMillis: 0,
		ActiveShardsPercentAsNumber: 100.0,
	}
)

func TestNewElasticRequest(t *testing.T) {
	Convey("testing the creation of an elastic request", t, func() {
		r := NewRequestFactory()
		So(r, ShouldNotBeNil)
	})
}

func TestElasticRequest_ExecuteRequest(t *testing.T) {
	server := http.Server{
		Addr:    serverAddress,
		Handler: &testHandler{},
	}
	go server.ListenAndServe()

	endpoint := fmt.Sprintf("http://%v", serverAddress)
	if !serverIsUp(3, endpoint) {
		t.Skip("test was skipped, no testing server was up")
	}
	Convey("testing the execution of the http request", t, func() {
		r := NewRequestFactory()
		So(r, ShouldNotBeNil)

		health := &internal.ClusterHealth{}
		err := r.ExecuteRequest("GET", endpoint, &health)
		So(err, ShouldBeNil)
		So(reflect.DeepEqual(health, expectedClusterHealth), ShouldBeTrue)
	})
	if err := server.Close(); err != nil {
		t.Fatal(err)
	}
}

type testHandler struct{}

func (th *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{
		"cluster_name" : "testing-cluster",
		"status" : "green",
		"timed_out" : false,
		"number_of_nodes" : 1,
		"number_of_data_nodes" : 1,
		"active_primary_shards" : 5,
		"active_shards" : 5,
		"relocating_shards" : 0,
		"initializing_shards" : 0,
		"unassigned_shards" : 0,
		"delayed_unassigned_shards" : 0,
		"number_of_pending_tasks" : 0,
		"number_of_in_flight_fetch" : 0,
		"task_max_waiting_in_queue_millis" : 0,
		"active_shards_percent_as_number" : 100.0
	}`)
}

func serverIsUp(maxTries int, address string) bool {
	tries := 0
	for {
		if tries < maxTries {
			if _, err := http.Get(address); err != nil {
				time.Sleep(200 * time.Millisecond)
				tries++
				continue
			}
			return true
		}
		return false
	}
}
