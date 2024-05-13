package main

import (
	"fmt"
)

func main() {

	fmt.Println("add two number without the plus sign")
	var num1, num2, sum int
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	sum = num1 - (-num2)
	fmt.Printf("%d -(- %d) = %d\n", num1, num2, sum)
}
