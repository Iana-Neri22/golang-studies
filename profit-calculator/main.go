package main

import (
	"fmt"
)

func main() {

	revenue := getUserInput("Enter total revenue: ")
	expenses := getUserInput("Enter total expenses: ")
	taxRate := getUserInput("Enter tax rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	formattedRatio := fmt.Sprintf("%.2f%%", ratio)

	fmt.Println("Earnings Before Tax: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("Profit Ratio: ", formattedRatio)
}

func getUserInput(prompt string) float64 {
	var userInput float64
	fmt.Print(prompt)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}