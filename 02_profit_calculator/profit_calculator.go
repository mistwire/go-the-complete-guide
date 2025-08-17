package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Revenue: ")

	if err != nil {
		fmt.Println("ERROR:")
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Expenses: ")

	if err != nil {
		fmt.Println("ERROR:")
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Tax Rate (%): ")

	if err != nil {
		fmt.Println("ERROR:")
		fmt.Println(err)
		return
	}

	EBT, profit, ratio := calculate(revenue, expenses, taxRate)

	fmt.Println("Earning Before Tax:", EBT)
	fmt.Println("Profit", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)

	writeToFile(EBT, profit, ratio)
}

func writeToFile(EBT, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", EBT, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func calculate(revenue, expenses, taxRate float64) (float64, float64, float64) {
	EBT := revenue - expenses
	profit := EBT * (1 - taxRate/100)
	ratio := EBT / profit
	return EBT, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return userInput, errors.New("input cannot be 0 or negative")
	}

	return userInput, nil
}
