package middleware

import (
	"net/http"

	"github.com/pappz/ha-homework/service"
)

// Middleware receive http requests, pass through the unmarshaled inputs to the
// controllers and handle errors with the proper status codes. The middlware
// manage the http headers and status codes in the response.
type Middleware struct {
	service service.Sector
}

// NewMiddleware instantiate a new Middleware
func NewMiddleware(service service.Sector) Middleware {
	return Middleware{
		service: service,
	}
}

// Handle doing validation on the Json request. In case of err it send response
// with specific error message in json format. After the controller returned with
// results the middleware send out the response data in json or in case of error
// response with error code and message reason in json format.
func (m Middleware) Handle(h Handler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ri := RequestInfo{
			Data:    h.RequestDataType(),
			W:       w,
			R:       r,
			Service: m.service,
		}
		if err := unmarshalAndValidate(r.Body, ri.Data); err != nil {
			responseError(w, err)
			return
		}

		v, err := h.Do(ri)
		if err != nil {
			responseError(w, err)
			return
		}

		if v == nil {
			return
		}

		if err := responseJson(w, v); err != nil {
			// log.Errorf("failed to send json: %s", err.Error())
		}
	}
}
