package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Space struct {
	ID     int
	IsFree bool
	Color  string
}

type Parking struct {
	background        *canvas.Image
	Spaces           []Space
	BackgroundImages []*canvas.Image
}

func NewParking(spaces int, image *canvas.Image, spacingX, spacingY int) *Parking {
	parking := &Parking{
		background: image,
		Spaces:     make([]Space, spaces),
	}

	for i := 0; i < spaces; i++ {
		// Crear una copia de la imagen de fondo
		imageCopy := canvas.NewImageFromImage(image.Image)

		// Ajustar la posición de la copia para separarlas
		imageCopy.Move(fyne.NewPos(float32(i)*float32(spacingX), float32(i)*float32(spacingY)))

		// Almacenar la copia en BackgroundImages
		parking.BackgroundImages = append(parking.BackgroundImages, imageCopy)
	}

	return parking
}

func (p *Parking) Run() {
	
}
