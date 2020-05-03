package main

import (
	"fmt"
)

// x is in the package scope
var x int

func main() {
	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()
	fmt.Println(x)

	// y is in the main function scope
	var y int
	fmt.Println(y)

	{
		// z is only in this code block scope
		var z int = 42
		fmt.Println(z)
	}

}

func foo() {
	x++
}
