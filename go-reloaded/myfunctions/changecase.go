package myfunctions

import (
	"fmt"
	"strconv"
	"strings"
)

func Capitalize(s []string) []string {
	for i := 0; i < len(s); i++ {
		if strings.Contains(s[i], "(cap") || strings.Contains(s[i], "(CAP") || strings.Contains(s[i], "(Cap") || strings.Contains(s[i], "(CAp") || strings.Contains(s[i], "(cAp") {
			if strings.Contains(s[i], "(cap,") || strings.Contains(s[i], "(CAP,") {
				number, err := strconv.Atoi(s[i+1][:len(s[i+1])-1])
				if err != nil { //|| number > len(s) || number < 0 || number <= len(s[:i]) || number> len(s[:i]){
					fmt.Println("Error at conversion  or cap is out of range.", err)
					continue
				}

				if number <= len(s[:i]) {
					for j := i - number; j < i; j++ {
						s[j] = strings.ToUpper(string(s[j][0])) + strings.ToLower(string(s[j][1:]))
					}
					s = append(s[:i], s[i+2:]...)
				} else if number > len(s[:i]) {
					for r := 0; r < len(s[:i]); r++ {
						s[r] = strings.ToUpper(string(s[r][0])) + strings.ToLower(string(s[r][1:]))
					}
					s = append(s[:i], s[i+2:]...)
				}
			} else {
				s[i-1] = strings.ToUpper(string(s[i-1][0])) + strings.ToLower(string(s[i-1][1:]))
				s = append(s[:i], s[i+1:]...)

			}
		}
	}
	return s
}

func Upp(s []string) []string {
	for i := 0; i < len(s); i++ {
		if strings.Contains(s[i], "(up") || strings.Contains(s[i], "(UP") {
			if strings.Contains(s[i], "(up,") || strings.Contains(s[i], "(UP,") {
				number, err := strconv.Atoi(s[i+1][:len(s[i+1])-1]) // converting the string to integer
				if err != nil {
					fmt.Println(err)
				}
				if number <= len(s[:i]) {
					for j := i - number; j < i; j++ {
						s[j] = strings.ToUpper(s[j])
					}
					s = append(s[:i], s[i+2:]...)
				} else if number > len(s[:i]) {
					for k := 0; k < len(s[:i]); k++ {
						s[k] = strings.ToUpper(s[k])
					}
					s = append(s[:i], s[i+2:]...)
				}
			} else {
				s[i-1] = strings.ToUpper(s[i-1])
				s = append(s[:i], s[i+1:]...)

			}
		}
	}
	return s
}

func Low(s []string) []string {
	for i := 0; i < len(s); i++ {
		if strings.Contains(s[i], "(low") || strings.Contains(s[i], "(LOW") {
			if strings.Contains(s[i], "(low,") {
				number, err := strconv.Atoi(s[i+1][:len(s[i+1])-1])
				if err != nil {
					fmt.Println(err)
				}
				if number <= len(s[:i]) {
					for j := i - number; j < i; j++ {
						s[j] = strings.ToLower(s[j])
					}
					s = append(s[:i], s[i+2:]...)
				} else if number > len(s[:i]) {
					for k := 0; k < len(s[:i]); k++ {
						s[k] = strings.ToLower(s[k])
					}
					s = append(s[:i], s[i+2:]...)
				}
			} else {
				s[i-1] = strings.ToLower(s[i-1])
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
	return s
}
