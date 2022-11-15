package service

// Sector service provide the location of databank for drone
type Sector struct {
	id float64
}

// NewSector create new sector service
func NewSector(id int) Sector {
	return Sector{
		float64(id),
	}
}

// Location calculate location of databank. The float will be truncated.
func (s Sector) Location(droneData DroneData) (float64, error) {
	if err := droneData.validate(); err != nil {
		return 0, err
	}
	num := droneData.X*s.id + droneData.Y*s.id + droneData.Z*s.id + droneData.Velocity
	return float64(int(num*100)) / 100, nil
}
