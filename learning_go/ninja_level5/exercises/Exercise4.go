package exercises

import (
	"fmt"
)

// Exercise4 ... code for ninja level 5 exercise 4
func Exercise4() {

	// creating an anonamous struct
	fmt.Println("\nThis is the solution to exercise 4:")
	solarSystem := struct {
		name        string
		planets     int
		planetNames []string
		moons       map[string][]string
	}{
		name:        "sol",
		planets:     8,
		planetNames: []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"},
		moons: map[string][]string{
			"Earth": []string{"the Moon"},
			"Mars":  []string{"Phobos", "Deimos"},
			// skip other planets moons to save time
		},
	}

	// access field values in the solarsystem struct
	fmt.Println("How many planets does sol have:", solarSystem.planets)
	fmt.Println("What is Earth's moon called:", solarSystem.moons["Earth"][0])
	fmt.Println("How many moons does the 4th planet have:", len(solarSystem.moons[solarSystem.planetNames[3]]))

}
