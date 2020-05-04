package exercises

import (
	"fmt"
	"strconv"
)

// Exercise1 ... code for ninja level 6 exercise 1
func Exercise1() {
	fmt.Println("Exercise 1:")
	fmt.Println("foo print:", foo(42))

	i, s := bar(69)
	fmt.Printf("bar prints: %v type: %T", i, i)
	fmt.Printf("\nbar prints: %v type: %T", s, s)
}

func foo(i int) int {
	return i
}

func bar(i int) (int, string) {
	// string conversion from int
	return i, strconv.Itoa(i)
}
