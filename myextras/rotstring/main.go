package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2{
		fmt.Println()
	}
	word := os.Args[1]
	splitted := split(word)
	words := ""
	for i := 1; i < len(splitted); i++ {
		if i != len(splitted) - 1{
			words += splitted[i] + " " 
		} else {
			words += splitted[i] +  " " + splitted[0] 
		}
	}  
	for _, ch := range words {
		z01.PrintRune(ch)
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