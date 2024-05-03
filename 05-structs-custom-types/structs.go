package main

import (
	"fmt"

	"example.com/structs/user"
)


type customString string

func (text *customString) log() {
	fmt.Println(*text)
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Plaese enter your birthdate (MM/DD/YYYY): ")

	var name customString = "Ham"
	name.log()

	appUser, err := user.New(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("admin@example.com", "test123")
	admin.OutputUserDetails()

	// do something with gathered data

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}