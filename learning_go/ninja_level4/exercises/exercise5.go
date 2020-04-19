package exercises

import (
	"fmt"
)

// Exercise5 ... exercise 5 solution
func Exercise5(a []int) {
	fmt.Println("\nExercise 5:")
	a = append(a[0:3], a[6:]...)
	fmt.Println(a)
}
