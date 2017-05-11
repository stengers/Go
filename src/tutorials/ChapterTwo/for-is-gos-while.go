package main

import "fmt"

//If the init and/or Post are left out, the semicolons can be dropped. This is essentially a while statement in Go.
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println("Total:", sum)
}
