package main

import (
	"fmt"
	"go-reloaded/myfunctions"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("expecting 3 arguments")
		return
	}
	args := os.Args
	inputfile := args[1]
	outputfile := args[2]

	text, err := os.ReadFile(inputfile)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	splitnew := strings.Split(string(text), "\n")
	empty := ""
	for _, word := range splitnew {

		words := strings.Fields(string(word))
		words = myfunctions.Capitalize(words)
		words = myfunctions.Upp(words)
		words = myfunctions.Low(words)
		words = myfunctions.Hex(words)
		words = myfunctions.Bin(words)
		words = myfunctions.Punctuations(words)
		words = myfunctions.Vowel(words)
		words = myfunctions.Apostrophe(words)
		result := strings.Join(words, " ") + "\n"
		empty += result
	}
	results := []byte(empty)
	os.WriteFile(outputfile, results, 0644)

}
