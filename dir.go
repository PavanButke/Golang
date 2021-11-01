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

	w := a.NewWindow("Contacts")

	w.Resize(fyne.NewSize(800, 400))

	type Contancts struct {
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Birthday     string `json:"birthday"`
		PhoneNumber  string `json:"phone_number"`
		Email        string `json:"email"`
		Organization string `json:"organization"`
	}

	var myContacts []Contancts

	jsonFile, err := os.Open("Contacts.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened Contacts.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &myContacts)

	l_name := widget.NewLabel("...")

	l_name.TextStyle = fyne.TextStyle{Bold: true}

	l_phone := widget.NewLabel("...")

	l_org := widget.NewLabel("...")

	e_fname := widget.NewEntry()
	e_fname.SetPlaceHolder("Enter Firstname here...")
	e_lname := widget.NewEntry()
	e_lname.SetPlaceHolder("Enter Lastname here...")
	e_birth := widget.NewEntry()
	e_birth.SetPlaceHolder("Enter Birthdate here(Optional)")
	e_phone := widget.NewEntry()
	e_phone.SetPlaceHolder("Enter Phone here...")
	e_mail := widget.NewEntry()
	e_lname.SetPlaceHolder("Enter Email here...")

	e_org := widget.NewEntry()
	e_org.SetPlaceHolder("Enter Organization Name here(Optional)...")

	submit_btn := widget.NewButton("Submit", func() {

		myData1 := &Contancts{
			FirstName:    e_fname.Text,
			LastName:     e_lname.Text,
			Birthday:     e_birth.Text,
			PhoneNumber:  e_phone.Text,
			Email:        e_mail.Text,
			Organization: e_org.Text,
		}

		myContacts = append(myContacts, *myData1)

		final_data, _ := json.MarshalIndent(myContacts, "", " ")

		ioutil.WriteFile("Contacts.json", final_data, 0644)

		e_fname.Text = ""
		e_lname.Text = ""
		e_birth.Text = ""
		e_phone.Text = ""
		e_mail.Text = ""
		e_org.Text = ""

		e_fname.Refresh()
		e_lname.Refresh()
		e_birth.Refresh()
		e_phone.Refresh()
		e_mail.Refresh()
		e_org.Refresh()

	})

	list := widget.NewList(

		func() int { return len(myContacts) },

		func() fyne.CanvasObject { return widget.NewLabel("") },

		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(myContacts[lii].FirstName)
		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		l_name.Text = myContacts[id].LastName
		l_name.Refresh()

		l_phone.Text = myContacts[id].PhoneNumber
		l_phone.Refresh()
		l_org.Text = myContacts[id].Organization
		l_org.Refresh()

		call := widget.NewButtonWithIcon("Call Now", theme.VolumeUpIcon(), func() {
			image := canvas.NewImageFromFile("call.jpg")
			image.FillMode = canvas.ImageFillContain
		})

		message := widget.NewButtonWithIcon("Write a Mail", theme.MailComposeIcon(), func() {
			showText()
		})

		w.SetContent(container.NewVBox(l_name, l_phone, l_org, call, message))

	}

	w.SetContent(

		container.NewHSplit(

			list,

			container.NewVBox(e_fname, e_lname, e_phone, e_mail, e_birth, e_org, submit_btn),
		),
	)
	w.ShowAndRun()

}
