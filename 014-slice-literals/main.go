package main

import "fmt"

func main() {
	// common collection type and it is dynamic and grow
	//[]type{...}

	results := []string{"shibu", "rupesh"}
	fmt.Println(results, results[0], results[len(results) - 1])
	results[1] = "priya"
	fmt.Println(results, results[0], results[len(results) - 1])

	var nums []int
	nums = append(nums, 10)
	nums = append(nums, 20)
	fmt.Println(nums)
}
