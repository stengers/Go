package main

import "fmt"

func main() {
	//sum is an inferred integer of '0'
	sum := 0
	//can also be written as
	//var sum int

	//for i equals 0, i is less than 10, increment i by 1
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}
	fmt.Println("Total:", sum)
}
