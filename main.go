// Author: Natnael Debesay
// Version: 1.0.0
// Date 2025-11-26
// Fileoverview: This program decides whether you should keep going on a car ride or not depending on your gas tank.

package main

import (
	"fmt"
	"os"
)

func main() {
	// Declare variables for input (using float64 for precise calculation)
	var tankCapacity float64    
	var gasPercentage float64   
	var averageMileage float64

	// input and conversion
	fmt.Print("Car Gas Capacity (L): ")
	if _, err := fmt.Scanf("%f", &tankCapacity); err != nil {
		fmt.Println("Error: Invalid capacity input. Please enter a number.")
		os.Exit(1)
	}

	fmt.Print("Gas Gauge (in percentage): ")
	if _, err := fmt.Scanf("%f", &gasPercentage); err != nil {
		fmt.Println("Error: Invalid percentage input. Please enter a number.")
		os.Exit(1)
	}

	fmt.Print("Average Km/L mileage value for the car: ")
	if _, err := fmt.Scanf("%f", &averageMileage); err != nil {
		fmt.Println("Error: Invalid mileage input. Please enter a number.")
		os.Exit(1)
	}

	// process
	fuelLeft := (gasPercentage / 100.0) * tankCapacity 

	distancePossible := fuelLeft * averageMileage

	// output
	
	fmt.Printf("\nCalculated Fuel Left: %.2f L\n", fuelLeft)
	fmt.Printf("Calculated Distance Possible: %.2f Km\n", distancePossible)
	fmt.Println("------------------------------------")

	if distancePossible >= 50 {
		fmt.Println("Safe to proceed")
	} else {
		fmt.Println("Get Gas")
	}

	fmt.Println("\nDone.")
}