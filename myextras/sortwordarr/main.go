// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func SortWordArr(a []string) {
// 	sort.Strings(a)
// }

// func main() {
// 	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
// 	SortWordArr(result)

// 	fmt.Println(result)
// }

package main

import (
	"fmt"
	"strings"
)

func SortWordArr(a []string) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if strings.Compare(a[j], a[j+1]) > 0 {
				// Swap elements if they are in the wrong order
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)

	fmt.Println(result)
}


