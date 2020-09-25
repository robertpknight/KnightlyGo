package exercise2

import (
	"fmt"
)

// Exercise2 ... code for ninja level 6 exercise 2
func Exercise2() {
	fmt.Println("\n\nExercise 2:")

	slice := []int{1, 2, 3, 4, 8}
	fmt.Println("Foo of ", slice, "is:", foo(slice...))

	fmt.Println("Bar of ", slice, "is:", bar(slice))
}

func foo(i ...int) int {
	t := 0
	for _, v := range i {
		t += v
	}
	return t
}

func bar(i []int) int {
	t := 0
	for _, v := range i {
		t += v
	}
	return t
}
