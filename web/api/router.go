package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pappz/ha-homework/service"
	"github.com/pappz/ha-homework/web/middleware"
)

func Router(s service.Sector) *mux.Router {
	m := middleware.NewJsonParser(s)
	router := mux.NewRouter()
	router.StrictSlash(true)
	router.HandleFunc("/health", m.Handle(Health{})).Methods(http.MethodGet)
	router.HandleFunc("/databank", m.Handle(Location{})).Methods(http.MethodPost)
	return router
}
