package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if the number of arguments is exactly 2
	if len(os.Args) != 3 {
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]
	seen := make(map[rune]bool)
	for _, char := range str2 {
		seen[char] = true
	}
	var result string
	for _, char := range str1 {
		if seen[char] {
			result += string(char)
			seen[char] = false 
		}
	}
	fmt.Println(result)
}
