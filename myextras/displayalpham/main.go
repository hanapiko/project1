// // package main

// // import (
// // 	"os"

// // 	"github.com/01-edu/z01"
// // )

// // func main() {
// // 	if len(os.Args) != 3 {
// // 		return
// // 	}
// // 	str1 := os.Args[1]
// // 	str2 := os.Args[2]
// // 	count := 0

// // 	for _, ch := range str2{
// // 		if ch == rune(str1[count]){
// // 			count++
// // 			if len(str1) == count{
// // 				for _, ch := range str1{
// // 					z01.PrintRune(ch)
// // 				}
// // 				z01.PrintRune('\n')
// // 				count = 0
// // 			}
// // 		}
// // 	}
// // }

// package main

// import (
// 	//"strconv"

// 	"github.com/01-edu/z01"
// )

// //"fmt"
// //"os"

// //"github.com/01-edu/z01"

// // import (
// // 	"github.com/01-edu/z01"
// // )

// // func main() {
// // 	z01.PrintRune(FirstRune("Hello!"))
// // 	z01.PrintRune(FirstRune("Salut!"))
// // 	z01.PrintRune(FirstRune("Ola!"))
// // 	z01.PrintRune('\n')
// // }

// // func FirstRune(s string) rune {
// // 	srune := []rune(s)
// // 	return srune[len(srune)-1]
// // }

// // func main() {
// // 	if len(os.Args) != 2 {
// // 		return
// // 	}
// // 	args := os.Args[1]
// // 	last := ""
// // 	for i := len(args)-1; i >= 0; i-- {
// // 		if string(args[i]) != " "{
// // 			last = string(args[i]) + last
// // 		} else {
// // 			if last != "" {
// // 				break
// // 			}
// // 		}
// // 	}
// // 	for _, ch := range last {
// // 		z01.PrintRune(ch)
// // 	}
// // }

// func main() {
// 	mul := func(acc int, cur int) int {
// 		return acc * cur
// 	}
// 	sum := func(acc int, cur int) int {
// 		return acc + cur
// 	}
// 	div := func(acc int, cur int) int {
// 		return acc / cur
// 	}
// 	as := []int{500, 2}
// 	ReduceInt(as, mul)
// 	ReduceInt(as, sum)
// 	ReduceInt(as, div)
// }

// func ReduceInt(a []int, f func(int, int) int) {
// 	first := a[0]
// 	for i := 1; i < len(a); i++{
// 		first = f(first, a[i])
// 	}
// 	result := Itoa(first)
// 	for _, ch := range result{
// 		z01.PrintRune(ch)
// 	}
// 	z01.PrintRune('\n')
// 	}

// func Itoa(n int) string {
// 	str := ""
// 	for n != 0 {
// 		firstrune := '0'
// 		mod := n % 10
// 		for i := 0; i < mod; i++ {
// 			firstrune++
// 		}
// 		str = string(firstrune) + str
// 		n = n / 10
// 	}
// 	return str
// }

package main

import (
	"os"

	"github.com/01-edu/z01"
)

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

