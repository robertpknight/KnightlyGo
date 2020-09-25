package exercise7

import (
	"fmt"
)

// Exercise7 ... code for ninja level 6 exercise 7
func Exercise7() {
	fmt.Println("\nExercise 7:")

	greet := func(person string) {
		fmt.Println("Hello", person)
	}
	fmt.Printf("Type of greet is: %T\n", greet)
	greet("Function Expression Master")
}
