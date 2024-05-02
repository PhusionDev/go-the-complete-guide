package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFile = "balance.txt"

func main() {
	accountBalance, err := getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	}

	printWelcome()

	for {
		fmt.Printf("Account Balance: %.2f\n", accountBalance)
		printMenu()
		choice := getInput()

		validInput, err := validInput(choice)
		if !validInput {
			fmt.Println(err)
			return
		}

		switch choice {
		case 1:
			fmt.Println("Check Balance")
		case 2:
			fmt.Println("Deposit Money")
			accountBalance++
		case 3:
			fmt.Println("Withdraw Money")
			accountBalance--
		case 4:
			printGoodbye()
		default:
			fmt.Println("Invalid Selection")
		}

		if (choice == 4) {
			writeBalanceToFile(accountBalance)
			return
		}
	}
}

func getInput() int {
	var choice int;
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)
	return choice
}

func printWelcome() {
		fmt.Println("Welcome to Go Bank!")
}

func printGoodbye() {
	fmt.Println("Thanks for choosing our bank")
}

func printMenu() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(balanceFile)

	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	
	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return balance, nil
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