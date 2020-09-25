package main

import (
	"fmt"
)

func main() {
	// assign increment to different variables
	// this creates two locations in memory
	a := incrementor()
	b := incrementor()

	// increment var x in the scope of a
	fmt.Println(a())
	fmt.Println(a())

	// this will increment x in the scope of b
	// same identifier with different memory location
	fmt.Println(b())

	// this will increment  var x in the scope of a again
	fmt.Println(a())
}

func incrementor() func() int {
	// closure is enclosing a variable in the code block
	// x is in the scope of this function
	var x int
	return func() int {
		x++
		return x
	}
}
