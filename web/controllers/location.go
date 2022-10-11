package controllers

import (
	"errors"

	"github.com/pappz/ha-homework/service"
	"github.com/pappz/ha-homework/web/handler"
)

var (
	errInvalidCoordinate = errors.New("invalid coordinate")
	errInvalidVelocity   = errors.New("invalid velocity")
)

type LocationRequest struct {
	X   float64 `json:",string"`
	Y   float64 `json:",string"`
	Z   float64 `json:",string"`
	Vel float64 `json:",string"`
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
	Location float64 `json:"loc,string"`
}

type Location struct {
}

func (h Location) RequestDataType() interface{} {
	return LocationRequest{}
}

func (h Location) Do(ri handler.RequestInfo) (handler.ResponseData, error) {
	rd := ri.Data.(LocationRequest)
	dd := service.DroneData{
		X:        rd.X,
		Y:        rd.Y,
		Z:        rd.Z,
		Velocity: rd.Vel,
	}
	resp := LocationResponse{
		ri.Service.Location(dd),
	}
	return resp, nil
}
