package exercises

import (
	"fmt"
)

// Exercise2 ... code for ninja level 5 exercise 2
func Exercise2() {

	fmt.Println("\nThis is the solution to exercise 2:")
	// person type declared in Exercise1
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

	// create map of type person
	people := map[string]person{
		rob.lastName: rob,
		liz.lastName: liz,
	}

	fmt.Println(people["Knight"].firstName + " likes the icecream flavours: ")
	for _, v := range people["Knight"].favouritIceCreams {
		fmt.Println("\t" + v)
	}
}
