package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5


func main() {
	var investmentAmount float64
	var years float64  
	expectedReturnRate := 5.5

	outputText("Investment Amount:")
	//fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected Return Rate:")
	//fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Years:")
	//fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	//futureValue := investmentAmount * math.Pow(1 + expectedReturnRate/100, years) 
	//futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f", futureValue)
	formattedFRV := fmt.Sprintf("Future Real Value: %.2f", futureRealValue)

	fmt.Println(formattedFV)
	fmt.Println(formattedFRV)
}

func outputText(text string) {
	fmt.Println(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
	rfv = fv / math.Pow(1 + inflationRate/100, years)
	return fv, rfv
}
	