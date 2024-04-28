package main

import "fmt"

func main() {
	var totalMiles, totalGallons int
	trips := 0

	for {
		var miles, gallons int
		fmt.Printf("Enter miles driven for trip %d (or -1 to exit): ", trips+1)
		fmt.Scanln(&miles)

		if miles == -1 {
			break
		}

		fmt.Printf("Enter gallons used for trip %d: ", trips+1)
		fmt.Scanln(&gallons)

		trips++
		totalMiles += miles
		totalGallons += gallons

		mpg := float64(miles) / float64(gallons)
		fmt.Printf("Miles per gallon for trip %d: %.2f\n", trips, mpg)

		combinedMPG := float64(totalMiles) / float64(totalGallons)
		fmt.Printf("Combined miles per gallon up to trip %d: %.2f\n", trips, combinedMPG)
	}
}
