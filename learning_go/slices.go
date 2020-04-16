package main

import (
	"fmt"
)

func main() {
	// composite literal or a slice
	// a slice allows you to group together VALUES of the same TYPE
	// they create a new value every time they are evaluated
	x := []int{4, 5, 7, 8, 42}
	y := []int{69, 27}

	// length of a slice
	fmt.Println(len(x))

	// accessing indexes
	fmt.Println("The first value is:", x[0])

	// slicing a slice
	fmt.Println(x[1:3])

	// iterate over a slice (shorthand)
	for i, v := range x {
		fmt.Println(i, ":", v)
	}

	// iterate over a slice
	for i := 0; i <= 4; i++ {
		// this is using the iterator index to access the slice values at index i
		fmt.Println(i, "-", x[i])
	}

	// append returns a type of the same time i.e. a slice
	z := append(x, 2, 3)
	fmt.Println(z)

	// append but unfurl y variatic parameter, unlimited number of values in y
	x = append(x, y...)
	fmt.Println(x)

	// delete from a slice - remove index 2 and 3
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	// using make built-in function for increased efficiency if you know the length
	// underlying array for slice has values, length and capacity
	m := make([]int, 10, 11)
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))

	// fill length of slice
	m = append(m, 343)
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))

	// since this breaches capacity the size of the underlying array has doubled
	m = append(m, 999)
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))

	// multidimensional slices - slice of slice
	rk := []string{"rob", "knight"}
	ll := []string{"liz", "lord"}
	rel := [][]string{rk, ll}
	fmt.Println(rel)
}
