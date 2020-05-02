package main

import (
	"fmt"
)

type person struct {
	first, last string
}

type secretAgent struct {
	person
	license bool
}

// func (r receiver) identifier(parameters) (returns(s))
func (s secretAgent) speak() {
	fmt.Println("the name's", s.last+",", s.first, s.last)
}

// method repeated but with different functionality
func (p person) speak() {
	fmt.Println("my name is", p.first)
}

// interfaces let us define behaviour - polymorphism
// a value can be of more than one type
// any type with method speak() is of type human
type human interface {
	speak()
}

// this function takes values of type human
// the functionality can change depending on
// what other types the value has
func printHuman(h human) {
	fmt.Println("\nprinting value of type human:", h)
	fmt.Printf("\tThis human is of type:%T", h)

	// using type in switch statement to define behaviour
	switch h.(type) {
	// checking tyoe person
	case person:
		fmt.Println("\n\tthis person has first name: " + h.(person).first + " and " + h.(person).last) // type person is asserted here
	case secretAgent:
		fmt.Println("\n\tthis person has first name: ", h.(secretAgent).first,
			"and "+h.(secretAgent).last+". License to kill status:",
			h.(secretAgent).license,
		)
	}
}

func main() {
	agent := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		license: true,
	}

	rob := person{
		first: "Rob",
		last:  "Knight",
	}

	fmt.Println(agent)

	// call method on value of type secretAgent
	agent.speak()
	// call method on value of type person
	rob.speak()

	// call function which can take multiple types using polymorphism
	// this means a function can take lots of types as long as
	// the human interface is implemented
	printHuman(agent)
	printHuman(rob)
}
