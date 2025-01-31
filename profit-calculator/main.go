package main

import (
	"fmt"
	"math"
)

func main() {

	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter total revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter total expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println("Earnings Before Tax: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("Profit Ratio: ", math.Round(ratio), "%")
}