package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	str1 := args[0]
	str2 := []rune(args[1])
	str3 := []rune(args[2])


	for _, ch := range str1{
		if ch == str2[0]{
			ch = str3[0]
		}
		z01.PrintRune(ch)
	}
}