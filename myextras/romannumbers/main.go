package main

import (
	"fmt"
	"os"
)
func main() {
	args := os.Args[1]
	num := Atoi(args)
	if num <= 0 || num >= 4000 {
		fmt.Println("ERROR: cannot convert to roman digit.")
		os.Exit(0)

	}
	cal, res := Rom(num)
	fmt. Println(cal)
	fmt.Println(res)

}

func Rom(n int) (string, string) {
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX","V", "VI", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	calculation := []string{"M","(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)","V", "(V-I)", "I"}

	cal := ""
	result := ""
	for i, c := range values{
		for n >= c{
			n = n - c
			result += roman[i]
			if len(cal) > 0 {
				cal = cal + "+"
			}
			cal += calculation[i]
		}
	}
	return cal, result
}

func Atoi(s string) int {
	val := 0
	for _, ch := range s{
		if ch < '0' || ch > '9' {
			fmt.Println("ERROR: cannot convert to roman digit.")
			os.Exit(0)
		}
		val = val * 10 + int(ch -'0')
	}
	return val
}