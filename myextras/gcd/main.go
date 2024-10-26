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

// method2
// if a == 0 || b == 0 {
	// 	return 0
	// }
	// var num uint
	// if a < b {
	// 	num = a
	// } else {
	// 	num = b
	// }

	// for i := num; i > 0; i-- {
	// 	if a % i == 0 && b % i == 0 {
	// 		return i
	// 	}
	// }
	// return 0