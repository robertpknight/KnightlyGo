package exercise8

import (
	"fmt"
	"math/rand"
	"time"
)

// Exercise8 ... code for ninja level 6 exercise 8
func Exercise8() {
	fmt.Println("\nExercise 8:")
	C3P0 := getBestYear()
	R2D2 := getBestYear()

	fmt.Println("C3P0 thinks the best year is:", C3P0())
	fmt.Println("R2D2 thinks the best year is:", R2D2())

}

func getBestYear() func() int {
	// randomm see is different for each value of the func
	rand.Seed(time.Now().Unix()) // initialise sudo random generator
	return func() int {
		l := []int{1946, 1999, 1969, 1993}
		return l[rand.Intn(len(l))]
	}
}
