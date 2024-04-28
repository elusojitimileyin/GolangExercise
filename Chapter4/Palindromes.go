package main

import (
	"fmt"
	"strconv"
)

func main() {

	word := 22232
	wordNum := strconv.Itoa(word)

	left := 0
	right := len(wordNum) - 1
	palindrome := true

	for left < right {
		if wordNum[left] != wordNum[right] {
			palindrome = false
			break
		}
		left++
		right--
	}

	if palindrome {
		fmt.Println(word, "is a palindrome.")
	} else {
		fmt.Println(word, "is not a palindrome.")
	}
}
