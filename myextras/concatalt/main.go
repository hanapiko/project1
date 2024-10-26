package main

import "fmt"

// "github.com/01-edu/go-tests/lib/challenge"
// "github.com/01-edu/go-tests/solutions"

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11, 56, 7}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}

func ConcatAlternate(slice1, slice2 []int) []int {
	slc := []int{}
	if len(slice1) == 0 && len(slice2) == 0 {
		return []int(nil)
	}
	if len(slice1) < len(slice2) {
		slice1, slice2 = slice2, slice1
	}
	for i, val := range slice1 {
		slc = append(slc, val)

		if i < len(slice2) {
			slc = append(slc, slice2[i])
		}
	}
	return slc
}