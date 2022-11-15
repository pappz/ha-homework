package service

import "testing"

func TestSector_Location(t *testing.T) {

	tests := []struct {
		sectorId  int
		droneData DroneData
		want      float64
		wantErr   bool
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
			false,
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
			false,
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
			false,
		},
		{
			1,
			DroneData{
				-1,
				-1,
				-1,
				-1,
			},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run("location", func(t *testing.T) {
			s := NewSector(tt.sectorId)
			got, err := s.Location(tt.droneData)
			if tt.wantErr && err == nil {
				t.Errorf("expected err but got nil")
			}

			if !tt.wantErr && err != nil {
				t.Errorf("expected nil err but got err: %s", err)
			}

			if got != tt.want {
				t.Errorf("Location() = %v, want %v", got, tt.want)
			}
		})
	}
}
