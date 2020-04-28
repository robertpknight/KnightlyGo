package exercises

import (
	"fmt"
)

// Exercise9 ... exercise 9 solution
func Exercise9(agents map[string][]string) map[string][]string {
	fmt.Println("\nExercise 9:")
	// add to map
	agents["knight_rob"] = []string{"data", "house music", "travelling"}

	for k, v := range agents {
		fmt.Println("key:", k, "\tvalues:", v)
	}
	return agents
}
