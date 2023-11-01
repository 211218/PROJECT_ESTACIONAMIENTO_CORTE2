package scenes

import (
	"estacionamiento/models"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"sync"
	"time"
)

type MainMenuScene struct {
	window fyne.Window
}

func NewMainMenuScene(window fyne.Window) *MainMenuScene {
	return &MainMenuScene{window: window}
}

var parking *models.Parking
var simulation *models.Simulation

func (s *MainMenuScene) Show() {
	backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/estacionamiento.jpg"))
	backgroundImage.Resize(fyne.NewSize(1080, 720))
	backgroundImage.Move(fyne.NewPos(0, 0))

	// Creamos un semáforo para exclusión mutua
	vehicleMutex := new(sync.Mutex)

	// Creamos fondo para el parking
	parkingBackgroundImg := canvas.NewImageFromURI(storage.NewFileURI("./assets/background.jpg"))
	parkingBackgroundImg.Resize(fyne.NewSize(60, 100))
	parkingBackgroundImg.Move(fyne.NewPos(0, 0))

	// Creamos el modelo del parking
	spacingX := 20 // No hay separación horizontal
	spacingY := 20 // Separación vertical de 20 píxeles
	parking = models.NewParking(20, parkingBackgroundImg, spacingX, spacingY, vehicleMutex)

	// Iniciar el modelo del parking
	go parking.Run()

	// Crear un contenedor para los vehículos
	vehicleContainer := container.NewWithoutLayout()

	// Creamos un canal para sincronizar la creación de vehículos
	vehicleCreationTicker := time.NewTicker(3 * time.Second)
	defer vehicleCreationTicker.Stop()

	// Función para crear un nuevo vehículo y agregarlo al contenedor
	createVehicle := func() {
		vehicleImg := canvas.NewImageFromURI(storage.NewFileURI("./assets/car.png"))
		vehicleImg.Resize(fyne.NewSize(60, 100))
		vehicleImg.Move(fyne.NewPos(0, 0)) // Ajusta la posición según tus necesidades

		go func() {
			vehicle := models.NewVehicle(vehicleImg, vehicleMutex)
			vehicle.Run()
		}()

		vehicleContainer.Add(vehicleImg)
	}

	// Bucle para crear vehículos cada 3 segundos
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(3 * time.Second)
			s.window.Canvas().Refresh(vehicleContainer)
			createVehicle()
		}
	}()

	// Crear 20 espacios para cada vehiculo usando el modelo parking
	for i := 0; i < 20; i++ {
		parkingImg := parking.BackgroundImages[i]
		parkingImg.Move(fyne.NewPos(0, float32(i*110))) // Asegúrate de que los vehículos no se superpongan
		vehicleContainer.Add(parkingImg)
	}

	// Crear un contenedor principal para la escena
	content := container.NewWithoutLayout(backgroundImage, vehicleContainer)
	s.window.SetContent(content)
}
