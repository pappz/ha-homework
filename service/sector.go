package service

type Sector struct {
	id float64
}

func NewSector(id int) Sector {
	return Sector{
		float64(id),
	}
}

// Location calculate location of databank. The float will be truncated.
func (s Sector) Location(drone DroneData) float64 {
	num := drone.X*s.id + drone.Y*s.id + drone.Z*s.id + drone.Velocity
	return float64(int(num*100)) / 100
}
