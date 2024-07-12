package main

import (
	"fmt"
	"os"
	"strconv"
)

// Function to check if a number is prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Function to calculate the sum of all prime numbers <= n
func sumOfPrimes(n int) int {
	sum := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	return sum
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("0")
		return
	}

	input := os.Args[1]
	number, err := strconv.Atoi(input)
	if err != nil || number <= 0 {
		fmt.Println("0")
		return
	}

	sum := sumOfPrimes(number)
	fmt.Println(sum)
}
