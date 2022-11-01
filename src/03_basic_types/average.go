package main

import (
	"fmt"
	"os"
)

func main() {

	totalSum := 0.0

	n := 0

	for {
		val := 0.0

		_, err := fmt.Fscanln(os.Stdin, &val)

		if err != nil {
			break
		}

		totalSum += val
		n++
	}

	if n == 0 {
		os.Exit((-1))
	}

	fmt.Println("The average is", totalSum/float64(n))

}
