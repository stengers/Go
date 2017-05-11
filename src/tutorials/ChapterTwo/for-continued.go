package main

import "fmt"

//for loops consist of three parts. The init, the condition, and the post. The init and post statements are OPTIONAL.
func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println("Total:", sum)
}
