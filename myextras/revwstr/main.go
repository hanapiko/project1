package main

import (
	"fmt"
	"os"
)

// Function to reverse the order of words in a string
func reverseWords(input string) string {
	if input == "" {
		return ""
	}

	var reversed string
	var word string

	for _, ch := range input {
		if ch == ' ' {
			if word != "" {
				reversed = word + " " + reversed
				word = ""
			}
		} else {
			word += string(ch)
		}
	}

	if word != "" {
		reversed = word + " " + reversed
	}

	return reversed[:len(reversed)-1] // Remove the trailing space
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	fmt.Println(reverseWords(input))
}
