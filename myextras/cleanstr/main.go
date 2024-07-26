package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
	//"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]

	slc := split(args)

	for i := 0; i < len(slc); i++{
		if i != len(slc) - 1 {
			fmt.Print(slc[i] + " ")
		}else{
			fmt.Print(slc[i])
		}
		
	}
	z01.PrintRune('\n')
}

func split(s string) []string {
	strslice := []string{}
	result := ""

	for i := 0; i < len(s); i++{
		if s[i] != ' ' {
			result += string(s[i])
		} else if s[i] == ' ' && result != ""{
			strslice = append(strslice, result)
			result = ""
		}
	}
	if result != ""{
		strslice = append(strslice, result)
	}
	return strslice
}
