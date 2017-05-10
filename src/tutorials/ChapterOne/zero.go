package main

import "fmt"

//number, boolean and string variables not declared with an explicit value are given zero, false or empty respectively.
func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
