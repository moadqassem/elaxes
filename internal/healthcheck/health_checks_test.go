package healthcheck

import (
	"testing"

	"github.com/moadqassem/elaxes/config"
	"github.com/moadqassem/elaxes/internal"
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"errors"
)

var (
	elasticsearchConfig = &config.Elasticsearch{
		ClusterAddress: "http://localhost:8000",
	}

	expectedClusterHealth = &internal.ClusterHealth{
		ClusterName:                 "elasticsearch-testing-cluster",
		Status:                      "green",
		TimedOut:                    false,
		NumberOfNodes:               2,
		NumberOfDataNodes:           2,
		ActivePrimaryShards:         40,
		ActiveShards:                50,
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

func TestNewDefaultChecker(t *testing.T) {
	Convey("testing a new health checker", t, func() {
		checker, err := NewHealthChecker(elasticsearchConfig, nil)
		So(err, ShouldBeNil)
		So(checker, ShouldNotBeNil)
	})
}

func TestDefaultChecker_ClusterHealth(t *testing.T) {
	Convey("testing a the cluster health check", t, func() {
		checker, err := NewHealthChecker(elasticsearchConfig, &mockCluster{
			data: expectedClusterHealth,
		})
		So(err, ShouldBeNil)
		So(checker, ShouldNotBeNil)

		health, err := checker.ClusterHealth()
		So(err, ShouldBeNil)
		So(reflect.DeepEqual(health, expectedClusterHealth), ShouldBeTrue)
	})
}

type mockCluster struct {
	data interface{}
}

func (m *mockCluster) ExecuteRequest(method string, endpoint string, model interface{}) error {
	modelType := reflect.ValueOf(model)
	if modelType.Kind() != reflect.Ptr {
		return errors.New("unsupported model type")
	}

	modelType.Elem().Set(reflect.ValueOf(m.data))
	return nil
}
