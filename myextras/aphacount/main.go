package main

import (
	"fmt"
)

func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}

func AlphaCount(s string) int {
	count := 0
	for _, ch := range s {
		if ch > 'a' && ch <= 'z' || ch > 'A' && ch <= 'Z' {
			count++
		}
	}
	return count
}