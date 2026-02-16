package main

import "fmt"

func main() {
	//map[keyType]valueType
	ages := map[string]int {
		"shibu": 35,
		"priyanks": 30,
	}

	fmt.Println(ages["shibu"], len(ages))

	for k, v := range ages {
		fmt.Println(k, v)
	}

	//var scores map[string]int
	scores := make(map[string]int)

	scores["math"] = 10

	fmt.Println(scores)

	delete(scores, "math")
	fmt.Println(scores)
}
