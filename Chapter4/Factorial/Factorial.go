package main

import "fmt"

func main() {
	var number, factorial int
	fmt.Println("Enter a number:")
	fmt.Scanln(&number)

	factorial = 1
	for count := number; count > 0; count-- {
		factorial *= count
	}

	fmt.Printf("%d factorial: %d\n", number, factorial)
}
