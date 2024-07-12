package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("no arguments found")
		return
	}
	args := os.Args[1]
	g := strings.Split(args, " ")
	for i := len(g)-1; i >= 0; i-- {
		b := g[i]
		if b == " " {
			continue
		} else {
			fmt.Println(b)
			break
		}
	}
}