package main

import (
	"fmt"
)

func main() {
	var baseLength int

	fmt.Println("Enter the length of the base of the triangle (between 1 and 10):")
	fmt.Scanln(&baseLength)

	if baseLength < 1 || baseLength > 10 {
		fmt.Println("Invalid input. Please enter a base length between 1 and 10.")
		return
	}

	for i := 1; i <= baseLength; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
