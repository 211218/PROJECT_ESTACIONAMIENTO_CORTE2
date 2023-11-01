package models

import (
    "fyne.io/fyne/v2/canvas"
    "sync"
)

type Space struct {
    ID     int
    IsFree bool
    Color  string
}

type Parking struct {
    BackgroundImage  *canvas.Image
    Spaces           []Space
    mutex            *sync.Mutex // Semáforo para exclusión mutua
}

func NewParking(spaces int, image *canvas.Image, mutex *sync.Mutex) *Parking {
    parking := &Parking{
        BackgroundImage: image,
        Spaces:          make([]Space, spaces),
        mutex:           mutex, // Proporciona un semáforo para garantizar la exclusión mutua
    }
    return parking
}

func (p *Parking) Run() {
    // Utiliza el semáforo para garantizar la exclusión mutua en la sección crítica
    p.mutex.Lock()
    // Realiza las operaciones críticas aquí
    p.mutex.Unlock()
}
