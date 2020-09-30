package exercises

import (
	"fmt"
)

// Exercise4 ... exercise 4 solution
func Exercise4(a []int) {
	fmt.Println("\nExercise 4:")
	fmt.Println(a)

	a = append(a, 52)
	fmt.Println(a)

	a = append(a, 53, 54, 55)
	fmt.Println(a)

	y := []int{56, 57, 58, 59, 60}
	a = append(a, y...)
	fmt.Println(a)
}
