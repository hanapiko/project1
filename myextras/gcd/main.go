package main

import (
	"fmt"
)

func main() {
	fmt.Println(Gcd(14, 77))
	fmt.Println(Gcd(17, 3))
}

func Gcd(a, b int) int {
	for b != 0 {
		a,b = b, a%b
	}
	return a
}