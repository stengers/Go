package main

import "fmt"

func main() {
	//the variable declared determines the inferred type.
	i := 42			//integer inference
	f := 42.123		//float inference
	s := "42"		//string inference
	c := 0.867 + 0.5i	//complex128 inference
	
	fmt.Printf("%T, %T, %T, %T\n", i, f, s, c)
}
