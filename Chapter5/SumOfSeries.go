package main

import "fmt"

func main() {
	fmt.Println("n\tSum")

	var sum int64

	for n := 1; n <= 100; n++ {
		sum += int64(n)
		fmt.Printf("%d\t%d\n", n, sum)
	}
}
