package main

import (
	"fmt"
	"math"
)

//using Newton's method to estimate the square root
func sqrt(x float64) float64 {
	z := 1.
	for i := 0; i < 2; i++ {
		z = z - ((math.Pow(z, 2) - x) / (2 * z))
	}
	return z
}

//comparing guess to actual sqrt
func main() {
	x := 3.
	fmt.Println(
		"Guess: ", sqrt(x),
		"\nActual:", math.Sqrt(x),
	)
}
