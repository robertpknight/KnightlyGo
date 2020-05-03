package main

import (
	"fmt"
)

// can do pretty much anything in recursion with loops
// but here is an example anyway

func main() {
	n := factorial(5)
	fmt.Println(n)

	l := loopFactorial(5)
	fmt.Println(l)
}

func factorial(n int) int {
	// must have exit condition to stop recursion running forever
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// better to implement in a loop
// "true sign of a master, is they make the difficult easy"
func loopFactorial(n int) int {
	total := 1
	// you can have no initial condition
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
