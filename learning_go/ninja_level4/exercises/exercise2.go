package exercises

import (
	"fmt"
)

// Exercise2 ... exercise 2 solution
func Exercise2(a []int) {
	fmt.Println("\nExercise 2:")

	fmt.Println("length of slice is:", len(a))
	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Printf("%v %T\n", "a is of type:", a)
}
