package main

import (
	"fmt"
)

// struct is an aggregate type
type person struct {
	name     string // fields explicity specified in indentifier list
	location string
	age      int
}

// embedded person struct
type secretAgent struct {
	person                   // anonamous field or embedded field, unqualified type name acts as field name
	name, catchPhrase string // same types can be in one declaration, or can choose to declare each field
	licence           bool
}

func main() {
	structExample()
	fmt.Println()
	anonamousStructExample()
}

func structExample() {
	fmt.Println("Below are some structs:")

	r := person{
		name: "Rob Knight",
	}

	p := secretAgent{
		person: person{
			name: "James Bond",
			age:  32,
		},
		name:        "007",
		catchPhrase: "Shaken, not stirred",
		licence:     true,
	}

	fmt.Println(r)
	// unused field defaults to empty or 0
	fmt.Println(r.name, r.age, r.location)

	fmt.Println(p)
	// to prevent a name collision we can use the full name spacing
	// otherwise the inner type gets promoted to the outer type in the case of age
	fmt.Println(p.person.name, p.name, p.age, p.licence, p.catchPhrase)
}

func anonamousStructExample() {
	fmt.Println("Below is an anonamous struct:")
	r := struct {
		name     string
		location string
		age      int
	}{
		name:     "Rob Knight",
		location: "London",
		age:      27,
	}

	fmt.Println(r)
}
