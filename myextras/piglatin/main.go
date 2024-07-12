package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	word := os.Args[1]

	pigLatin := convertToPigLatin(word)

	if pigLatin == "" {
		fmt.Println("No vowels")
	} else {
		fmt.Println(pigLatin)
	}
}

func convertToPigLatin(word string) string {
	for i, c := range word {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
			c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
			if i == 0 {
				return word + "ay"
			}
			return word[i:] + word[:i] + "ay"
		}
	}
	return ""
}
