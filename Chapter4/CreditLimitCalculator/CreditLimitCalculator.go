package main

import "fmt"

func main() {
	var accountNumber, beginningBalance, totalCharges, totalCredits, creditLimit int

	fmt.Println("Enter account number:")
	fmt.Scanln(&accountNumber)

	fmt.Println("Enter beginning balance:")
	fmt.Scanln(&beginningBalance)

	fmt.Println("Enter total charges for this month:")
	fmt.Scanln(&totalCharges)

	fmt.Println("Enter total credits for this month:")
	fmt.Scanln(&totalCredits)

	fmt.Println("Enter credit limit:")
	fmt.Scanln(&creditLimit)

	newBalance := beginningBalance + totalCharges - totalCredits

	fmt.Printf("New balance for account %d: %d\n", accountNumber, newBalance)

	if newBalance > creditLimit {
		fmt.Println("Credit limit exceeded")
	}
}
