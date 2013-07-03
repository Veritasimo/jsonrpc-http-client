package jsonrpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ServiceProxy struct {
	Version     string
	ServiceUrl  string
	ServiceName string
	id          int
}

func NewProxy(url string, name string) *ServiceProxy {
	proxy := new(ServiceProxy)
	proxy.Version = "1.0"
	proxy.ServiceUrl = url
	proxy.ServiceName = name

	return proxy
}

func (s *ServiceProxy) Call(method string, params ...interface{}) (interface{}, error) {
	if s.Version != "1.0" {
		return nil, errors.New("Unsupported version")
	}

	if s.ServiceUrl == "" {
		return nil, errors.New("No service url specified.")
	}

	if s.ServiceName != "" {
		method = fmt.Sprintf("%s.%s", s.ServiceName, method)
	}
	var payload = map[string]interface{}{
		"version": s.Version,
		"method":  method,
		"params":  params,
		"id":      s.id,
	}

	s.id += 1
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(data)
	resp, err := http.Post(s.ServiceUrl, "application/json", buf)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	var respPayload interface{}
	err = decoder.Decode(&respPayload)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid response from server. Status: %s. %s", resp.Status, err.Error()))
	}

	return respPayload, nil
}
