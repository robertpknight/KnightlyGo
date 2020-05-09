package main

import (
	"fmt"
)

type person struct {
	name   string
	gender string
	sex    int
}

func main() {
	fmt.Println("Exercise 1: ")
	h := "Hello world"
	fmt.Println(h)
	fmt.Println("Hello world address:", &h)
	fmt.Printf("Hello world address is of type: %T\n", &h)

	fmt.Println("\nExercise 2: ")
	jb := person{
		name:   "Joe Bloggs",
		gender: "male",
		sex:    27,
	}

	fmt.Println("Joe Bloggs is:", jb.gender)

	mutate_sex(&jb)
	fmt.Println("Joe Bloggs is:", jb.gender)
}

func mutate_sex(p *person) {
	fmt.Println("Swapping gender of person at address:", p)
	if (*p).gender == "male" {
		(*p).gender = "female"
		// can also use p.gender = "female"
	}
}
