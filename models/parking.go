package models

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "sync"
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
    mutex            *sync.Mutex // Semáforo para exclusión mutua
}

func NewParking(spaces int, image *canvas.Image, spacingX, spacingY int, mutex *sync.Mutex) *Parking {
    parking := &Parking{
        background: image,
        Spaces:     make([]Space, spaces),
        mutex:      mutex, // Proporciona un semáforo para garantizar la exclusión mutua
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
    // Utiliza el semáforo para garantizar la exclusión mutua en la sección crítica
    p.mutex.Lock()
    // Realiza las operaciones críticas aquí
    p.mutex.Unlock()
}
