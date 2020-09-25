package main

import (
	"fmt"
)

func main() {
	defer foo()
	bar()
	// foo prints after bar
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
