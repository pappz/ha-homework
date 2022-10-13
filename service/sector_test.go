package service

import "testing"

func TestSector_Location(t *testing.T) {

	tests := []struct {
		sectorId  int
		droneData DroneData
		want      float64
	}{
		{
			1,
			DroneData{
				123.12,
				456.56,
				789.89,
				20,
			},
			1389.57,
		},
		{
			2,
			DroneData{
				123.12,
				456.56,
				789.89,
				20,
			},
			2759.14,
		},
		{
			1,
			DroneData{
				1,
				1,
				1,
				1,
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run("location", func(t *testing.T) {
			s := NewSector(tt.sectorId)
			if got := s.Location(tt.droneData); got != tt.want {
				t.Errorf("Location() = %v, want %v", got, tt.want)
			}
		})
	}
}
