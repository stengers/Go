package main

import "fmt"

func main() {
	//a defer statement defers the execution of a funciton until the surrounding function returns
	defer fmt.Println("world.")

	fmt.Printf("Hello ")
}
