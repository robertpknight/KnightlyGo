package exercises

import (
	"fmt"
)

type person struct {
	firstName, lastName string
	favouritIceCreams   []string
}

// Exercise1 ... code for ninja level 5 exercise 1
func Exercise1() {

	fmt.Println("\nThis is the solution to exercise 1:")

	rob := person{
		firstName:         "Rob",
		lastName:          "Knight",
		favouritIceCreams: []string{"mango", "stracciatella"},
	}

	liz := person{
		firstName:         "Liz",
		lastName:          "Lord",
		favouritIceCreams: []string{"chocolate", "jaffa cake"},
	}

	// create slice of type person
	people := []person{rob, liz}

	for _, p := range people {
		fmt.Println("\n" + p.firstName + " likes icecream flavours:")
		for _, flavour := range p.favouritIceCreams {
			fmt.Println("\t" + flavour)
		}

	}

}
