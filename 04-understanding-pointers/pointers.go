package main

import "fmt"

func main() {
	age := 32 // regular variable
	agePtr := &age

	fmt.Println("Age:", *agePtr)
	convertToAdultYears(agePtr)
	fmt.Println("Adult Years:", age)
}

func convertToAdultYears(age *int) {
	*age = *age - 18
}