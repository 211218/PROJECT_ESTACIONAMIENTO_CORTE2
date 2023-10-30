package scenes

import (
    "estacionamiento/models"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/storage"
    "image/color"
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

// Función para traducir el color de Space a color.Color
func getSpaceColor(space *models.Space) color.Color {
    if space.IsFree {
        return color.RGBA{0, 255, 0, 255} // Verde para espacios libres
    }
    return color.RGBA{255, 0, 0, 255} // Rojo para espacios ocupados
}

func (s *MainMenuScene) Show() {
    backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/estacionamiento.jpg"))
    backgroundImage.Resize(fyne.NewSize(1080, 720))
    backgroundImage.Move(fyne.NewPos(0, 0))

    // Crear el modelo de parking con 10 espacios
    parking = models.NewParking(10)

    // Crear la imagen del vehiculo
    vehicleImg := canvas.NewImageFromURI(storage.NewFileURI("./assets/car.png"))
    vehicleImg.Resize(fyne.NewSize(60, 100))
    vehicleImg.Move(fyne.NewPos(0, 0))
    // Creamos el modelo del vehiculo
    vehicle = models.NewVehicle(vehicleImg)

    // Crear un contenedor para mostrar los espacios del estacionamiento
spacesContainer := container.NewGridWithColumns(20)
for _, space := range parking.Spaces {
	spaceWidth := 110 // Cambia el ancho según tus preferencias
	spaceHeight := 110 // Cambia la altura según tus preferencias
	cell := canvas.NewRectangle(getSpaceColor(&space))
	cell.Resize(fyne.NewSize(float32(spaceWidth), float32(spaceHeight)))
	spacesContainer.Add(cell)
}

// Iniciar los goroutines de los modelos
go parking.Run()
go vehicle.Run()

// Crear un contenedor principal para la escena
s.window.SetContent(container.NewWithoutLayout(backgroundImage, spacesContainer, vehicleImg))
}
