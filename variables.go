package main

import (
	"fmt"
)

// either assign values to variables by defining type first
var a, b, c int = 1, 2, 3

// or by not defining type at all and doing it magically
var d, e, f, g = 1, true, 2.0, "Piyush"

func main() {
	// Inside function, short variable declarations can be used ":=" omitting var
	// However, this is not valid for variables outside function. ":=" or "var" with "="
	a1, a2 := 9, "Nine"
	var a3, a4 = 8, 7

	fmt.Println(a, b, c, d, e, f, g)
	fmt.Println(a1, a2, a3, a4)
}
