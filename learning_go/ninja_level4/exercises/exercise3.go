package exercises

import (
	"fmt"
)

// Exercise3 ... exercise 3 solution
func Exercise3(a []int) {
	fmt.Println("\nExercise 3:")
	// slices of a slice
	fmt.Println(a[0:5])
	fmt.Println(a[5:])
	fmt.Println(a[2:7])
	fmt.Println(a[1:6])
	fmt.Println(a)
}
