// simulation.go
package models

import "time"

type Simulation struct {
	Vehicles  []Vehicle
	Parking   Parking
	StartTime time.Time
}