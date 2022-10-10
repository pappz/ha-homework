package server

import (
	"errors"
	"net/http"
)

var (
	errInvalidCoordinate = errors.New("invalid coordinate")
	errInvalidVelocity   = errors.New("invalid velocity")
)

type RequestData struct {
	X   float64 `json:",string"`
	Y   float64 `json:",string"`
	Z   float64 `json:",string"`
	Vel float64 `json:",string"`
}

func (rd RequestData) Validate() error {
	if rd.X < 0 {
		return errInvalidCoordinate
	}
	if rd.Y < 0 {
		return errInvalidCoordinate
	}
	if rd.Z < 0 {
		return errInvalidCoordinate
	}
	if rd.Vel <= 0 {
		return errInvalidVelocity
	}
	return nil
}

func location(w http.ResponseWriter, r *http.Request) {
	rd := &RequestData{}
	if err := unmarshalAndValidate(r, rd); err != nil {
		responseError(w, err)
		return
	}
}
