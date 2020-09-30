package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Writer interface implemented if the Write method is defined
	fmt.Fprintln(os.Stdout, "Hello Writer")
	io.WriteString(os.Stdout, "Hello Writer")
}
