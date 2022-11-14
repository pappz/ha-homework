package api

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pappz/ha-homework/web/middleware"
)

const (
	HealthResponse = `{"alive": true}`
)

// RegisterHealthHandler sets up the routing of the HTTP handlers.
func RegisterHealthHandler(router *mux.Router) {
	h := health{}
	router.HandleFunc("/health", middleware.Handle(h.Handle, nil)).Methods(http.MethodGet)
}

// health controller to ensure the service is alive
type health struct {
}

func (h health) Handle(ri *middleware.RequestInfo) (middleware.ResponseData, error) {
	_, err := io.WriteString(ri.W, HealthResponse)
	return nil, err
}
