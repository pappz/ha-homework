package handler

import (
	"net/http"

	"github.com/pappz/ha-homework/service"
)

type RequestInfo struct {
	Data    interface{}
	W       http.ResponseWriter
	R       *http.Request
	Service service.Sector
}

type ResponseData interface{}

type Handler interface {
	RequestDataType() interface{}
	Do(RequestInfo) (ResponseData, error)
}
