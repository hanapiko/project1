package check

import (
	"fmt"
	"os"
	"strings"
)

// Usage checks if the program was supplied the expected command-line arguments, exits on usage error
func Usage(args []string) {
	// args =  os.Args[1:]
	if len(args) == 0 {
		// case: go run .
		os.Exit(0)
	} else if len(args) > 3 {
		// checks if the there are more than three arguements
		PrintUsage()
	}
}

// PrintUsage prints the program usage information, then exits the program with the error code 1.
func PrintUsage() {
	usage := "Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard"
	fmt.Println(usage)
	os.Exit(1)
}

// Text
func Text(text string) string {
	// Replace all the newline characters `\n` in the arguement with `\\n`
	text = strings.ReplaceAll(text, "\n", "\\n")

	if text == "" {
		// Nothing to draw
		os.Exit(0)
	} else if text == "\\n" {
		// if the arguement is "\n" the program prints a new line
		fmt.Println()
		os.Exit(0)
	}

	out := ""
	// checking if the runes of the string in the arguement is not of an ascii decimal value of more than 126
	for _, char := range text {
		if char > '~' {
			fmt.Println("Error : Non Ascii character found!! can not display the graphic representation")
			os.Exit(1)
		} else if char >= ' ' {
			// Ignore ASCII characters before the ' ' (space) character
			out += string(char)
		}
	}

	return out
}

// ArtFile given a banner, returns the name of the file with the graphics for the given banner 
func ArtFile(banner string) string {
	switch banner {
	case "standard":
		return "standard.txt"
	case "thinkertoy":
		return "thinkertoy.txt"
	case "shadow":
		return "shadow.txt"
	case "lean":
		return "lean.txt"
	default:
		fmt.Printf("Invalid banner: %q\n", banner)
		PrintUsage()
	}

	return ""
}
