package controllers

import (
	"io"

	"github.com/pappz/ha-homework/web/handler"
)

const (
	HealthResponse = `{"alive": true}`
)

type Health struct {
}

func (h Health) RequestDataType() interface{} {
	return nil
}

func (h Health) Do(ri handler.RequestInfo) (handler.ResponseData, error) {
	_, err := io.WriteString(ri.W, HealthResponse)
	return nil, err
}
