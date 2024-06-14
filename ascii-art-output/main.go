package main

import (
	"fmt"
	"os"
	"strings"

	"ascii/check"
	"ascii/file"
	"ascii/graphics"
)

func main() {
	args := os.Args[1:]
	check.Usage(args)

	text := args[0]
	text = check.Text(text)

	banner := "standard"
	if len(os.Args) > 2 {
		banner = os.Args[2]
	}
	asciiArtFile := check.ArtFile(banner)
	bannerFileContents := file.ReadArtFile(asciiArtFile)

	// the variable inputArgs splits the string(arguement) into substrings, by a separator "\n"
	inputArgs := strings.Split(text, "\\n")

	output := graphics.Graphic(inputArgs, bannerFileContents)
	fmt.Print(output)
}
