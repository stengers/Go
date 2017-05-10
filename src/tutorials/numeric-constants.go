package main

import "fmt"

//an untyped constant takes the type needed by it's context
const (
	//Create a huge number by shifting a 1 bit left 100 places.
	//a.k.a., the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	//Shift it right again 99 places, so we end up with 1 << 1, or 2.
	Small = Big >> 99
	//Playing around with >> bit manipulation
	Zero = 1 >> 1 //0
	One = 1 << 1 //2
	Two = 1 << 2 //4
	Three = 1 << 3 //8
	Four = 1 << 4 //16
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func needNum(x int) int{
	return x
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	//bit manipulation
	fmt.Println(needNum(Zero))
	fmt.Println(needNum(One))
	fmt.Println(needNum(Two))
	fmt.Println(needNum(Three))
	fmt.Println(needNum(Four))
}
