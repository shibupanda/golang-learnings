package main

import "fmt"

func main() {

	// &x -> address of x
	// *p -> dereference (go to that address and read/write)
	score := 10
	fmt.Println("before ", score)
	addScore(&score)
	fmt.Println("after ", score)
}

func addScore(score *int) {
	*score = *score + 5

}