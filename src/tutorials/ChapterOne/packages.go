//every go program is made up of packages. Programs start running in package main.
package main

//The program is using packages with imported paths "fmt" and "math/rand".
import (
	"fmt"
	"math/rand"
)

//Main function, prints a random integer. The seed needs to be changed for a different random number
func main() {
	fmt.Println("My random number is", rand.Intn(100))
}
