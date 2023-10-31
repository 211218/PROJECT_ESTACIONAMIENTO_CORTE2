package controllers

import (
	"fmt"
	"time"
	"estacionamiento/models"
)

type SimulationController struct {
	Vehicles []*models.Vehicle
}

func NewSimulationController() *SimulationController {
	return &SimulationController{}
}

func (s *SimulationController) AddVehicle(vehicle *models.Vehicle) {
	s.Vehicles = append(s.Vehicles, vehicle)
}

func (s *SimulationController) Run() {
	for _, vehicle := range s.Vehicles {
		go func(v *models.Vehicle) {
			fmt.Println("Iniciando vehículo")
			v.Run()
		}(vehicle)
		time.Sleep(3 * time.Second)
	}
}
