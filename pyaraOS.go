package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App = app.New()
var myWindow fyne.Window = myApp.NewWindow("Windows ki Chutti")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var img fyne.CanvasObject
var panelContent *fyne.Container
var DeskBtn fyne.Widget

func main() {

	myApp.Settings().SetTheme(theme.LightTheme())

	img := canvas.NewImageFromFile("/config/Desktop/Golang/OS/pexels-eberhard-grossgasteiger-844297.jpg")
	img.FillMode = canvas.ImageFillContain

	btn1 = widget.NewButtonWithIcon("Weather App", theme.InfoIcon(), func() {

		showWeatherApp(myWindow)
	})

	btn2 = widget.NewButtonWithIcon("Calculator ", theme.ComputerIcon(), func() {

		showCalcApp()
	})

	btn3 = widget.NewButtonWithIcon("Gallery ", theme.MediaPhotoIcon(), func() {

		showGalleryApp(myWindow)
	})

	DeskBtn = widget.NewButtonWithIcon("Home ", theme.HomeIcon(), func() {

		myWindow.SetContent(container.NewBorder(panelContent, nil, nil, nil, img))
	})
	panelContent = container.NewVBox((container.NewGridWithColumns(4, DeskBtn, btn1, btn2, btn3)))

	myWindow.Resize(fyne.NewSize(1200, 600))
	myWindow.CenterOnScreen()

	myWindow.SetContent(container.NewBorder(panelContent, nil, nil, img))

	myWindow.ShowAndRun()
}
