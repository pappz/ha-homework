package middleware

import (
	"encoding/json"
	"errors"
	"io"
)

var (
	errFailedToReadBody = errors.New("error during read body")
	errMissingBody      = errors.New("missing body")
)

type Json interface {
	Validate() error
}

func unmarshalAndValidate(r io.Reader, v interface{}) error {
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
