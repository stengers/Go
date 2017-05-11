package main

import "fmt"

func main() {
	fmt.Println("Counting down from 10")
	//the for loop counts down from 10, the defer should reverse the order
	for i := 10; i > 0; i-- {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
