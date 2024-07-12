package main

import (
	"fmt"
)

func Itoa(n int) string {
	// Handle special case for 0
	if n == 0 {
		return "0"
	}

	// Handle negative numbers
	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}

	// Calculate number of digits
	numDigits := 0
	temp := n
	for temp > 0 {
		temp /= 10
		numDigits++
	}

	// Create a byte slice for storing the digits
	result := make([]byte, numDigits)

	// Build the string by extracting digits
	for i := numDigits - 1; i >= 0; i-- {
		result[i] = byte(n % 10 + '0')
		n /= 10
	}

	// Add negative sign if necessary
	if isNegative {
		result = append([]byte{'-'}, result...)
	}

	return string(result)
}

func main() {
	fmt.Println(Itoa(12345))     // Output: "12345"
	fmt.Println(Itoa(0))         // Output: "0"
	fmt.Println(Itoa(-1234))     // Output: "-1234"
	fmt.Println(Itoa(987654321)) // Output: "987654321"
}

// package main

// import (
// 	"fmt"
// )

// func Itoa(n int) string {
// 	// Handle the case when n is 0
// 	if n == 0 {
// 		return "0"
// 	}

// 	// Initialize an empty string to store the result
// 	var result string

// 	// Handle negative numbers
// 	if n < 0 {
// 		result += "-"
// 		n = -n
// 	}

// 	// Extract digits from n and build the string in reverse order
// 	for n > 0 {
// 		digit := n % 10
// 		result = string('0'+digit) + result
// 		n /= 10
// 	}

// 	return result
// }

// func main() {
// 	fmt.Println(Itoa(12345))
// 	fmt.Println(Itoa(0))
// 	fmt.Println(Itoa(-1234))
// 	fmt.Println(Itoa(987654321))
// }
