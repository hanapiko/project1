package main

import "github.com/01-edu/z01"

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}

func ReduceInt(a []int, f func(int, int) int) {
	first := a[0]
	for i := 1; i < len(a); i++ {
		first = f(first, a[i])
		a := Itoa(first)
		for _, ch := range a {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}

func Itoa(n int) string {
	r := ""
	for n != 0 {
		mod := n % 10
		rn := '0'
		for i := 0; i < mod; i++ {
			rn++
		}
		r = string(rn) + r
		n = n / 10
	}
	return r
}
