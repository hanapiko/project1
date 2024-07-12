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

	input := os.Args[1]

	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("00000000")
		return
	}

	// Format the number as an 8-bit binary string
	binaryStr := fmt.Sprintf("%08b", num)

	fmt.Printf("%s", binaryStr)
}
