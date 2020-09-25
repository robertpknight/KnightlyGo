package main

import (
	"fmt"
)

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 10}
	s := sum(ii...)
	fmt.Println("sum all numbers in slice:", ii, "\nsum:", s)

	// using callback with sum function
	s2 := applyEven(sum, ii...)
	fmt.Println("apply sum function to even numbers only:", s2)

	s3 := applyOdd(sum, ii...)
	fmt.Println("apply sum function to odd numbers only:", s3)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

// callback
// takes function as parameter
// and variadic parameter of int
func applyEven(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func applyOdd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 1 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
