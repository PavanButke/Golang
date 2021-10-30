package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	wdw := a.NewWindow("Photos")
	wdw.Resize(fyne.NewSize(800, 600))

	root_src := "/config/Desktop/Golang/Gallery/"

	files, err := ioutil.ReadDir(root_src)
	if err != nil {
		log.Fatal(err)
	}

	Tabs := container.NewAppTabs()


	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())

		if !file.IsDir() {
			extension := strings.Split(file.Name(), ".")[1]
			image.
			if extension == "png" || extension == "jpeg" {
				image := canvas.NewImageFromFile( root_src+"/"+file.Name())
				image.FillMode= canvas.ImageFillOriginal
				Tabs.Append(container.NewTabItem(file.Name(),image))
			}
		}
	}
	Tabs.SetTabLocation(container.TabLocationLeading)

	wdw.SetContent(Tabs)

	wdw.ShowAndRun()
}