// package main

// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )

func main() {
	if len(os.Args) != 2{
		return
	}
	args := os.Args[1]
	result := ""
	for _, ch := range args{
		if ch >= 'a' && ch <= 'z'{
			result += string((ch - 'a' + 13)%26 + 'a')
		} else if ch >= 'A' && ch <= 'Z'{
			result += string((ch - 'A' + 13)%26 + 'A')
		} else {
			result += string(ch)
		}
	}
	for _, ch := range result{
		z01.PrintRune(ch)
	}
}

// package main

// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )

func main(){
	if len(os.Args) != 2 {
		return
	}
	result := ""
	args := os.Args[1]
	for _, ch := range args{
		if ch >= 'a' && ch <= 'z' {
			result += string('z'- (ch - 'a'))
		}else if ch >= 'A' && ch <= 'Z' {
			result += string('Z'- (ch - 'A'))
		} else {
			result += string(ch)
		}
	}
	for _, ch := range result{
		z01.PrintRune(ch)
	}
}


package main

import (
	"fmt"
)

func main() {
	fmt.Println(Compare("Hello!", "Hello!"))
	fmt.Println(Compare("Salut!", "lut!"))
	fmt.Println(Compare("Ola!", "Ol"))
}

func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}
