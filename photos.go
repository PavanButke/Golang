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
)

func main() {
	a := app.New()
	wdw := a.NewWindow("Photos")
	wdw.Resize(fyne.NewSize(800, 600))

	root_src := "/config/Downloads/"

	files, err := ioutil.ReadDir(root_src)
	if err != nil {
		log.Fatal(err)
	}

	tabs := container.NewAppTabs()

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())

		if file.IsDir() == false {
			extension := strings.Split(file.Name(), ".")[1]

			if extension == "png" || extension == "jpeg" {
				image := canvas.NewImageFromFile(root_src + "/" + file.Name())

				tabs.Append(container.NewTabItem(file.Name(), image))
			}
		}
	}

	wdw.SetContent(tabs)

	wdw.ShowAndRun()
}
