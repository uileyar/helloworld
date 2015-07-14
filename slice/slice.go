// slice.go
package main

import (
	"fmt"
)

func sum1(a *[]int) int {
	var sum int

	for i := 0; i < len(*a); i++ {
		sum += (*a)[i]
	}

	for i, v := range *a {
		(*a)[i] += v
	}

	return sum
}

func sum2(a []int) int {
	var sum int

	for i := 0; i < len(a); i++ {
		sum += a[i]
	}

	for i, v := range a {
		a[i] += v
	}

	return sum
}

func main() {
	a := []int{1, 2, 3, 4}
	s := sum1(&a)
	fmt.Println(a, s)

	s = sum2(a[0:])
	s = sum2([]int{1, 2, 3, 4})
	fmt.Println(a, s)
}
