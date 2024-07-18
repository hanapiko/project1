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
	//"fmt"

	"fmt"

	//"github.com/01-edu/z01"


)


func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)
	fmt.Println(result)
	
}
func SortWordArr(a []string) {
	for i:=0; i < len(a); i++ {
		for j:=i+1; j < len(a); j++ {
			if a[i] > a[j]{
				a[i],a[j] = a[j],a[i]
			}
		}
	}
	// fmt.Println(a)
}



