package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
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
	var picsArr []string

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())

		if !file.IsDir() {
			extension := strings.Split(file.Name(), ".")[1]
			if extension == "png" || extension == "jpeg" {
				picsArr = append(picsArr, root_src+"\\"+file.Name())
			}
		}
	}

	image := canvas.NewImageFromFile(picsArr[0])
	wdw.SetContent(image)

	wdw.ShowAndRun()
}
