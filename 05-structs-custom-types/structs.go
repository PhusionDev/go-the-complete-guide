package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Plaese enter your birthdate (MM/DD/YYYY): ")

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