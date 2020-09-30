package exercises

import (
	"fmt"
)

// Exercise10 ... exercise 10 solution
func Exercise10(agents map[string][]string) {
	fmt.Println("\nExercise 10:")
	// delete value from map
	if _, ok := agents["no_dr"]; ok {
		fmt.Println("removing dr no")
		delete(agents, "no_dr")
	}

	for k, v := range agents {
		fmt.Println("key:", k, "values:", v)
	}

	// try removing again
	if _, ok := agents["no_dr"]; ok {
		fmt.Println("removing dr no")
		delete(agents, "no_dr")
	} else {
		fmt.Println("dr no not okay - can't remove")
	}
}
