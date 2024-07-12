package main

import (
	//"fmt"

	"github.com/01-edu/z01"
)

func main() {
	octet := byte(0b01000001)
	swapped := SwapBits(octet)
	z01.PrintRune(rune(swapped))
	z01.PrintRune('\n')
	//fmt.Println(swapped)
}

func SwapBits(octet byte) byte {
	a := octet >> 4 | octet << 4
	return a
}