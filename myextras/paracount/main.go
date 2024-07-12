package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	a := len(args)
	z01.PrintRune(rune(a + '0'))
}