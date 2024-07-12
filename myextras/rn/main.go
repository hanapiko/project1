package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <number>")
		return
	}

	input := os.Args[1]
	number, err := strconv.Atoi(input)
	if err != nil || number < 1 || number > 4000 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	// Define mappings for Roman numerals and their values
	numeralValues := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	numeralSymbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	// Initialize variables for constructing the Roman numeral representation
	romanNumeral := ""
	calculation := ""

	// Iterate through the numeralValues
	for i := 0; i < len(numeralValues); i++ {
		value := numeralValues[i]
		symbol := numeralSymbols[i]

		// Repeat subtracting the current value while it fits
		for number >= value {
			romanNumeral += symbol
			calculation += symbol
			number -= value
			if i < len(numeralValues)-1 {
				calculation += "-"
			}
		}
	}

	// Print the Roman numeral calculation and the final Roman numeral
	fmt.Println(calculation)
	fmt.Println(romanNumeral)
}
