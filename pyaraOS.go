package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

var myApp fyne.App = app.New()
var myWindow fyne.Window = myApp.NewWindow("Windows ki Chutti")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var img fyne.CanvasObject
var DeskBtn fyne.Widget

func main() {

	myApp.Settings().SetTheme(theme.LightTheme())

	img := canvas.NewImageFromFile()

}
