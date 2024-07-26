package main

import (
	//"fmt"
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	word := os.Args[1]
	if word == ""{
		z01.PrintRune('\n')
	}
	slic := split(word)
	

	for i := len(slic)-1; i >= 0; i-- {
		if i != 0 {
			fmt.Print(slic[i] + " ")
		} else {
			fmt.Print(slic[i])
		}
	}
	z01.PrintRune('\n')
	//fmt.Println(slic)
	
}

func split(str string)[]string{
	slc := []string{}
	result := ""
	for _, ch := range str{
		if ch != ' '{
			result += string(ch)
		} else if ch == ' ' && result!= ""{
			slc = append(slc, result)
			result = ""
		}
	}
	if result != ""{
		slc =append(slc, result)
	}
	return slc
}