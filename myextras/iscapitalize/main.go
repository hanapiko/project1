

package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}

func IsCapitalized(s string) bool {

if len(s) == 0{
	return false
}

for i := 0; i < len(s); i++{
	if i == 0 && (s[i] >= 'a' && s[i] <= 'z'){
		return false
	} else {
		if s[i] == ' ' && (s[i + 1] >= 'a' && s[i + 1] <= 'z'){
			return false
		}
	}
}
return true
 }
 
    // checker := false
// B := true
// 	if len(s) == 0 {
// 		return false
// 	}
// 	slc := split(s)

// 	for i:= 0; i < len(slc); i++{
// 		if isAlpha(rune(slc[i][0])){
// B =  true
// 		}else{
// 			B = false
// 		}
// 	}
    

//   return B
// }
// func split(s string) []string {
// 	strslice := []string{}
// 	result := ""

// 	for i := 0; i < len(s); i++{
// 		if s[i] != ' ' {
// 			result += string(s[i])
// 		} else if s[i] == ' ' && result != ""{
// 			strslice = append(strslice, result)
// 			result = ""
// 		}
// 	}
// 	if result != ""{
// 		strslice = append(strslice, result)
// 	}
// 	return strslice
// }

// func isAlpha(char rune) bool {
// 	return char >= 'A' && char <= 'Z' || char >= '0' && char <= '9' || char >= 33 && char <= 64




