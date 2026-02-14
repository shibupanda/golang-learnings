package main

import (
	"fmt"
)

func main() {
	score := 72

	if score > 90 {
		fmt.Println("A")
	}else if score >= 75 {
		fmt.Println("B")
	}else {
		fmt.Println(("C"))
	}
}
