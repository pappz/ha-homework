package service

type Sector struct {
	id float64
}

func NewSector(id int) Sector {
	return Sector{
		float64(id),
	}
}

func (s Sector) Location(drone DroneData) float64 {
	return drone.X*s.id + drone.Y*s.id + drone.Z*s.id + drone.Velocity
}
