package main

import "fmt"

//a var declaration can included initializers, one per variable. An initializer is a string, integer, boolean, etc.
var i, j int = 1, 2

func main() {
	//if an initializer is present, the type can be omitted; the variable will take the type of the initializer.
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
