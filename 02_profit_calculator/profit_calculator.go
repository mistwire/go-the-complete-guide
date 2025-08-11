package main

import (
	"fmt"
)

func main () {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = getUserInput("Revenue: ")
	expenses = getUserInput("Expenses: ")
	taxRate = getUserInput("Tax Rate (%): ")

	EBT, profit, ratio := calculate(revenue, expenses, taxRate)
	fmt.Println("Earning Before Tax:", EBT)
	fmt.Println("Profit", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

func getUserInput (infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput 
}


func calculate (revenueAmount, expenses, taxRate float64) (float64, float64, float64) {
	EBT := revenueAmount - expenses
	profit := EBT * (1 - taxRate/100)
	ratio := EBT/profit
	return EBT, profit, ratio 
}