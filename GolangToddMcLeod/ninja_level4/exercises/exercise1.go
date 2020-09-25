package exercises

import (
	"fmt"
)

// Exercise1 ... exercise 1 solution
func Exercise1() {
	fmt.Println("Exercise 1:")
	// declare array
	var x [5]int
	x[1] = 2
	// iterate over array and print values
	for _, v := range x {
		fmt.Println(v)
	}
	// print string value and array type using fmt.Printf
	fmt.Printf("%v %T\n", "x is of type:", x)

}
