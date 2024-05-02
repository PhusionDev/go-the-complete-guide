package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5

	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow((1 + expectedReturnRate / 100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate / 100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

	CalculuateProfit()
}

func CalculuateProfit() {
	var revenue, expenses, taxRate float64;

	fmt.Print("Enter your revenue: ")
	fmt.Scan(&revenue)
	
	fmt.Print("Enter your expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter your tax rate: ")
	fmt.Scan(&taxRate)

	var earningsBeforeTax float64 = revenue - expenses
	var earningsAfterTax float64 = revenue * (1 - taxRate/100) - expenses
	var ratio = earningsBeforeTax / earningsAfterTax

	fmt.Println("Earnings before tax:", earningsBeforeTax)
	fmt.Println("Profit:", earningsAfterTax)
	fmt.Println("Ratio:", ratio)
}