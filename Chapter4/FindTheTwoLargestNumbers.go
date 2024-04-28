package main

import (
	"fmt"
)

func main() {
	var counter int
	var largest, secondLargest int

	fmt.Println("Enter 10 integers:")

	for counter < 10 {
		var input int
		fmt.Scanln(&input)

		if counter == 0 || input > largest {
			secondLargest = largest
			largest = input
		} else if input > secondLargest {
			secondLargest = input
		}

		counter++
	}

	fmt.Printf("The largest integer is: %d\n", largest)
	fmt.Printf("The second largest integer is: %d\n", secondLargest)
}
