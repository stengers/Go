package main

import (
	"fmt"
	"math"
)

//a name is exported if it begins with a capital letter. Any "unexported" names are not accessible outside the package.
func main() {
	fmt.Println(math.Pi)
}
