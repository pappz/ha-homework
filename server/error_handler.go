package server

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrRespInternalError = errors.New("internal error")
)

type ErrorResponse struct {
	Error   bool
	Message string
}

func responseError(w http.ResponseWriter, e error) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	if e == ErrRespInternalError {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

	resp := ErrorResponse{
		true,
		e.Error(),
	}
	if j, err := json.Marshal(resp); err == nil {
		_, _ = w.Write(j)
	}
}
