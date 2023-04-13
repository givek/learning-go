package main

import (
	"fmt"
	"os"
)

// The defer statement captures a function call to run later.

// example

// We need to ensure the file closes

func main() {
	f, err := os.Open("my_file.txt")

	if err != nil {
		// ...
	}

	defer f.Close()

	// and do somthing with the file.
}

// The call to Close is guaranteed to run at function exit.
// (don't defer closing the file until we know it really opened.)

// If we have more than 1 defer statements, on function exit, they will be executed in opposite order.

// NOTE: defer operates on function scope, which is a good thing, but not always.

func notUsefulDefer() {
	for i := 0; i < len(os.Args); i++ {
		f, err := os.Open(os.Args[i])

		if err != nil {
			// ...
		}

		defer f.Close()

		// do something
	}
}

// The deferred calls to f.Close() must wait until function exit.
// (we might run out of file descriptors before that)

// So, in such case instead of using defer we should close the file immediately, after use.

// defer copies arguments to the deferred call

func something() {
	a := 10

	defer fmt.Println(a)

	a = 15

	fmt.Println(a)
}

// prints 10, 15 - defer copies arguments.

// defer with naked return

// A defer statement runs before the return is done.

// When return types are named, they act like local variable to the function and can be assined to and are return automatically. (naked return)

func doIt() (a int) {

	// As the statement only execute functions,
	// if we need do an assingment we have to create an anonymous function for that and envoke it.

	defer func() { a = 2 }()

	a = 1

	return
}

// returns 2, as `a` is the return value to the fucntion.
