package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	//the if statmenet can start with a short statement to exectue before the condition
	//variables declared by the statement are only in scope until the end of the if statement.
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
