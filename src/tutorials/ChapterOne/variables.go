package main

import "fmt"

//var statement delcares a list of variables; as in function arguments lists, the type is last.
var c, python, java bool

func main() {
	//if an integer var is not given a value, it's default value is '0'
	var i int
	fmt.Println(i, c, python, java)
}
