package main

import "fmt"

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

func itoa(n int) string {
	result := ""

	for n != 0 {
		mod := n% 10
		startrune := '0'
		for i := 0; i < mod; i++ {
			startrune++
		}
		result = string(startrune) + result
		n = n / 10
	}
	return result
}

func ZipString(s string) string {
	seen := make(map[rune]int)
	result := []rune{}

	for _, char := range s {
		if seen[char] == 0 {
			result = append(result, char)
		}
		seen[char]++
	}
	res := ""
	for _, ch := range result {
		count := seen[ch]
		res += itoa(count) + string(ch)
	}
	return res
}
