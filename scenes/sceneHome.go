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
    parkingImg := canvas.NewImageFromURI(storage.NewFileURI("./assets/background.jpg"))
    parkingImg.Resize(fyne.NewSize(60, 100))
    parkingImg.Move(fyne.NewPos(0, 0))

    // Creamos el modelo del parking
    spacingX := 20        // No hay separación horizontal
    spacingY := 20       // Separación vertical de 20 píxeles
    parking = models.NewParking(20, parkingImg, spacingX, spacingY, vehicleMutex)

    // Iniciar el modelo del parking
    go parking.Run()

    // Crear un contenedor principal para la escena
    // Crear un contenedor para los vehículos
    vehicleContainer := container.NewWithoutLayout()

    // Crear 100 vehículos
    for i := 0; i < 5; i++ {
        vehicleImg := canvas.NewImageFromURI(storage.NewFileURI("./assets/car.png"))
        vehicleImg.Resize(fyne.NewSize(60, 100))
        vehicleImg.Move(fyne.NewPos(0, float32(i*110))) // Asegúrate de que los vehículos no se superpongan
        go func() {
            vehicle := models.NewVehicle(vehicleImg, vehicleMutex)
            vehicle.Run()
            time.Sleep(time.Duration(100) * time.Millisecond)
        }()
        vehicleContainer.Add(vehicleImg)
    }

    // Crear 20 espacios para cada vehiculo usando el modelo parking
    for i := 0; i < 20; i++ {
        parkingImg := parking.BackgroundImages[i]
        parkingImg.Move(fyne.NewPos(0, float32(i*110))) // Asegúrate de que los vehículos no se superpongan
        vehicleContainer.Add(parkingImg)
    }

    // Crear un contenedor principal para la escena
    content := container.NewWithoutLayout(backgroundImage, parkingImg, vehicleContainer)
    s.window.SetContent(content)
}
