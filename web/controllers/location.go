package controllers

import (
	"errors"
	"fmt"

	"github.com/pappz/ha-homework/service"
	"github.com/pappz/ha-homework/web/middleware"
)

var (
	errInvalidCoordinate = errors.New("invalid coordinate")
	errInvalidVelocity   = errors.New("invalid velocity")
)

type LocationRequest struct {
	X   float64 `json:"x,string"`
	Y   float64 `json:"y,string"`
	Z   float64 `json:"z,string"`
	Vel float64 `json:"vel,string"`
}

func (rd LocationRequest) Validate() error {
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

type LocationResponse struct {
	Location string `json:"loc"`
}

type Location struct {
}

func (h Location) RequestDataType() middleware.Json {
	return &LocationRequest{}
}

func (h Location) Do(ri middleware.RequestInfo) (middleware.ResponseData, error) {
	rd := ri.Data.(*LocationRequest)
	dd := service.DroneData{
		X:        rd.X,
		Y:        rd.Y,
		Z:        rd.Z,
		Velocity: rd.Vel,
	}

	loc := ri.Service.Location(dd)
	resp := LocationResponse{
		h.formatFloat(loc),
	}
	return resp, nil
}

func (h Location) formatFloat(f float64) string {
	return fmt.Sprintf("%.2f", f)
}
