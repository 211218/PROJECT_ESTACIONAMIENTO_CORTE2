package main

import (
	"estacionamiento/scenes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Estacionamiento")
	myWindow.CenterOnScreen()
	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(1080, 720))

	mainMenuScene := scenes.NewMainMenuScene(myWindow)
	mainMenuScene.Show()
	myWindow.ShowAndRun()
}