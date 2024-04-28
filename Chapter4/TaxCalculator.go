package main

import (
	"fmt"
)

func main() {
	var excess, taxRate, count float64
	count = 1

	for {
		var earning float64
		fmt.Println("Enter citizen earnings: ")
		_, err := fmt.Scanln(&earning)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		if earning == 0 {
			break
		}

		var name string
		fmt.Println("Enter citizen names: ")
		_, err = fmt.Scanln(&name)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid name.")
			continue
		}

		taxRate = earning * (15.00 / 100.00)
		count++

		excess = earning * (20.00 / 100.00)
		count++
	}

	totalTax := taxRate + excess
	fmt.Println("The citizens' total tax:", totalTax)
}
