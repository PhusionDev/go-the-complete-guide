package main

import (
	"errors"
	"fmt"
)

func getInput() int {
	var choice int;
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)
	return choice
}

func validInput(choice int) (bool, error) {
	if choice == 0 {
		return false, errors.New("value cannot be 0")
	}

	if choice < 0 {
		return false, errors.New("value cannot be negative")
	}

	return true, nil
}