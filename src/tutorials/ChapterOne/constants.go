package main

import "fmt"

//decalre a constant with the const declaration, can include or exclude type. It will become the type needed based on inference of the variable.
const Pi = 3.14
const One int = 1

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "day!")
	
	const Truth = true
	fmt.Println("Go is awesome?", Truth)
	fmt.Println(One)
}
