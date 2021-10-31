package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
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

	img := canvas.NewImageFromFile("/config/Desktop/Golang/OS/e5wMGQy.webp")

	btn1 = widget.NewButtonWindowIcon("Weather App", theme.InfoIcon(), func() {

		showWeatherApp(myWindow)
	})

	btn2 = widget.NewButtonWindowIcon("Calculator ", theme.ComputerIcon(), func() {

		showCalc()
	})

}
