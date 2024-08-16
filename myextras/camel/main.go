// Write a function that converts a string from camelCase to snake_case.

// If the string is empty, return an empty string.
// If the string is not camelCase, return the string unchanged.
// If the string is camelCase, return the snake_case version of the string.
// For this exercise you need to know that camelCase has two different writing alternatives that will be accepted:

// lowerCamelCase
// UpperCamelCase
// Rules for writing in camelCase:

// The word does not end on a capitalized letter (CamelCasE).
// No two capitalized letters shall follow directly each other (CamelCAse).
// Numbers or punctuation are not allowed in the word anywhere (camelCase1).

package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	if !isvalid(s) {
		return s
	}
	result := ""
	for i := 0; i < len(s); i++ {
		r := s[i]
		if r >= 'A' && r <= 'Z' {
			if i > 0 {
				result += "-"
			}
			result += string(r)
		} else {
			result += string(r)

		}
	}
	return result
}
func isvalid(s string) bool {
	if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {
		return false
	}
	for i := 0; i < len(s); i++ {
		r := s[i]
		if r >= 'a' && r <= 'z' {
			continue
		}
		if i > 0 && r >= 'A' && r <= 'Z' && s[i+1] >= 'A' && s[i+1] <= 'Z' {
			return false

		}
	}
	return true
}
