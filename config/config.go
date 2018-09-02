package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// Configs contains all the configurations for the app to start.
type Configs struct {
	Server        *Server
	Elasticsearch *Elasticsearch
	Prometheus    *Prometheus
}

type Server struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}

type Elasticsearch struct {
	ClusterAddress string `json:"cluster_address"`
}

type Prometheus struct {
	Port     int32  `json:"port"`
	Endpoint string `json:"endpoint"`
}

// LoadConfigs load the configuration from the provided path.
func LoadConfigs(path string) (*Configs, error) {
	if path != "" {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}

		cfg := &Configs{}
		if err := json.Unmarshal(data, &cfg); err != nil {
			return nil, err
		}

		return cfg, nil
	}
	return nil, errors.New("invalid file path")
}
