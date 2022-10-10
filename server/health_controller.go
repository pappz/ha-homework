package server

import (
	"io"
	"net/http"
)

const (
	healthResponse = `{"alive": true}`
)

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = io.WriteString(w, healthResponse)
}
