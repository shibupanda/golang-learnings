package main

import "fmt"

func main() {
	scores := make([]int, 0, 5)
	fmt.Println(scores)
	scores = append(scores, 12)
	scores = append(scores, 12)
	scores = append(scores, 12)
	scores = append(scores, 12)
	scores = append(scores, 12)
	fmt.Println(scores)
	scores = append(scores, 50, 80, 90)
	fmt.Println(scores, len(scores), cap(scores))

	scores = append(scores, 50, 80, 90)
	fmt.Println(scores, len(scores), cap(scores))

	// NOTE: if it will exceed the capacity it will double the capacity

	a := []string{"I", "LOVE"}
	b := []string{"GO"}

	//...
	a = append(a, b...)
	fmt.Println(a)
	
}
