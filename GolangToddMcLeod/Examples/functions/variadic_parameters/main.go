package main

import (
	"fmt"
)

func main() {
	s := sum(10, 2, 3)
	fmt.Println("double the sum is:", 2*s)

	// unfurl a slice
	nums := []int{1, 2, 3, 4, 5, 7, 8}
	numsSum := sum(nums...)
	fmt.Println("double the sum is:", 2*numsSum)

	// variadic means 0 or more - so can pass nothing to function
	fmt.Println("can give 0 values to sum to get:", sum())
}

// variadic parameter - takes any number values of type int and stores them in a slice
// variadic parameter has to be the last paramater, cannot put normal parameters after
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	s := 0
	for i, v := range x {
		s += v
		fmt.Println("for item number", i, "current sum:", s)
	}
	return s
}
