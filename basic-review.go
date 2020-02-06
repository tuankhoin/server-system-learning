/** 
	This file can be used to demonstrate sample use of Golang basics. What is included:
	
	_Basic declaration
	_Pointers, arrays and struct
	_Slices and its operations
	
	Recommended background knowledge for Go: C and OOP
	Written by Tuan Khoi Nguyen, Feb 6th 2020.
**/

// What to begin

package main

import (
	"fmt"
	"math"
)

// Struct

type Location struct {
	X, Y int
}

// Function declaration
func main() {
	// defer: Only carry out the command after everything within function is finished
	// Order of multiple defers: last in-first out
	defer fmt.Println("Thank you")
	defer fmt.Println("If you are reading this, wish you a good day!")
	fmt.Println(Sqrt(2, 3))
	fmt.Printf("Comparing function to math.Sqrt: %f and %f\n", math.Sqrt(2), math.Sqrt(3))
	fmt.Println(compute(Sqrt))
	MultipleOps()
}

// Type declaration is reversed, and can return multiple values
// A void function doesn't have to declare return type
// Vars with same type can be declared together
func Sqrt(x, y float64) (float64, float64) {

	// Variable declaration always starts with ":="
	z1 := 0.0        //Automatic type initialization
	z2 := float64(1) // Direct type conversion

	// Pointers, and using them to change value
	p := &z1
	*p = 1.0

	// Control statement does not require brackets for arguments
	// "for" loop only need the condition. Other arguments are optional
	for ; z1 < 10; z1++ {
		z1 -= (z1*z1 - x) / (2 * z1)
	}

	//"for" loop in Go is "while" loop in C
	for z2 < 15 {
		z2 -= (z2*z2 - y) / (2 * z2)
	}
	return z1, z2
}

// A void function doesn't have to declare return type
func MultipleOps() {

	// Initialization will be made consecutively
	one, two, three := 1, 2, 3
	fmt.Printf("%d %d %d\n", one, two, three)

	// Struct components: Same as C
	loc := Location{69,96}
	fmt.Println(loc.X)

	// Array initialization
	array := [5]int{1, 2, 3, 4, 5}

	// Slice is a component of array, but more dynamically
	// Range selection: include first, excludes last
	var slice []int = array[1:4] // Try array[:4] as well

	// Slice can act as a reference to array
	slice[1] = 3
	fmt.Println(array)

	// Only difference to array is that array size can't be changed
	// Slice has length and capacity
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// Make a slice with a length of 0 and capacity of 5
	newSlice := make([]int, 0, 5)
	// Extend length (assigning values)
	newSlice = newSlice[:4]
	// Append will extend capacity and length by adding assigned values to slice
	newSlice = append(newSlice, 1, 2, 3, 4)
	// Drop values, will reduce length and capacity
	newSlice = newSlice[2:]
	// "for-range" loop for slices and maps
	// either i or v can be replaced with "_" if not necessary
	for i, v := range newSlice {
		fmt.Printf("newSlice[%d]=%d\n", i, v)
	}

	// Maps and keys
	var m = map[string]int{
		"one":1,
		"two":2,
		"three":3,
	}
	// Delete
	delete(m, "three")
	fmt.Println(m)

	// Check if key has a value
	ok := m["three"]
	fmt.Println("Present?", ok)
}

// Function as a variable
func compute(fn func(float64, float64) (float64, float64)) (float64,float64) {
	return fn(5,6)
}

