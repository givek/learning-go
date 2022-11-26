package main

import "fmt"

// packages
//
// Everything lives in a package.
// Every standalone program has a main package.
// You can declare anything at package scope. But you can't you the short declaration operator.
//
// Every name that is capitalized is exported.

func main() {
	var a = 0
	a = 8

	fmt.Println(a)

	{
		var x = 0
		fmt.Println(x)
	}
}

// Imports
//
// Each source file in your package must import what it needs.
//
// No cyclic imports, A package "A" cannot import a package that imports "A".

// Items within the package get initialized before main()


// Declarations
//
// There are six ways to introduce a name:
// 1. Constant declarations with `const`.
// 2. Type declarations with `type`.
// 3. Variable declarations with `var`.
// 4. Short hand, initialized variable declaration of any type `:=`.
// 5. Function declarations with `func`.
// 6. Formal parameters and named returns of a function.
