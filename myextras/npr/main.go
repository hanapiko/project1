package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		// fmt.Println("Error")
		return
	}
	args := splitt(os.Args[1])
	// arg := splitt(args)
	num, er := nprc(args)
	if er != "" {
		fmt.Println("Error")
		return
	}
	fmt.Println(num)
}

func nprc(s []string) (int, string) {
	num := []int{}
	for _, str := range s {
		if str == "+" || str == "-" || str == "*" || str == "/" || str == "%" {
			if len(num) < 2 {
				return 0, "error"
			}
			a := num[len(num)-2]
			b := num[len(num)-1]
			num = num[:len(num)-2]
			// ing, err := strconv.Atoi(a)
			// if err != nil {
			// 	return 0, "error"
			// }
			// b, err := strconv.Atoi(b)
			// if err != nil {
			// 	return 0, "error"
			// }

			cal := 0
			if str == "+" {
				cal = a + b
			}
			if str == "-" {
				cal = a - b
			}
			if str == "/" {
				cal = a / b
			}
			if str == "*" {
				cal = a * b
			}
			if str == "%" {
				cal = a % b
			}
			num = append(num, cal)
		} else {
			nums, err := strconv.Atoi(str)
			if err != nil {
				return 0, "error"
			}
			num = append(num, nums)
		}
	}
	if len(num) != 1 {
		return 0, "error"
	}
	return num[0], ""
}

func splitt(s string) []string {
	result := ""
	slc := []string{}
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			result += string(s[i])
		} else if s[i] == ' ' && result != "" {
			slc = append(slc, result)
			result = ""
		}
	}
	if result != "" {
		slc = append(slc, result)
	}
	return slc
}
