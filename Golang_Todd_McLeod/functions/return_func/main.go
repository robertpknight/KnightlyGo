package main

import (
	"fmt"
)

func main() {
	s1 := foo()
	fmt.Println(s1)

	x := bar()
	fmt.Printf("x is of type: %T\n", x)
	fmt.Println("run x():", x())
	fmt.Println("run bar()():", bar()())

}

func foo() string {
	s := "Hello world"
	return s
}

// function bar returns a function that returns an int
// functions are first class citizens
// so can be returned like any other type
func bar() func() int {
	return func() int {
		return 451
	} // this return s the function with no arguments that returns int 451
}
