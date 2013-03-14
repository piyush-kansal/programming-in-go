package main

import (
	"fmt"
	"math"
)

func main() {
	// In Go, you can refer to the exported names by using <package-name>.<first-letter-in-caps><rest-letters-upto-you-if-available>
	fmt.Println("Happy", math.Pi, "Day")

	// Error
	// fmt.Println("Happy", math.PI, "Day")

	// Error
	// fmt.Println("Happy", math.pi, "Day")
}
