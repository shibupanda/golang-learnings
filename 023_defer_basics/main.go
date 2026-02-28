package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Case1 : sucess")
	if err := dowork(true); err != nil {
		fmt.Println("error", err)
	}

	fmt.Println("Case2 : fail early")

	if err := dowork(false); err != nil {
		fmt.Println("error", err)
	}


}

func dowork(success bool) error {
	// start message -> resource aquired
	// cleanup message -> resource released

	fmt.Println("start: resource acuired")
	defer fmt.Println("cleanup: resource released")

	if !success {
		return errors.New("something went wrong! returning early")
	}

	fmt.Println("work: doing something imp")
	fmt.Println("work: this work is done")

	return nil
}
