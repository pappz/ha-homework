package service

import "testing"

func TestDroneData_validate(t *testing.T) {

	tests := []struct {
		name      string
		droneData DroneData
		wantErr   bool
	}{
		{"valid",
			DroneData{
				1,
				2,
				3,
				4,
			},
			false,
		},
		{"invalid x",
			DroneData{
				-1,
				2,
				3,
				4,
			},
			true,
		},
		{"invalid y",
			DroneData{
				1,
				-2,
				3,
				4,
			},
			true,
		},
		{"invalid z",
			DroneData{
				1,
				2,
				-3,
				4,
			},
			true,
		},
		{"invalid vel",
			DroneData{
				1,
				2,
				3,
				-4,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.droneData.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
