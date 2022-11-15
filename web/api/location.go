package api

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/pappz/ha-homework/service"
	"github.com/pappz/ha-homework/web/middleware"
)

var (
	errMissingMandatoryField = errors.New("missing field(s)")
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
	if rd.X == nil {
		return errMissingMandatoryField
	}
	if rd.Y == nil {
		return errMissingMandatoryField
	}
	if rd.Z == nil {
		return errMissingMandatoryField
	}
	if rd.Vel == nil {
		return errMissingMandatoryField
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

	loc, err := h.service.Location(dd)
	if err != nil {
		return h.handleError(err)
	}

	resp := LocationResponse{
		loc,
	}
	return resp, nil
}

func (h location) jsonRequest() interface{} {
	return &LocationRequest{}
}

// handleError hide the internal error from user
func (h location) handleError(err error) (middleware.ResponseData, error) {
	if err == service.ErrInvalidCoordinate || err == service.ErrInvalidVelocity {
		return nil, err
	}

	log.Errorf("unexpected error: %s", err.Error())
	return nil, middleware.ErrRespInternalError
}
