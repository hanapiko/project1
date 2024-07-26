package main

import "fmt"

func main() {
		fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
		fmt.Print(FifthAndSkip("This is a short sentence"))
		fmt.Print(FifthAndSkip("1234"))
	

}

func FifthAndSkip(str string) string {
	if len(str) < 5 {
		return "Invalid Input\n"
	}
	word := ""
	for _, ch := range str{
		if ch == ' '{
			continue
		}
		word += string(ch)
	}
	result := ""
	for i := 0; i < len(word); i += 5{
		if i+5 >  len(word){
			result += word[i:]
		} else {
			result += word[i:i+5] + " "
		}
		i++
	}
	return result + "\n"
}
