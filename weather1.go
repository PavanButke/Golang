package main

import (
	"encoding/json"
	"fmt"

	"image/color"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func showWeatherApp(wdw fyne.Window) {
	// myapp := app.New()
	// wdw := myapp.NewWindow("Weather")
	// wdw.Resize(fyne.NewSize(500, 359))

	res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=mumbai&APPID=c8213346839a3bf450747bbb19415d9e")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	img1 := canvas.NewImageFromFile("index.jpg")
	img1.FillMode = canvas.ImageFillOriginal

	weather, err := UnmarshalWeather(body)
	if err != nil {
		fmt.Println(err)
	}

	label1 := canvas.NewText("Mausam ", color.Opaque)

	label1.TextStyle = fyne.TextStyle{Bold: true}

	label2 := canvas.NewText(fmt.Sprintf("Des %s", weather.Sys.Country), color.White)
	label3 := canvas.NewText(fmt.Sprintf("Tej Hawa %2f", weather.Wind.Speed), color.Opaque)
	label4 := canvas.NewText(fmt.Sprintf("Garmi %2f", weather.Main.Temp), color.Opaque)

	weatherContainer := container.NewVBox(
		label1,
		img1,
		label2,
		label3,
		label4,
		container.NewGridWithColumns(1),
	)

	wdw.SetContent(container.NewBorder(panelContent, nil, nil, nil, weatherContainer))
	wdw.Show()

}

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    welcome, err := UnmarshalWelcome(bytes)
//    bytes, err = welcome.Marshal()

func UnmarshalWeather(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int64     `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int64     `json:"timezone"`
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Cod        int64     `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
}

type Sys struct {
	Type    int64  `json:"type"`
	ID      int64  `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type Weather struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
}
