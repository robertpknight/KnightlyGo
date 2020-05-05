package exercise9

import (
	"fmt"

	"strings"
)

// Exercise9 ... code for ninja level 6 exercise 9
func Exercise9() {
	fmt.Println("\nExercise 9:")
	g := greeting("Rob")
	fmt.Println(g)

	fmt.Println("Using shout function to shout a word returning func:")
	fmt.Println(shout(greeting, "Rob"))
	fmt.Println(shout(swear, "Rob"))
}

func greeting(person string) string {
	return "Hello, " + person
}

func swear(person string) string {
	return "Screw you, " + person
}

func shout(words func(string) string, person string) string {
	w := words(person)
	return strings.ToUpper(w)
}
