package exercise10

import (
	"fmt"
)

// Exercise10 ... code for ninja level 6 exercise 10
func Exercise10() {
	fmt.Println("\nExercise 10:")
	fmt.Println(greeting("Rob"))
	fmt.Println(greeting("Todd"))
}

func greeting(person string) string {
	// gList enclosed by function
	gList := []string{"Hello", "Morning", "Evening", "G'day"}
	greet := ""
	if len(person) >= 4 {
		greet = gList[3]
	} else {
		greet = gList[len(person)-1]
	}

	return greet + ", " + person
}
