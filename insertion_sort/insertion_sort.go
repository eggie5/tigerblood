// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	a := []int{7, 4, 6, 4, 9, 4, 2, 1}
	// a := []int{1,2,3,4,5,6,7}
	fmt.Println(a)
	insertion_sort(a)
}

func insertion_sort(a []int) {
	for i := 1; i < len(a); i++ {
		j := i
		for j > 0  && (a[j-1] > a[j]){ //while in go
			fmt.Println("inner")
			a[j], a[j-1] = a[j-1], a[j] //inplace swap FTW
			j--
		}

		fmt.Println(a)
	}
}
