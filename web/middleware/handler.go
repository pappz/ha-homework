package middleware

import (
	"net/http"
)

// RequestInfo will be passed to all controllers. The Payload is the struct
// what required by the controller as input parameter.
type RequestInfo struct {
	W       http.ResponseWriter
	R       *http.Request
	Payload Json
}

// ResponseData will be sent out by the middleware to the http
// request as a http response. It could be nil.
type ResponseData interface{}

// Handler interfaces used by the JsonParser.
type Handler func(*RequestInfo) (ResponseData, error)
