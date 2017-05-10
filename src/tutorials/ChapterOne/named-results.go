package main

import "fmt"

//go rutern values can be named. They are treated as variables defined at the top of the function.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	//a return statement without arguemtns returns the named return values (known as a "naked" return).
	return
}

func main() {
	fmt.Println(split(17))
}
