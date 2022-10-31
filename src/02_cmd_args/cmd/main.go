package main

import (
	"fmt"
	"os"
)

// In go, the main function does not take command line args.
func main() {
	programName := os.Args[0] // 0th arg is always the program name

	fmt.Printf("Hello %s \n", programName)

	if len(os.Args) > 1 {
		fmt.Println("Hello,", os.Args[1])
	}
}
