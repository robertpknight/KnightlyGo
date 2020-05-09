package main

import (
	"fmt"
)

func main() {
	a := 42

	// printing value and pointer
	fmt.Println(a)
	fmt.Println("memory address of pointer:", &a)

	// print types
	fmt.Printf("%T\n", a)
	fmt.Printf("Type of the address: %T\n", &a)

	// *int is pointer to an int. This is a different type to int.
	// Cannot assign the address to int, they are different types
	// var b int = &a

	// can assign the pointer if we user pointer to an int
	var b *int = &a
	fmt.Println("&a assigned to b:", b)

	// using * on an address - this operator * de-references the address
	// returning the value 42
	fmt.Println(*b)

	// you can chain a de-referencing of the address
	fmt.Println(*&a)

	// you can change the value of "a" by re-assigning a value to de-referencing b or &a
	*b = 8
	fmt.Println(a)

	// Using pointers
	// 1. Pass an address if the data at the address is large
	// 2. When you need to change a value
}
