package models
import (

	"fyne.io/fyne/v2/canvas"
)
type Parking struct {
    image  *canvas.Image
    numSpaces int
	X, Y   float32
	SpeedX float32 // Velocidad de desplazamiento en el eje X
}


func NewParking(numSpaces int, image *canvas.Image) *Parking {
    return &Parking{numSpaces: numSpaces, image: image}
}

func (p *Parking) Run() {

}
