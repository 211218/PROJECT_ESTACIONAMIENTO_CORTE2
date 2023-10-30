package models

import (
	"sync"
	"time"
)

type Space struct {
	ID     int
	IsFree bool
	Color  string
}

type Parking struct {
	Spaces []Space
	mu     sync.Mutex
}

func NewParking(spaces int) *Parking {
	p := &Parking{}
	for i := 0; i < spaces; i++ {
		p.Spaces = append(p.Spaces, Space{ID: i, IsFree: true, Color: "Green"}) // Espacios inicialmente libres (verde)
	}
	return p
}

func (p *Parking) OccupySpace() int {
	p.mu.Lock()
	defer p.mu.Unlock() 
	for i := range p.Spaces {
		if p.Spaces[i].IsFree {
			p.Spaces[i].IsFree = false
			p.Spaces[i].Color = "Red" // Espacio ocupado (rojo)
			return i
		}
	}
	return -1 // No hay espacios disponibles
}

func (p *Parking) ReleaseSpace(spaceID int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if spaceID >= 0 && spaceID < len(p.Spaces) {
		p.Spaces[spaceID].IsFree = true
		p.Spaces[spaceID].Color = "Green" // Espacio liberado (verde)
	}
}

func (p *Parking) Run() {
	go func() {
		for {
			// Simulación de tiempo en el estacionamiento
			time.Sleep(time.Second * 2)

			// Ejemplo: Liberar espacios aleatorios después de un tiempo
			spaceIDToRelease := 1 // Cambia esto según tu lógica
			p.ReleaseSpace(spaceIDToRelease)
		}
	}()
}
