package middleware

import (
	"net/http"

	"github.com/pappz/ha-homework/service"
)

// RequestInfo pass to all controllers. The Data is the struct
// what required by the controller as input parameter
type RequestInfo struct {
	Data    Json
	W       http.ResponseWriter
	R       *http.Request
	Service service.Sector
}

// ResponseData will be sent out by the middleware to the http
// request as a http response. It could be nil.
type ResponseData interface{}

// Handler interfaces used by the middleware. The RequestDataType
// return a struct what the middleware try to unmarshal from the
// http request. The Do will be called by middleware when received
// a http request. In this function should implement the
// controller's logic.
type Handler interface {
	RequestDataType() Json
	Do(RequestInfo) (ResponseData, error)
}
