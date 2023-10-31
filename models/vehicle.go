package models

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"time"
)

type Vehicle struct {
	image *canvas.Image
}

func NewVehicle(image *canvas.Image) *Vehicle {
	return &Vehicle{image: image}
}

func (v *Vehicle) Run() {

	// crear un ciclo for para crear mas vehiculos y que ala vez se pueda ejecutar la interfaz sin delay
	for i := 0; i < 1; i++ {
		fmt.Println("Iniciando vehículo")
		canvas.Refresh(v.image)
		time.Sleep(time.Millisecond * 500)
		canvas.Refresh(v.image)
	}



		for {
			v.image.Move(fyne.NewPos(v.image.Position().X+10, v.image.Position().Y))
			canvas.Refresh(v.image)
			time.Sleep(time.Millisecond * 500)
		}
	}

