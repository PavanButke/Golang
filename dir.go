package main

import (
	"encoding/json"

	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Open our jsonFile

	// myapp := app.New()
	// wdw := myapp.NewWindow("Weather")

	jsonFile, err := os.Open("users.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")

	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var users Users

	json.Unmarshal(byteValue, &users)
	for i := 0; i < len(users.Users); i++ {
		fmt.Println("First Name: " + users.Users[i].FirstName)
		fmt.Println("Last Name: " + users.Users[i].LastName)
		fmt.Println("Birthday: " + users.Users[i].Birthday)
		fmt.Println("PhoneNumber: " + users.Users[i].PhoneNumber)
		fmt.Println("Email: " + users.Users[i].Email)
		fmt.Println("Org: " + users.Users[i].Organization)

	}
	// wdw.ShowAndRun()

}

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Birthday     string `json:"birthday"`
	PhoneNumber  string `json:"phone_number"`
	Email        string `json:"email"`
	Organization string `json:"organization"`
}
