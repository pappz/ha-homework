package service

// Sector service provide the location of databank for drone
// TODO: write validation for the input parameters. Unfortunately
// it was not well specified.
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
