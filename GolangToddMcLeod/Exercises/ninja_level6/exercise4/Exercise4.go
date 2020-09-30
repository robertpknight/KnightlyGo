package exercise4

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

// Exercise4 ... code for ninja level 6 exercise 4
func Exercise4() {
	fmt.Println("\nExercise 4:")

	r := person{
		first: "Rob",
		last:  "Knight",
		age:   27,
	}

	r.speak()
}

func (p person) speak() {
	fmt.Println("Hi, my name is", p.first, ", I am", p.age, "years old")
}
