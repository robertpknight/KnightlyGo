package exercises

import (
	"fmt"
)

type vehicle struct {
	doors  int
	colour string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

// Exercise3 ... code for ninja level 5 exercise 3
func Exercise3() {

	fmt.Println("\nThis is the solution to exercise 3:")

	monsterTruck := truck{
		vehicle: vehicle{
			doors:  2,
			colour: "red",
		},
		fourWheel: true,
	}

	simpsonsSedan := sedan{
		vehicle: vehicle{
			doors:  4,
			colour: "pink",
		},
		luxury: false,
	}

	// print out the entire struct value
	fmt.Println("\nmonster truck value: ", monsterTruck)
	fmt.Println("\nsimpsons sedan value: ", simpsonsSedan)

	// print specific values
	fmt.Println("\nmonster truck has", monsterTruck.doors, "doors")
	fmt.Println("\nsimpsons sedan is luxury:", simpsonsSedan.luxury)
}
