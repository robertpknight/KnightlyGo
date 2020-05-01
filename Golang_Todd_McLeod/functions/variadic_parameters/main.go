package main

import (
	"fmt"
)

func main() {
	sum := sum(10, 2, 3)
	fmt.Println("double the sum is:", 2*sum)
}

// variadic parameter - takes any number values of type int and stores them in a slice
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
