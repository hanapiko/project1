package main

import "fmt"

func AtoiBase(s string, base string) int {
	if len(base) < 2 {
		return 0
	}

	// Check uniqueness of characters in base
	seen := make(map[rune]bool)
	for _, char := range base {
		if char == '+' || char == '-' {
			return 0
		}
		if seen[char] {
			return 0
		}
		seen[char] = true
	}

	// Convert string s from base to decimal
	baseMap := make(map[byte]int)
	for i := 0; i < len(base); i++ {
		baseMap[base[i]] = i
	}

	result := 0
	b := len(base)
	for _, char := range s {
		if _, ok := baseMap[byte(char)]; !ok {
			return 0 // character in s is not in the base
		}
		result = result*b + baseMap[byte(char)]
	}

	return result
}

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))               // Output: 125
	fmt.Println(AtoiBase("1111101", "01"))                  // Output: 125
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))         // Output: 125
	fmt.Println(AtoiBase("uoi", "choumi"))                  // Output: 125
	fmt.Println(AtoiBase("bbbbbab", "-ab"))                 // Output: 0
	fmt.Println(AtoiBase("1010", "01"))                     // Output: 10
	fmt.Println(AtoiBase("1A", "0123456789ABCDEF"))         // Output: 26
	fmt.Println(AtoiBase("ZZZ", "ABCDEFGHIJKLMNOPQRSTUVWXYZ")) // Output: 18279
}
