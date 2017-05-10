package main

//import "fmt"
//import "math"
//grouping imports rather than importing each one on a separate line
import (
	"fmt"
	"math"
)

//%g grabs the output of the function after the ','
func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
