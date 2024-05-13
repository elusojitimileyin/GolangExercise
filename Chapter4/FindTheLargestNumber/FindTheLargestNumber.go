package main

import (
	"fmt"
)

func main() {
	var counter, largest int

	fmt.Println("Enter 10 integers:")

	for counter < 10 {
		var input int
		fmt.Scanln(&input)

		if counter == 0 || input > largest {
			largest = input
		}

		counter++
	}

	fmt.Printf("The largest integer is: %d\n", largest)
}
