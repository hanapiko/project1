package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	number, err := strconv.Atoi(args)
	if err != nil || number < 0 {
		fmt.Println("ERROR")
		return
	}
	base := strconv.FormatInt(int64(number), 16)
	fmt.Println(base)
}