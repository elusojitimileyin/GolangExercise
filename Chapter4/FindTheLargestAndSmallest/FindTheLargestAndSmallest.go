package main

import (
	"fmt"
)

func main() {
	var counter int
	var largest, smallest int

	fmt.Println("Enter 10 integers:")

	for counter < 10 {
		var input int
		fmt.Scanln(&input)

		if counter == 0 || input > largest {
			largest = input
		}
		if counter == 0 || input < smallest {
			smallest = input
		}

		counter++
	}

	fmt.Printf("The largest integer is: %d\n", largest)
	fmt.Printf("The smallest integer is: %d\n", smallest)
}
