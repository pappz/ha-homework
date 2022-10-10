package sector

type Sector struct {
	id float64
}

func NewSector(id int) Sector {
	return Sector{
		float64(id),
	}
}

func (s Sector) Location(drone DroneData) float64 {
	return drone.x*s.id + drone.y*s.id + drone.z*s.id + drone.velocity
}
