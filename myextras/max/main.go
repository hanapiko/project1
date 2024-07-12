package main

import (
	"fmt"
)

func main() {
	a := []int{90, 1, 200}
	max := Max(a)
	fmt.Println(max)
}

func Max(a []int) int {
	min := a[0]
	for i := 0; i < len(a); i++ {
		if min < a[i] {
			min = a[i]
		}
	}
	return min
} 