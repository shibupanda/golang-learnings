package main

import "fmt"

func main()  {
	points := map[string]int {
		"a" : 10,
		"b": 0,
	}

	fmt.Println("a", points["a"])
	fmt.Println("b", points["b"])
	fmt.Println("c", points["c"])

	valA, okA := points["a"]
	fmt.Println(valA, okA)

	val, ok := points["c"]
	fmt.Println(val, ok)

	if valC, okC := points["c"]; okC {
		fmt.Println(valC)
	}else {
		fmt.Println("C is not present", okC)
	}


}
