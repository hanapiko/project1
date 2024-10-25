// Write a program that finds all pairs of elements in an integer array that sum up to a given target value. The program should output a list of pairs, each representing the indices of the elements that form the pair.

// In this exercise you must take in consideration the following:

//     Ensure it's possible to have positive or negative integers in the array.
//     Ensure each element is used only once in a pair, although the element can be repeated in different pairs.
//     Allow for multiple pairs to sum up to the target value.
//     The output messages should follow the one given in the examples bellow.
//     Return the message "No pairs found." when no pair is present.
//     Return the message "Invalid target sum." if the target is invalid.
//     Return the message "Invalid number: " if the number in the array is invalid.
//     For any input format that deviates from the specified format "[1, 2, 3, 4, 5]" "6", the program will return an "Invalid input." error message.

// Let's consider the input arr = [1, 2, 3, 4, 5] and the target sum targetSum = 6. When we run the program, the findPairs() function will search for pairs in the array that sum up to targetSum.

// Here is some example of outputs:

// $ go run . "[1, 2, 3, 4, 5]" "6"
// Pairs with sum 6: [[0 4] [1 3]]
// $ go run . "[-1, 2, -3, 4, -5]" "1"
// Pairs with sum 1: [[0 1] [2 3]]
// $ go run . "[1, 2, 3, 4, 5]" "10"
// No pairs found.
// $ go run . "[-1, -2, -3, -4, -5]" "-5"
// Pairs with sum -5: [[0 3] [1 2]]
// $ go run . "[1, 2, 3, 4, 20, -4, 5]" "2 5"
// Invalid target sum.
// $ go run . "[1, 2, 3, 4, 20, p, 5]" "5"
// Invalid number: p
// $ go run . "[1, 2, 3, 4" "5"
// Invalid input.
// $ go run . "1, 2, 3, 4" "5"
// Invalid input.
// $

package main

import (
	"fmt"
	"os"
	"strconv"
)

func Pairs(n []int, target int) [][]int {
	res := [][]int{}

	for i := 0; i < len(n); i++ {
		for j := i + 1; j < len(n); j++ {
			if n[i]+n[j] == target {
				res = append(res, []int{i, j})
			}
		}
	}
	return res

}
func parseArray(input string) []int {
	if len(input) < 3 || input[0] != '[' || input[len(input)-1] != ']' {
		fmt.Println("invalid input")
		os.Exit(0)
	} else {
		input = input[1 : len(input)-1]
	}
	arr := Split(input, ", ")
	res := []int{}
	for _, v := range arr {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("Invalid number %v\n", v)
			os.Exit(0)
		}
		res = append(res, num)
	}
	// fmt.Println(res)
	return res
}

func Split(s string, sep string) []string {
	res := []string{}
	start := 0
	for i := 0; i < len(s); i++{
		if i + len(sep) < len(s){
			if s[i:i+len(sep)] == sep{
				res = append(res, s[start:i])
				start = i + len(sep)
			}
		}
		if i == len(s) -1{
			res = append(res, s[start:])
		}
	}
	return res
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("invalid input")
		return
	}
	args := parseArray(os.Args[1])
	num, _ := strconv.Atoi(os.Args[2])

	fmt.Println(Pairs(args, num))
}