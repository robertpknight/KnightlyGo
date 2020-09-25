package main

import (
	"fmt"
)

func main() {
	foo()

	// anonamous function
	func(x int) {
		fmt.Println("printing number in anonamous function call:", x)
	}(42) // have to provide arguments straight after function definition

	// func expression
	fexp := func() {
		fmt.Println("this print comes from the function assigned to variable fexp")
	}
	fexp() // calls function defined and assinged to a variable

}

func foo() {
	fmt.Println("foo ran")
}
