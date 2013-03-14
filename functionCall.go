package main

import (
	"fmt"
)

// multiple arguments can be written
func add(x int, y int) (int, int) {
	return x + y, 0
}

// when more than 1 argument of the function shares a type, then it can be
// written just once
func add2(x, y int) (int, string) {
	return x + y, "Piyush"
}

// Result parameters - If the result parameters are named, a return statement 
// without arguments returns the current values of the results
func add3(x, y int) (z int, s string) {
	z = x + y
	s = "Piyush"
	return
}

func main() {
	var x int
	var y int
	x = 1
	y = 2

	a, b := add(x, y)
	fmt.Println(a, b)

	a2, b2 := add2(x, y)
	fmt.Println(a2, b2)

	a3, b3 := add3(x, y)
	fmt.Println(a3, b3)
}
