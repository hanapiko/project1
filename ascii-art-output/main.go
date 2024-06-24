package main

import (
    "ascii/functions"
    "fmt"
    "os"
    "strings"
)

func main() {
    args := os.Args[1:] // Exclude the program name itself

    // Check if there are no arguments or less than minimum required
    if len(args) < 2 {
        fmt.Println("Usage: go run . --output=<fileName.txt> [OPTION] [STRING] [BANNER]")
        return
    }

    var outputFile string
    var banner string
    var stringArg string

    // Parse arguments
    for _, arg := range args {
        if strings.HasPrefix(arg, "--output=") {
            outputFile = strings.TrimPrefix(arg, "--output=")
        } else if banner == "" && (arg == "standard" || arg == "thinkertoy" || arg == "shadow" || arg == "lean") {
            banner = arg
        } else if stringArg == "" {
            stringArg = arg
        }
    }

    // Validate and process arguments
    if outputFile == "" {
        fmt.Println("Error: Missing --output=<fileName.txt> flag.")
        fmt.Println("Usage: go run . --output=<fileName.txt> [OPTION] [STRING] [BANNER]")
        return
    }

    if stringArg == "" {
        fmt.Println("Error: Missing STRING argument.")
        fmt.Println("Usage: go run . --output=<fileName.txt> [OPTION] [STRING] [BANNER]")
        return
    }

    if stringArg == "\\n" {
        fmt.Println()
        return
    }

    // Validate ASCII characters
    for _, chr := range stringArg {
        if chr > '~' {
            fmt.Println("Error: Non ASCII character found! Cannot display the graphic representation.")
            return
        }
    }

    // Replace newline characters and split arguments
    stringArg = strings.ReplaceAll(stringArg, "\n", "\\n")
    splitArgs := strings.Split(stringArg, "\\n")

    // Determine ASCII art file based on banner type
    var asciiFile string
    switch banner {
    case "standard":
        asciiFile = "standard.txt"
    case "thinkertoy":
        asciiFile = "thinkertoy.txt"
    case "shadow":
        asciiFile = "shadow.txt"
    case "lean":
        asciiFile = "lean.txt"
    default:
        fmt.Printf("Invalid banner: %q\n", banner)
        fmt.Println("Usage: go run . --output=<fileName.txt> [OPTION] [STRING] [BANNER]")
        return
    }

    // Read ASCII art from file
    ascii, err := os.ReadFile(asciiFile)
    if err != nil {
        fmt.Println(err)
        return
    }
    asciiString := string(ascii)
    var characters []string
    if asciiFile == "thinkertoy.txt" {
        characters = strings.Split(asciiString, "\r\n")
    } else {
        characters = strings.Split(asciiString, "\n")
    }

    // Generate output from custom function
    output := functions.Graphic(splitArgs, characters)

    // Output to file
    file, err := os.Create(outputFile)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    _, err = file.WriteString(output)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("Output written to %s\n", outputFile)
}
