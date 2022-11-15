package service

import "errors"

var (
	ErrInvalidCoordinate = errors.New("invalid coordinate")
	ErrInvalidVelocity   = errors.New("invalid velocity")
)

// DroneData describe the required inputs for the math
type DroneData struct {
	X, Y, Z, Velocity float64
}

// TODO: write better validation. Unfortunately it was not well specified.
// validate check the user inputs
func (dd DroneData) validate() error {
	if dd.X < 0 {
		return ErrInvalidCoordinate
	}
	if dd.Y < 0 {
		return ErrInvalidCoordinate
	}
	if dd.Z < 0 {
		return ErrInvalidCoordinate
	}
	if dd.Velocity <= 0 {
		return ErrInvalidVelocity
	}
	return nil
}
