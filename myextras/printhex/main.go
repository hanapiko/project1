package main

import (
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	str := os.Args[1]
	hex := "0123456789abcdef"
	result := ""
	num := 0
	for _, ch := range str {
		if ch >= '0' && ch <= '9' {
			num = num*10+ int(ch -'0')
		}	
	}
	//num, _ := strconv.Atoi(str)

	for num > 0 {
		mod := num % 16
		result = string(hex[mod]) + result
		num /= 16
	}

	fmt.Println(result)


	// args := os.Args[1]
	// number, err := strconv.Atoi(args)
	// if err != nil || number < 0 {
	// 	fmt.Println("ERROR")
	// 	return
	// }
	// base := strconv.FormatInt(int64(number), 16)
	// fmt.Println(base)
}