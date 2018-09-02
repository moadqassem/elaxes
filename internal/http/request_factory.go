package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// FactoryManager manages the http requests between the app and the remote cluster.
type FactoryManager interface {
	ExecuteRequest(method string, endpoint string, model interface{}) error
}

type requestFactory struct {
	client *http.Client
}

// NewHealthRequest returns an http request for the elasticsearch cluster based on the provided endpoint.
func NewRequestFactory() FactoryManager {
	return &requestFactory{
		client: http.DefaultClient,
	}
}

// ExecuteRequest dispatches the prepared request and returns a response or an error.
func (e *requestFactory) ExecuteRequest(method string, endpoint string, model interface{}) error {
	request, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	res, err := e.client.Do(request)
	defer res.Body.Close()

	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, model)
}
