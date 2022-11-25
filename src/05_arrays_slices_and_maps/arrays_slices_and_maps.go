package main

import "fmt"

// composite types
//
// - string
// - array (fixed size)
// - slice (variable length)
// - map


// Arrays
// Arrays are typed by size, which is fixed at complie time.
// The size of an array is part of its type. The types [10]int and [20]int are distinct.
// It can be use usefull to store fixed size data, where values do not change.

// All these are equivalent.

var arr1 [3]int
var arr2 = [3]int{0, 0, 0} // providing initial values
var arr3 = [...]int{0, 0, 0} // sized by initializer

var arr4 = [3]int{1, 2, 4}

var arr5 [4]int

// func main() {
//     fmt.Println(arr1) // [0 0 0]
//     arr1 = arr4 // elements are copied
//     fmt.Println(arr1) // [1 2 3]
// 
//     // arr5 = arr4 // !TYPE MISMATCH
// }


// Slices
// Slices have variable lenght, backed by some array;
// Slices are passed by reference.

// The slice descriptor stores a pointer to the actuall array, the lenght of the array and the capacity.

var slice1 []int // nil, no storage
var slice2 = []int{1, 2} // initialized

// func main() {
//     fmt.Println(slice1) // []
// 
//     // The return value of the append function needs to be assigned to the orignal variable, 
//     // as a new array can be created on append.
//     slice1 = append(slice1, 1) // append to nil, OK
//     slice2 = append(slice2, 3)
// 
//     fmt.Println(slice1) // [1]
//     slice1 = slice2 // OK, as assignment only copies the descriptor.
//     fmt.Println(slice1) // [1 2 3]
// 
//     fmt.Println(slice1[0:2]) // [inclusive:exclusive] and the size is (2 - 0).
// }


// Maps
// Maps are dictionaries, (key: value) pairs.
// Maps are passed by reference.
// The type used for the key must have == and != defined.

// Reading from a nill map is OK, but inserting will panic.

var map1 map[string]int                       // nil, no storage.
var map2 = make(map[string]int)               // non-nil, but empty.
var map3 = map[string]int{"and": 1, "the": 0} // non-nil, non-empty.

func main() {
    fmt.Println(map1["the"]) // 0
    fmt.Println(map2["the"]) // 0

    // map1["and"] = 1 // PANIC! As map1 is nil.

    val1, ok1 := map1["the"]
    fmt.Println(val1, ok1) // 0, false. As the map is empty (returned 0 is the default value).

    val3, ok3 := map3["the"]
    fmt.Println(val3, ok3) // 0, true
}
