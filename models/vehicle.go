package models

import (
	"math/rand"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Vehicle struct {
    image     *canvas.Image
    mutex     *sync.Mutex // Semáforo para exclusión mutua
}

func NewVehicle(image *canvas.Image, mutex *sync.Mutex) *Vehicle {
    return &Vehicle{image: image, mutex: mutex}
}

func (v *Vehicle) Run() {
    // Simulación de tiempo aleatorio de espera antes de mover el vehículo
    time.Sleep(time.Duration(100+rand.Intn(400)) * time.Millisecond)

    v.mutex.Lock() // Bloquear exclusión mutua antes de mover el vehículo
    for i := 0; i < 5; i++ {
        v.image.Move(fyne.NewPos(v.image.Position().X+10, v.image.Position().Y))
        canvas.Refresh(v.image)
        time.Sleep(time.Millisecond * 500)
    }
    v.mutex.Unlock() // Desbloquear exclusión mutua después de mover el vehículo
}
   