package exercise6

import (
	"fmt"
)

// Exercise6 ... code for ninja level 6 exercise 6
func Exercise6() {
	fmt.Println("\nExercise 6:")

	func(person string) {
		fmt.Println("Hello", person, "(from anonamous func)")
	}("Rob")
}
