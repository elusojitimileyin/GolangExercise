package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	phoneNum     int64
	fullName     string
	cashierName  string
	discount     float64
	amountPaid   float64
	itemList     []string
	quantityList []int
	priceList    []float64
	totalList    []float64
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	fmt.Print("What Is The Customer's Name: ")
	input.Scan()
	fullName = input.Text()

	fmt.Printf("Welcome %s. Kindly Enter Your Phone Number: ", fullName)
	input.Scan()
	phoneStr := input.Text()
	phoneNum, _ = strconv.ParseInt(phoneStr, 10, 64)

	getItemList(input)

	fmt.Print("Enter Cashier's Name: ")
	input.Scan()
	cashierName = input.Text()

	fmt.Print("How Much Discount Will He Get: ")
	input.Scan()
	discountStr := input.Text()
	discount, _ = strconv.ParseFloat(discountStr, 64)

	calculateTotal()

	getDisplayAll()
	getDisplayAllWithPayment()

	fmt.Print("How Much Did The Customer Give To You: ")
	input.Scan()
	amountPaidStr := input.Text()
	amountPaid, _ = strconv.ParseFloat(amountPaidStr, 64)

	for amountPaid < getBillTotal() {
		fmt.Println("Amount paid is less than the bill total. Please enter a valid amount.")
		fmt.Print("How Much Did The Customer Give To You: ")
		input.Scan()
		amountPaidStr = input.Text()
		amountPaid, _ = strconv.ParseFloat(amountPaidStr, 64)
	}

	getDisplayAll()
	getDisplayForAfterPayment()
}

func getItemList(input *bufio.Scanner) {
	for {
		fmt.Print("What Did The User Buy? (yes/no): ")
		input.Scan()
		item := input.Text()
		if strings.ToLower(item) == "no" {
			break
		}
		itemList = append(itemList, item)

		fmt.Print("How Many Pieces? ")
		input.Scan()
		quantityStr := input.Text()
		quantity, _ := strconv.Atoi(quantityStr)
		quantityList = append(quantityList, quantity)

		fmt.Print("How Much Per Unit? ")
		input.Scan()
		priceStr := input.Text()
		price, _ := strconv.ParseFloat(priceStr, 64)
		priceList = append(priceList, price)
	}
}

func calculateTotal() {
	for i := 0; i < len(itemList); i++ {
		totalList = append(totalList, float64(quantityList[i])*priceList[i])
	}
}

func getDisplayAll() {
	fmt.Println()
	fmt.Println("SEMI-COLON STORES")
	fmt.Println()
	fmt.Println("MAIN BRANCH")
	fmt.Println()
	fmt.Println("LOCATION: 312, HERBERT MACULAY WAY, SABO YABA, LAGOS.")
	fmt.Println()
	fmt.Println("TEL:", phoneNum)
	fmt.Println()
	fmt.Println("DATE:", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println()
	fmt.Println("Customer Name:", fullName)
	fmt.Println()
	fmt.Println("Cashier's Name:", cashierName)
	fmt.Println()
	fmt.Println("===============================================================")
	fmt.Println()
	fmt.Printf("%-10s%-10s%-15s%-15s\n", "ITEM", "QTY", "PRICE", "TOTAL(NGN)")
	fmt.Println("---------------------------------------------------------------")
	for i := 0; i < len(itemList); i++ {
		fmt.Printf("%-10s%-10d%-15.2f%-15.2f\n", itemList[i], quantityList[i], priceList[i], totalList[i])
	}
	fmt.Println("---------------------------------------------------------------")
	fmt.Printf("%45s%.2f\n", "SubTotal:", getSubTotal())
	fmt.Printf("%45s%.2f\n", "Discount:", getDiscount())
	fmt.Printf("%45s%.2f\n", "VAT @ 17.50%:", getVAT())
	fmt.Println("===============================================================")
	fmt.Printf("%45s%.2f\n", "Bill Total:", getBillTotal())
	fmt.Println()
}

func getDisplayAllWithPayment() {
	fmt.Println("===============================================================")
	fmt.Printf("%-20s%.2f\n", "THIS IS NOT A RECEIPT KINDLY PAY", getBillTotal())
	fmt.Println("===============================================================")
}

func getSubTotal() float64 {
	var subTotal float64
	for _, total := range totalList {
		subTotal += total
	}
	return subTotal
}

func getDiscount() float64 {
	return getSubTotal() * (discount / 100)
}

func getVAT() float64 {
	return getSubTotal() * (17.50 / 100)
}

func getBillTotal() float64 {
	return getSubTotal() - getDiscount() + getVAT()
}

func getDisplayForAfterPayment() {
	fmt.Println("===============================================================")
	fmt.Printf("%-45s%.2f\n", "Bill Total:", getBillTotal())
	fmt.Printf("%-45s%.2f\n", "Amount Paid:", amountPaid)
	fmt.Printf("%-45s%.2f\n", "Balance:", getBalance())
	fmt.Println("===============================================================")
	fmt.Println("Thank You For Your Patronage")
	fmt.Println("===============================================================")
}

func getBalance() float64 {
	return amountPaid - getBillTotal()
}
