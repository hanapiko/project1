package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	check := false
	args := os.Args[1]
	for _, c := range args {
		if c != ' '{
			check = true
			break
		}
	}
	if check {
	result := ""
	for i := len(args)-1; i >= 0; i-- {
		if string(args[i]) != " "{
			result = string(args[i]) + result
		} else if result != ""{
			break
		}
	}
	for _, ch := range result{
		z01.PrintRune(ch)
	}
}

}