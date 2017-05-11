package main

import (
	"fmt"
	"math"
)

//Both calls to pow are executed and returned before the call to fmt.Println in main begins.
func pow(x, n, lim float64) float64 {
	if v:= math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	//can't use v, it's out of scope
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
		pow(3, 2, 12),
		pow(3, 2, 8),
	)
}
