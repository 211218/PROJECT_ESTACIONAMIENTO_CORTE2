package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
)



type MainMenuScene struct {
	window fyne.Window
}

func NewMainMenuScene(window fyne.Window) MainMenuScene {
	return MainMenuScene{
		window: window,
	}
}

func (s *MainMenuScene) Show() {

	// cargar la imagen de fondo
	backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/estacionamiento.jpg"))
	backgroundImage.Resize(fyne.NewSize(1080, 720))
	backgroundImage.Move(fyne.NewPos(0, 0))



	s.window.SetContent(container.NewWithoutLayout(backgroundImage))
}

