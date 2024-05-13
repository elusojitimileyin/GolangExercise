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

	for index := 1; index <= baseLength; index++ {
		for count := 1; count <= index; count++ {
			fmt.Print("# ")
		}
		fmt.Println()
	}
}
