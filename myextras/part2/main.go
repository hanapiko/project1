package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	data, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	words := strings.Split(string(data), " ")
	words = Capital(words)
	words = upp(words)
	words = low(words)
	words = vowel(words)
	words = Pncs(words)
	words = hex(words)
	words = bin(words)

	result := strings.Join(words, " ")
	results := []byte(result)
	os.WriteFile(args[1], results, 0644)
}
func Capital(a []string) []string {
	for i := 0; i < len(a); i++ {
		if strings.Contains(a[i], "cap") {
			if strings.Contains(a[i], "cap,") {
				number, _ := strconv.Atoi(a[i+1][:len(a[i+1])-1])
				for b := i - number; b < i; b++ {
					a[b] = strings.Title(a[b])
				}
				a = append(a[:i], a[i+2:]...)
			} else {
				a[i-1] = strings.Title(a[i-1])
				a = append(a[:i], a[i+1:]...)
			}
		}
	}
	return a
}

func upp(a []string) []string {
	for i := 0; i < len(a); i++ {
		if strings.Contains(a[i], "up") {
			if strings.Contains(a[i], "up,") {
				number, _ := strconv.Atoi(a[i+1][:len(a[i+1])-1])
				for b := i - number; b < i; b++ {
					a[b] = strings.ToUpper(a[b])
				}
			} else {
				a[i-1] = strings.ToUpper(a[i-1])
				a = append(a[:i], a[i+1:]...)
			}
		}
	}
	return a
}

func low(a []string) []string {
	for i := 0; i < len(a); i++ {
		if strings.Contains(a[i], "low") {
			if strings.Contains(a[i], "low,") {
				number, _ := strconv.Atoi(a[i+1][:len(a[i+1])-1])
				for b := i - number; b < i; b++ {
					a[b] = strings.ToLower(a[b])
				}
			} else {
				a[i-1] = strings.ToLower(a[i-1])
				a = append(a[:i], a[i+1:]...)
			}
		}
	}
	return a
}

func hex(s []string) []string {
	for i := 0; i < len(s); i++ {
		if s[i] == "(hex)" {
			result, _ := strconv.ParseInt(s[i-1], 16, 32)
			s[i-1] = fmt.Sprint(result)
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func bin(s []string) []string {
	for i := 0; i < len(s); i++ {
		if s[i] == "(bin)" {
			result, _ := strconv.ParseInt(s[i-1], 2, 32)
			s[i-1] = fmt.Sprint(result)
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func vowel(s []string) []string {
	vowels := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}

	for i, word := range s {
		for _, letter := range vowels {
			if word == "a" && string(s[i+1][0]) == letter {
				s[i] = "an"
			} else if word == "A" && string(s[i+1][0]) == letter {
				s[i] = "An"
			}
		}
	}
	return s
}

func Pncs(s []string) []string {
	puncs := []string{",", ".", "!", "?", ":", ";"}

	// punc in the middle of a string connecting to word after
	for i, word := range s {
		for _, punc := range puncs {
			if len(word) > 0 && string(word[0]) == punc && len(word) > 1 && string(word[len(word)-1]) != punc {
				s[i-1] += punc
				s[i] = word[1:]
			} else if (string(word[0]) == punc) && (s[len(s)-1] == s[i]) {
				s[i-1] += word
				s = s[:len(s)-1]
			} else if string(word[0]) == punc && string(word[len(word)-1]) == punc && s[i] != s[len(s)-1] {
				s[i-1] += word
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
	count := 0
	for i, word := range s {
		if word == "'" && count == 0 {
			count += 1
			s[i+1] = word + s[i+1]
			s = append(s[:i], s[i+1:]...)
		}

	}
	//  for second apostrophe
	for i, word := range s {
		if word == "'" {
			s[i-1] = s[i-1] + word
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}
