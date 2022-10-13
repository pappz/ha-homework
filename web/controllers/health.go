package controllers

import (
	"github.com/pappz/ha-homework/web/middleware"
	"io"
)

const (
	HealthResponse = `{"alive": true}`
)

// Health controller to ensure the service is alive
type Health struct {
}

func (h Health) RequestDataType() middleware.Json {
	return nil
}

func (h Health) Do(ri middleware.RequestInfo) (middleware.ResponseData, error) {
	_, err := io.WriteString(ri.W, HealthResponse)
	return nil, err
}
