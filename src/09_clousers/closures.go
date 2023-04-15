package main

import "fmt"

// Scope vs Lifetime

// Scope

// Scope is feature of the code (text) and checked at compile time.

// It's about who can access the variable.

// Lifetime

func doIt() *int {
	var b int

	return &b
}

// In a traditional language like c, this will blow up, as `b` is reference to a local variable of `doIt`,
// which will be cleared as soon as `doIt` finishes.

// In Go variable `b` can only be seen inside `doIt`, but it's value lives on.

// In Go when compiler see's such a case where a reference to a local variable is being returned from the function,
// the compiler does escape analysis. So that returned local value can live longer than the function.

// It is allocated on the heap.

// Closure

// A closure is when a function inside another function "closes over" one or more local variables of the outer function.

func fib() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b
		return b
	}
}

// The inner function gets a refernce to the outer function's variables.

// Those variables may end up with much longer lifetime than expected -- as long as there is a reference to the inner function.

// A closure is not just the function, the closure is what's being returned which may aslo includes any "closed over" variables.

// f := fib   // here f is a fucntion (fib function)
// f := fib() // here f is a closure (which includes the anonymous function as well as the closed scope.)

// A clouser has refernce to the closed variables, not the actual value.

func main() {
	f := fib()

	for x := f(); x < 100; x = f() {
		fmt.Println(x)
	}

	x, y := fib(), fib() // x and y get different local variables.

	fmt.Println(x(), x(), x())
	fmt.Println(y(), y(), y())

	println()

	anotherFunc()

	println()

	sliceOfFunc()

	println()

	problemSolved()
}

// Gotach's with clousers

func do(d func()) {
	d()
}

func anotherFunc() {
	for i := 0; i < 4; i++ {
		v := func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}
		do(v)
	}

}

// Here, the value of the `i` will change (0, 1, 2, 3), but the address will be the same.
// As `i` is the same variable, and clousers store the reference to the actual variable.

func sliceOfFunc() {
	s := make([]func(), 4)

	for i := 0; i < 4; i++ {
		s[i] = func() {
			fmt.Printf("%d @ %p\n", i, &i)

		}
	}

	for i := 0; i < 4; i++ {
		do(s[i])
	}
}

// This time the value as well as the address will same, as we are executing the functions after the completion of the loop.
// When the loop is completed the value of `i` becomes 4, and as closures store the reference to the actual variable,
// all of them print the value 4

func problemSolved() {
	s := make([]func(), 4)

	for i := 0; i < 4; i++ {

		// This is called closure capture.
		// here a new i2 is created for every iteration of the loop.
		// So it's value and refernce will be unique for every function created.
		i2 := i

		s[i2] = func() {
			fmt.Printf("%d @ %p\n", i2, &i2)

		}
	}

	for i := 0; i < 4; i++ {
		do(s[i])
	}
}

// So the Gotach is: As we are closing over by reference, if the closure is executed asynchronously (means later on),
// then it is possible that the variable we closed on mutates.
