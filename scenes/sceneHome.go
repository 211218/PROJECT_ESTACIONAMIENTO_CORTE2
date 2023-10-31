package scenes

import (
	"estacionamiento/models"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
)

type MainMenuScene struct {
    window fyne.Window
}

func NewMainMenuScene(window fyne.Window) *MainMenuScene {
    return &MainMenuScene{window: window}
}

var parking *models.Parking
var vehicle *models.Vehicle
var simulation *models.Simulation

func (s *MainMenuScene) Show() {
    backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/estacionamiento.jpg"))
    backgroundImage.Resize(fyne.NewSize(1080, 720))
    backgroundImage.Move(fyne.NewPos(0, 0))


    // Crear la imagen del vehiculo
    vehicleImg := canvas.NewImageFromURI(storage.NewFileURI("./assets/car.png"))
    vehicleImg.Resize(fyne.NewSize(60, 100))
    vehicleImg.Move(fyne.NewPos(0, 0))
    // Creamos el modelo del vehiculo
    vehicle = models.NewVehicle(vehicleImg)

    // Creamos fondo para el parking
    parkingImg := canvas.NewImageFromURI(storage.NewFileURI("./assets/background.jpg"))
    parkingImg.Resize(fyne.NewSize(60, 100))
    parkingImg.Move(fyne.NewPos(0, 0))

    // Creamos el modelo del parking
    spacingX := 20        // No hay separación horizontal
    spacingY := 20       // Separación vertical de 20 píxeles
    parking = models.NewParking(20, parkingImg, spacingX, spacingY)


    
    // Iniciar los goroutines de los modelos
    go parking.Run()
    
    for vehicleCount := 0; vehicleCount < 2; vehicleCount++ {
        go func(count int) {
            fmt.Printf("Iniciando vehículo número %d\n", count)
            go vehicle.Run()
            // Aquí va la lógica real de tu vehículo
            fmt.Printf("Vehículo número %d terminado\n", count)
        }(vehicleCount)
        time.Sleep(3 * time.Second) // Duerme durante 3 segundos
    }
    
    
	// Espera un poco para que las goroutines terminen


    // Crear un contenedor principal para la escena
    s.window.SetContent(container.NewWithoutLayout(backgroundImage, vehicleImg, parkingImg))
}