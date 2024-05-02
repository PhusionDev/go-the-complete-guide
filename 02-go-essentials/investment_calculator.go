package main

import (
	"fmt"
)

func calc() {
	revenue := getUserInput("Enter your revenue: ")
	expenses := getUserInput("Enter your expenses: ")
	taxRate := getUserInput("Enter your tax rate: ")

	ebt, profit, ratio := CalculateFinancials(revenue, expenses, taxRate)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", ebt)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.1f\n", profit)

	fmt.Print(formattedFV, formattedRFV, fmt.Sprintf("Ratio: %.3f\n", ratio))
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func CalculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	var earningsBeforeTax float64 = revenue - expenses
	var earningsAfterTax float64 = revenue * (1 - taxRate/100) - expenses
	var ratio = earningsBeforeTax / earningsAfterTax

	return earningsBeforeTax, earningsAfterTax, ratio
}
