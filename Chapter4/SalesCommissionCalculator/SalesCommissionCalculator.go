package main

import (
	"fmt"
)

func main() {
	var items, value int
	items = 1
	for items <= 4 {
		fmt.Println("Enter one salesperson's items sold for last week: ")
		fmt.Scanln(&value)
		items++
	}

	earnings := float64(value) * 0.9

	fmt.Println("Earnings =", earnings)
}
