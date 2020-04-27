package exercises

import (
	"fmt"
)

// Exercise7 ... exercise 7 solution
func Exercise7() {
	fmt.Println("\nExercise 7:")
	// create slice of slice
	james := []string{"James", "Bond", "Shaken, not stirred"}
	m := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	agents := [][]string{james, m}

	// range over matrix first over, slices then values within each slice
	fmt.Println("Ranging over Agent Records in matrix:")
	for _, agent := range agents {
		fmt.Println("\nAgent Record: ", agent)
		fmt.Println("\nRanging over words in Agent:", agent[1])
		for _, word := range agent {
			fmt.Println(word)
		}
	}
}
