package main

import "fmt"

func printMenu() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
}

func printWelcome() {
		fmt.Println("Welcome to Go Bank!")
}

func printGoodbye() {
	fmt.Println("Thanks for choosing our bank")
}