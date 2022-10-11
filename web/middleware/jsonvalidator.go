package middleware

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var (
	errFailedToReadBody = errors.New("error during read body")
	errMissingBody      = errors.New("missing body")
)

type InputValidation interface {
	Validate() error
}

func unmarshalAndValidate(r *http.Request, v interface{}) error {
	body, err := io.ReadAll(r.Body)
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

	iv, ok := v.(InputValidation)
	if !ok {
		return nil
	}

	return iv.Validate()
}
