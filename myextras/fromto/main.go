package main

import (
	"fmt"
	//"github.com/01-edu/z01"
)

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 100))
}

func FromTo(from int, to int) string {
	result := ""
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}
	if from < to {
		for i := from; i <= to; i++ {
			if i < 10 {
				result += "0" + itoa(i)
			} else {
				result += itoa(i)
			}
			if i != to {
				result += ", "
			}
		}
	} else {
		for i := from; i >= to; i-- {
			if i < 10 {
				result += "0" + itoa(i)
			} else {
				result += itoa(i)
			}
			if i != to {
				result += ", "
			}
		}
	}
	return result + "\n"
}

func itoa(n int) string {
	result := ""

	for n != 0 {
		mod := n% 10
		startrune := '0'
		for i := 0; i < mod; i++ {
			startrune++
		}
		result = string(startrune) + result
		n = n / 10
	}
	return result
}
