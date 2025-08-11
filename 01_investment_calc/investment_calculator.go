package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main () {
	
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate (Default is 5.5%): ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	// Outputs
	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Inflation Adjusted: %.2f\n", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
	rfv = fv / math.Pow(1 + inflationRate/100, years)
	return fv, rfv
	
}