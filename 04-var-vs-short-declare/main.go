package main

import (
	"fmt"
)


func main() {
	var city string
	city = "New York"

	var name = "shibu"; //infer type is string, so we can omit the type declaration

	// := is a short variable declaration, 
	// it is used to declare and initialize a variable in one line,
	// it can only be used inside a function

	x := 1000 // infer type is int, so we can omit the type declaration

	x = x + 500 // we can reassign a new value to the variable

	a, b := 100, 50 // we can declare and initialize multiple variables in one line

	fmt.Println(city, name, x, a, b)
}
