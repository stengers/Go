package main

import (
	"fmt"
	"math"
)

func main() {
	x, y := 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(f, x, y, z)
	fmt.Printf("f:%T, x:%T, y:%T, z:%T\n", f, x, y, z)
}
