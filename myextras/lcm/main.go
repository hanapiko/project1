
package main

import (
	"fmt"
)

func main() {
	fmt.Println(Lcm(5, 10))
	fmt.Println(Lcm(42, 12))
}
func Gcd(a, b int) int {
	for b != 0 {
		a,b = b, a%b
	}
	return a
}
func Lcm(a, b int) int {
	// if a == 0 || b == 0 {
	// 	return 0
	// }
	return (a / Gcd(a, b)) * b
}
