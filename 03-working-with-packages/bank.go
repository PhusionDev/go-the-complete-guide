package main

import (
	"fmt"

	"example.com/go-bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const balanceFile = "balance.txt"

func main() {
	accountBalance := initializeBank()
	accountBalance = mainLoop(accountBalance)
	exitBank(accountBalance)
}

func initializeBank() float64 {
	accountBalance, err := fileops.GetFloatFromFile(balanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	}

	printWelcome()
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())
	return accountBalance
}

func exitBank(accountBalance float64) {
	fileops.WriteFloatToFile(balanceFile, accountBalance)
	printGoodbye()
}

func mainLoop(accountBalance float64) float64 {
	for {
		fmt.Printf("Account Balance: %.2f\n", accountBalance)
		printMenu()
		choice := getInput()

		validInput, err := validInput(choice)
		if !validInput {
			fmt.Println(err)
			return accountBalance
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
			return accountBalance
		default:
			fmt.Println("Invalid Selection")
		}
	}
}