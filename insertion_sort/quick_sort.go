package main

import "fmt"

func main() {

	a := []int{5, 3, 2, 6, 4, 1, 3, 7}
	// a := []int{1,2,3,4,5,6,7}
	fmt.Println(a)
	quick_sort(a, 0, len(a))
}

func quick_sort(a []int, p int, r int) {
	if p < r {
		q := partition(a, p, r)
		quick_sort(a, p, q)
		quick_sort(a, q+1, r)

	}

}

func partition(a []int, p int, r int) (pivot int) {
	x := a[p] //guesing this is median of a
	i := p
	j := r-1
	fmt.Println("pivot=", x)
	fmt.Println(j)
	fmt.Println(i)

	for i < j {

		for (a[j] >= x) { //iterate until we found elemetn in wrong index
			j--
			fmt.Println(fmt.Sprintf("loop j=%d", j))
		}
		fmt.Println("need to swap right index ", j)

		for a[i] < x {
			i++
			fmt.Println(fmt.Sprintf("i=%d", i))
		}
		fmt.Println("need to swap left index ", i)

		fmt.Println("broke")
		if i < j {
			fmt.Println("swapping ", j, i)
			a[i], a[j] = a[j], a[i] //swap
			fmt.Println(a)
			//i++
			//j-- //for next iteration
		}
	}

	return i

}
