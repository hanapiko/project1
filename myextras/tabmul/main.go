package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	first := args[0]
	num := 0
	for _, r := range first{
		if r >= '0' && r <= '9'{
			num = num * 10 + int(r - '0')
		} else {
			return
		}
	}
	ans := 0
	for i := 1; i <= 9; i++{
		z01.PrintRune(rune(i + '0'))
		z01.PrintRune(' ')
		z01.PrintRune('*')
		z01.PrintRune(' ')
		Itoa(num)
		z01.PrintRune(' ')
		z01.PrintRune('=')
		z01.PrintRune(' ')
		ans = i * num
		Itoa(ans)
		z01.PrintRune('\n')

	}

}
 func Itoa(n int) string {
	result := ""
	for n != 0 {
		mod := n % 10
		startrune := '0'
		for i := 0; i < mod; i++ {
			startrune++
		}
		result = string(startrune) + result
		n = n / 10
	}
	for _, ch := range result{
		z01.PrintRune(ch)
	}
	return result
 }

//  func Atoi(s string) int {
// 	num := 0
// 	for _, ch := range s {
// 		num = num *10 + int(ch - '0')
// 	}
// 	return num
//  }

