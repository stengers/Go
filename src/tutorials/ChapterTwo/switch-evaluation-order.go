package main

import (
	"fmt"
	"time"
)

//switch cases evaluate from top to bottom, stopping when a case succeeds.
func main() {
	fmt.Println("When's Saturday?")
	switch today := time.Now().Weekday(); time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	case today + 3:
		fmt.Println("Too far away.")
	}
}
