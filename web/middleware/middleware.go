package middleware

import (
	"net/http"

	"github.com/pappz/ha-homework/service"
	"github.com/pappz/ha-homework/web/handler"
)

type Middleware struct {
	service service.Sector
}

func NewMiddleware(service service.Sector) Middleware {
	return Middleware{
		service: service,
	}
}

func (m Middleware) Handle(h handler.Handler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ri := handler.RequestInfo{
			Data:    h.RequestDataType(),
			W:       w,
			R:       r,
			Service: m.service,
		}
		if err := unmarshalAndValidate(r, ri.Data); err != nil {
			//log.Errorf(err.Error())
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
