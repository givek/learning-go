package main

import "fmt"

// if-then-else
// All if-then-else statements require braces, even if we have one statement inside the condition.

func main() {

	var a = 1
	var b = 2

	if a == b {
		fmt.Println("a == b")
	} else {
		fmt.Println("a != b")
	}

	// if statement can also start with a short declaration (:=) or statement.
	//
	// if (declaration); (condition) {
	//     ...
	// }

	if a = 0; a == 0 {
		fmt.Println("a == 0")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	colors := []string{"red", "green", "blue"}

	// one var: i is an index 0, 1, 2, ...
	for i := range colors {
		fmt.Println(i)
	}

	// two var: i is an index, v is the value.
	// the value (v) is a copy, not a reference
	for i, v := range colors {
		fmt.Println(i, v)
	}

	// loops can have labels.
outer:
	for i := range colors {
		for j := range colors {
			if i == j {
				continue outer
			}
		}
	}
}
