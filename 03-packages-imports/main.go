package main

// import -> is used to include external packages in our program, 
// it allows us to use functions and types defined in other packages
import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand.Int31()) // Generates a random float64 number between 0.0 and 1.0
}
