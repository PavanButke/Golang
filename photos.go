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

	var picsArr []string
	picsArr = make([]string, 5, 8)

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())

		if !file.IsDir() {
			extension := strings.Split(file.Name(), ".")[1]
			if extension == "png" || extension == "jpeg" {
				picsArr = append(picsArr, root_src+"\\"+file.Name())
			}
		}
	}

	Tabs := container.NewAppTabs(
		container.NewTabItem("Image1", canvas.NewImageFromFile(picsArr[3])),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)
	wdw.SetContent(Tabs)

	wdw.ShowAndRun()
}
