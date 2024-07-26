package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {
	result := []string{}
	for _, ch := range arr {
		if ch == 0 {
			result = append(result, "00")
		} else {
			result = append(result, Hex(int(ch)))
		}
	}

	// res := ""
	for i := 0; i < len(result); i++ {
		Pp(result[i])
		z01.PrintRune(' ')
		if i == 3 || i == 7 || i == 9 {
			z01.PrintRune('\n')
		}
	}
	for _, ch := range arr {
		if ch >= 32 && ch <= 128 {
			z01.PrintRune(rune(ch))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

func Pp(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

func Hex(num int) string {
	res := ""

	hexv := "0123456789abcdef"

	for num > 0 {
		dig := num % 16
		res = string(hexv[dig]) + res
		num /= 16
	}
	return res
}
