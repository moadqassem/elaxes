package config

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"reflect"
)

var (
	expectedConfigs = &Configs{
		Server: &Server{
			Host: "localhost",
			Port: 8080,
		},
		Elasticsearch: &Elasticsearch{
			ClusterAddress: "localhost:9200",
		},
		Prometheus: &Prometheus{
			Port:     8090,
			Endpoint: "/metrics",
		},
	}
)

func TestLoadConfigs(t *testing.T) {
	Convey("testing the configuration loading functionality", t, func() {
		cfg, err := LoadConfigs("config_test.json")
		So(err, ShouldBeNil)
		So(cfg, ShouldNotBeNil)
		So(reflect.DeepEqual(cfg, expectedConfigs), ShouldBeTrue)
	})
}
