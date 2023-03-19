package main

// Function are first-class objects.

// Function Scope
//
// Almost anything can be defined inside a function.
//
// Methods cannot be defined in a function (only in package scope)

// A function declaration lists `formal` parameters.

func do(a int, b int) int // here a & b are formal parameters

// A function call has `actual` parameters. (also know as arguments)

// result := do(1, 2) // 1, 2 are actual parameters

// Parameter Passing

// By value

// - numbers
// - bool
// - arrays
// - structs

// By reference

// - things passed by pointer (&x)
// - strings (but they are immutable)
// - slices
// - maps
// - channels

// What actually happens

// Acutally all parameters are passed by copying something (i.e. by value)

// For By value the actual data structure is copied.
// For By reference the descriptor is copied.

// If the thing copied is a pointer or descriptor, then the shared backing store (array, hash table, etc.) can be changed through it.

// Thus we think of it as "by reference".

// Return Values

// Functions can have multiple return values

func doIt(a int, b int) int {
	return 1
}

func doItAgain(a int, b int) (int, error) {
	return 1, nil
}

// Every return statement must have all the values specified.
