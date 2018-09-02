package healthcheck

import (
	"fmt"
	"github.com/moadqassem/elaxes/config"
	"github.com/moadqassem/elaxes/internal"
	"github.com/moadqassem/elaxes/internal/http"
)

const (
	clusterHealthCheck = iota
)

var healthChecks = map[int]string{
	clusterHealthCheck: `/_cluster/health`,
}

type healthChecker struct {
	clusterEndpoint string
	requestFactory  http.FactoryManager
}

// NewHealthChecker creates a new default health checker with the provided configurations.
func NewHealthChecker(configs *config.Elasticsearch, factory http.FactoryManager) (*healthChecker, error) {
	if factory == nil {
		factory = http.NewRequestFactory()
	}
	return &healthChecker{
		clusterEndpoint: configs.ClusterAddress,
		requestFactory:  factory,
	}, nil
}

// ClusterHealth returns general status about the cluster health.
func (d *healthChecker) ClusterHealth() (*internal.ClusterHealth, error) {
	ch := &internal.ClusterHealth{}
	err := d.requestFactory.ExecuteRequest("GET", fmt.Sprintf("%v%v", d.clusterEndpoint,
		healthChecks[clusterHealthCheck]), &ch)

	if err != nil {
		return nil, err
	}

	return ch, nil
}
