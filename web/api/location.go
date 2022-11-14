package api

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pappz/ha-homework/service"
	"github.com/pappz/ha-homework/web/middleware"
)

var (
	errInvalidCoordinate = errors.New("invalid coordinate")
	errInvalidVelocity   = errors.New("invalid velocity")
)

// RegisterLocationHandler sets up the routing of the HTTP handlers.
func RegisterLocationHandler(router *mux.Router, service service.Sector) {
	l := location{
		service: service,
	}
	router.HandleFunc("/databank", middleware.Handle(l.handle, l.jsonRequest)).Methods(http.MethodPost)
}

// LocationRequest input parameters from the drones
type LocationRequest struct {
	X   *float64 `json:"x,string"`
	Y   *float64 `json:"y,string"`
	Z   *float64 `json:"z,string"`
	Vel *float64 `json:"vel,string"`
}

func (rd LocationRequest) Validate() error {
	if rd.X == nil || *rd.X < 0 {
		return errInvalidCoordinate
	}
	if rd.Y == nil || *rd.Y < 0 {
		return errInvalidCoordinate
	}
	if rd.Z == nil || *rd.Z < 0 {
		return errInvalidCoordinate
	}
	if rd.Vel == nil || *rd.Vel <= 0 {
		return errInvalidVelocity
	}
	return nil
}

// LocationResponse to the request
type LocationResponse struct {
	Location float64 `json:"loc"`
}

// location is the http controller for the location of databank
type location struct {
	service service.Sector
}

func (h location) handle(ri *middleware.RequestInfo) (middleware.ResponseData, error) {
	rd := ri.Payload.(*LocationRequest)
	dd := service.DroneData{
		X:        *rd.X,
		Y:        *rd.Y,
		Z:        *rd.Z,
		Velocity: *rd.Vel,
	}

	loc := h.service.Location(dd)
	resp := LocationResponse{
		loc,
	}
	return resp, nil
}

func (h location) jsonRequest() middleware.Json {
	return &LocationRequest{}
}
