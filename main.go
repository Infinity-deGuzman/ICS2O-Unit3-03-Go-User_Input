// Created by: Infinity de Guzman
// Created on: May 2021
//
// This program calculates the area of a trapezoid

package main

import (
	"fmt"

	"github.com/leekchan/accounting"
)

func main() {
	// This function calculates the sphere's volume
	var radius float64

	// input
	accountingFormater := accounting.Accounting{Precision: 2}

	fmt.Println("This program calculates the volume of a sphere.")
	fmt.Println()
	fmt.Print("Enter the radius: ")
	fmt.Scanln(&radius)

	// process
	var volume = (4 / 3) * (radius * *3) * 3.14159265

	// output
	fmt.Println("The volume is: ", accountingFormater.FormatMoney(volume), "cm³")
	fmt.Println("\nDone.")
}
