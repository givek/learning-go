package main

import (
	"fmt"
	"os"
)

// Standard I/O
//
// Unix has the notion of three standard I/O streams.
// - Standard input
// - Standard output
// - Standard error (output)

func main() {
	fmt.Println("print a line to standard output.")
	fmt.Fprintln(os.Stderr, "print a error to standard error.")
}
