package main

import (
	"encoding/json"

	"fmt"
	"io/ioutil"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	// new title and window
	w := a.NewWindow("Contacts")
	// resize window
	w.Resize(fyne.NewSize(400, 400))

	type Student struct {
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Birthday     string `json:"birthday"`
		PhoneNumber  string `json:"phone_number"`
		Email        string `json:"email"`
		Organization string `json:"organization"`
	}

	var myStudentData []Student

	jsonFile, err := os.Open("Student.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened Student.json")

	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var arr []Student

	json.Unmarshal(byteValue, &arr)

	l_name := widget.NewLabel("...")
	l_name.TextStyle = fyne.TextStyle{Bold: true}
	// for phone number
	l_phone := widget.NewLabel("...")
	// entry widget for name
	e_fname := widget.NewEntry()
	e_fname.SetPlaceHolder("Enter name here...")
	e_lname := widget.NewEntry()
	e_lname.SetPlaceHolder("Enter name here...")
	e_mail := widget.NewEntry()
	e_lname.SetPlaceHolder("Enter Email here...")
	// entry widget for phone
	e_phone := widget.NewEntry()
	e_phone.SetPlaceHolder("Enter phone here...")

	e_birth := widget.NewEntry()
	e_birth.SetPlaceHolder("Enter Birthdate here(Optional)")

	e_org := widget.NewEntry()
	e_org.SetPlaceHolder("Enter Organization Name here(Optional)...")
	// submit button
	submit_btn := widget.NewButton("Submit", func() {
		// logic part- store our data in json format
		// let create a struct for our data
		// Get data from entry widget and push to slice
		myData1 := &Student{
			FirstName:    e_fname.Text,
			LastName:     e_lname.Text, // data from name entry widget
			PhoneNumber:  e_phone.Text,
			Email:        e_mail.Text,
			Birthday:     e_birth.Text,
			Organization: e_org.Text,
		}
		// append / push data to our slice
		myStudentData = append(myStudentData, *myData1)
		// * star is very important
		// convert / parse data to json format
		// 3 arguments
		// 1st is our slice data
		// ignore 2nd
		// 3rd is identation, we are using space to indent our data
		final_data, _ := json.MarshalIndent(myStudentData, "", " ")
		// writing data to file
		// it take 3 things
		//file name .txt or .json or anything you want to use
		// data source, final_data in our case
		// writing/reading/execute permission
		ioutil.WriteFile("Student.json", final_data, 0644)
		/// we are done 🙂
		// lets test
		// empty input field after data insertion
		e_fname.Text = ""
		e_lname.Text = ""
		e_phone.Text = ""
		e_mail.Text = ""
		e_birth.Text = ""
		e_org.Text = ""

		// refresh name & phone entry object
		e_fname.Refresh()
		e_lname.Text = ""
		e_lname.Refresh()

		e_phone.Refresh()
		e_mail.Refresh()
		e_birth.Refresh()
		e_org.Refresh()

	})
	// list widget
	list := widget.NewList(
		// first argument is item count
		// len() function to get myStudentData slice len
		func() int { return len(myStudentData) },
		// 2nd argument is for widget choice. I want to use label
		func() fyne.CanvasObject { return widget.NewLabel("") },
		// 3rd argument is to update widget with our new data
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(myStudentData[lii].FirstName)
		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		l_name.Text = myStudentData[id].LastName
		l_name.Refresh()
		// for phone number
		l_phone.Text = myStudentData[id].PhoneNumber
		l_phone.Refresh()

		call := widget.NewButtonWithIcon("Call Now", theme.VolumeUpIcon(), func() {
			image := canvas.NewImageFromFile("call.jpg")
			image.FillMode = canvas.ImageFillContain
		})

		message := widget.NewButtonWithIcon("Write a Mail", theme.MailComposeIcon(), func() {

		})

		w.SetContent(container.NewVBox(l_name, l_phone, call, message))

	}

	// show and run
	w.SetContent(
		// lets create Hsplite container
		container.NewHSplit(
			// first argument is list of data
			list,
			// 2nd is
			// vbox container
			// show both label
			container.NewVBox(l_name, l_phone, e_fname, e_lname, e_phone, submit_btn),
		),
	)
	w.ShowAndRun()

}
