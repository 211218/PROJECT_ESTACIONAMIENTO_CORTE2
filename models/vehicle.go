package models

import (
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
		for {
			v.image.Move(fyne.NewPos(v.image.Position().X+10, v.image.Position().Y))
			canvas.Refresh(v.image)
			time.Sleep(time.Millisecond * 500)
		}
}

