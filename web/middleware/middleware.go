package middleware

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/pappz/ha-homework/service"
)

var (
	ErrRespInternalError = errors.New("internal error")
	errFailedToReadBody  = errors.New("error during read body")
	errMissingBody       = errors.New("missing body")
)

// ErrorResponse is the generic Json format for http error responses
type ErrorResponse struct {
	Message string
}

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
		if err := m.unmarshalAndValidate(r.Body, ri.Data); err != nil {
			m.responseError(w, err)
			return
		}

		v, err := h.Do(ri)
		if err != nil {
			m.responseError(w, err)
			return
		}

		if v == nil {
			return
		}

		if err := m.responseJson(w, v); err != nil {
			// log.Errorf("failed to send json: %s", err.Error())
		}
	}
}

func (m Middleware) unmarshalAndValidate(r io.Reader, v interface{}) error {
	body, err := io.ReadAll(r)
	if err != nil {
		return errFailedToReadBody
	}

	if v == nil {
		return nil
	}

	if len(body) == 0 {
		return errMissingBody
	}

	if err := json.Unmarshal(body, v); err != nil {
		return err
	}

	iv, ok := v.(Json)
	if !ok {
		return nil
	}

	return iv.Validate()
}

// responseError response with error to the request. Set the proper http headers
// and based on the error type send out the required error message.
func (m Middleware) responseError(w http.ResponseWriter, e error) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	if e == ErrRespInternalError {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

	resp := ErrorResponse{
		e.Error(),
	}

	// json marshal error never should happen so ignore it
	j, _ := json.Marshal(resp)
	_, _ = w.Write(j)
	return
}

// responseJson marshal the response content and send out to the http request with
// the proper headers.
func (m Middleware) responseJson(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	j, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = w.Write(j)
	return err
}
