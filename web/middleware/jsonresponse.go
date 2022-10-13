package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrRespInternalError = errors.New("internal error")
)

type ErrorResponse struct {
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
		e.Error(),
	}

	// json marshal error never should happen so ignore it
	j, _ := json.Marshal(resp)
	_, _ = w.Write(j)
	return
}

func responseJson(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	j, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = w.Write(j)
	return err
}
