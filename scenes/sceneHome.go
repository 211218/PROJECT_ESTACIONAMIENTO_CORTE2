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

func (s *MainMenuScene) Show() {
	backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/estacionamiento.jpg"))
	backgroundImage.Resize(fyne.NewSize(1080, 720))
	backgroundImage.Move(fyne.NewPos(0, 0))

    // Creamos un modelo de cada espacio del estacionamiento
    parkingImg := canvas.NewImageFromURI(storage.NewFileURI("./assets/background.jpg"))
    parkingImg.Resize(fyne.NewSize(60, 100))
    parkingImg.Move(fyne.NewPos(0, 0)) // Ajusta la posición según tus necesidades

	// Creamos un semáforo para exclusión mutua
	vehicleMutex := new(sync.Mutex)

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
			time.Sleep(1 * time.Second)
			s.window.Canvas().Refresh(vehicleContainer)
			createVehicle()
		}
	}()






// Crear un contenedor para los espacios de estacionamiento
spaceContainer := container.NewWithoutLayout()

// Espaciado horizontal entre los espacios
horizontalSpacing := 80 // Ajusta según tus necesidades

// Función para crear un nuevo espacio y agregarlo al contenedor
createSpace := func(x, y float32) {
    spaceImg := canvas.NewImageFromURI(storage.NewFileURI("./assets/background.jpg"))
    spaceImg.Resize(fyne.NewSize(60, 100))
    spaceImg.Move(fyne.NewPos(x, y)) // Ajusta la posición según tus necesidades
    spaceContainer.Add(spaceImg)
}

// Bucle para crear 20 espacios de estacionamiento
for i := 0; i < 20; i++ {
    var x, y float32

    if i < 10 {
        // Primeros 10 espacios (arriba)
        x = float32(i * horizontalSpacing)
        y = 250
    } else {
        // Siguientes 10 espacios (abajo)
        x = float32((i - 10) * horizontalSpacing)
        y = 370 // Ajusta la posición vertical si es necesario
    }

	// mover los espacios a la derecha
	x += 250


    createSpace(x, y)
}

// Establecer el contenedor de espacios en la posición deseada en la interfaz
content := container.NewWithoutLayout(backgroundImage, spaceContainer, vehicleContainer)
s.window.SetContent(content)



}
