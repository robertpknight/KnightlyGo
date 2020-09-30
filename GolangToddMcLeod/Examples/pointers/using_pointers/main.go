package main

import (
	"fmt"
)

func main() {
	x := 1
	fmt.Println("x is at address:", &x)
	fmt.Println("x is value:", x)

	mutate_double_value(&x)
	fmt.Println("x is at address:", &x)
	fmt.Println("x is value:", x)

	mutate_double_value(&x)
	fmt.Println("x is at address:", &x)
	fmt.Println("x is value:", x)

}

// takes type pointer to an int: *int
func mutate_double_value(y *int) {
	fmt.Println("\nmutate by doubling value at address:", y)
	*y = *y * 2
}
