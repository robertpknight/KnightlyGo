package exercises

import (
	"fmt"
)

// Exercise8 ... exercise 8 solution
func Exercise8() {
	fmt.Println("\nExercise 8:")
	// create map
	agents := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	for k, slice := range agents {
		fmt.Println("key:", k)
		for i, v := range slice {
			fmt.Println("\tvalue: ", v, "index:", i)
		}

	}
}
