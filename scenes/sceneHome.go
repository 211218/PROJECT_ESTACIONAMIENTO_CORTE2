package scenes

import (
    "estacionamiento/models"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/storage"
    "time"
)

type MainMenuScene struct {
    window fyne.Window
}

func NewMainMenuScene(window fyne.Window) *MainMenuScene {
    return &MainMenuScene{window: window}
}

var parking *models.Parking
var vehicles []*models.Vehicle
var simulation *models.Simulation

func (s *MainMenuScene) Show() {
    backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/estacionamiento.jpg"))
    backgroundImage.Resize(fyne.NewSize(1080, 720))
    backgroundImage.Move(fyne.NewPos(0, 0))

    // Crear la imagen del vehículo
    vehicleImg := canvas.NewImageFromURI(storage.NewFileURI("./assets/car.png"))
    vehicleImg.Resize(fyne.NewSize(60, 100))
    vehicleImg.Move(fyne.NewPos(0, 0))

    // Creamos el modelo del vehículo
    vehicle := models.NewVehicle(vehicleImg)

    // Creamos fondo para el estacionamiento
    parkingImg := canvas.NewImageFromURI(storage.NewFileURI("./assets/background.jpg"))
    parkingImg.Resize(fyne.NewSize(60, 100))
    parkingImg.Move(fyne.NewPos(0, 0))

    // Creamos el modelo del estacionamiento
    spacingX := 20        // No hay separación horizontal
    spacingY := 20       // Separación vertical de 20 píxeles
    parking = models.NewParking(20, parkingImg, spacingX, spacingY)

    // Iniciar los goroutines de los modelos
    go parking.Run()
    go vehicle.Run()

    // Crear un contenedor principal para la escena
    mainContainer := container.NewWithoutLayout(backgroundImage, parkingImg, vehicleImg)

    // Mostrar la escena
    s.window.SetContent(mainContainer)

    // Crear vehículos cada 3 segundos
    go func() {
        for {
            vehicleImgCopy := canvas.NewImageFromURI(storage.NewFileURI("./assets/car.png"))
            vehicleImgCopy.Resize(fyne.NewSize(60, 100))
            vehicleImgCopy.Move(fyne.NewPos(0, 0))
            vehicleCopy := models.NewVehicle(vehicleImgCopy)
            vehicles = append(vehicles, vehicleCopy)
            s.window.SetContent(container.NewWithoutLayout(vehicleImgCopy))
            time.Sleep(3 * time.Second)
        }
    }()
}
