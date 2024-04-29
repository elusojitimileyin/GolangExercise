package main

import (
	"fmt"
	"math"
)

func main() {
	var count int
	fmt.Print("Enter the number of values: ")
	fmt.Scanln(&count)

	if count <= 0 {
		fmt.Println("Invalid input. Number of values must be greater than 0.")
		return
	}

	min, max := math.MaxInt64, math.MinInt64
	var sum int

	for index := 1; index <= count; index++ {
		var num int
		fmt.Printf("Enter value %d: ", index)
		fmt.Scanln(&num)

		if num < min {
			min = num
		}
		if num > max {
			max = num
		}

		sum += num
	}

	fmt.Println("Minimum value:", min)
	fmt.Println("Maximum value:", max)
	fmt.Println("Sum of value:", min+max)
}
