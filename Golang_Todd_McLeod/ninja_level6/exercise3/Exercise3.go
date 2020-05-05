package exercise3

import (
	"fmt"
)

// Exercise3 ... code for ninja level 6 exercise 3
func Exercise3() {
	fmt.Println("\nExercise 3:")

	slice := []int{1, 2, 3, 4, 8}
	defer fmt.Println("Foo of ", slice, "is:", foo(slice...))
	fmt.Println("... deferred Foo")
	fmt.Println("Bar of ", slice, "is:", bar(slice))
	fmt.Println("... Now print Foo")
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
