package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// go don't use exception for normal failures
	// functions -> return errors as normal return values

	// val, err := something()
	/*
	if err != nil {
	}
	*/
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error{
	input := "30"
	level, err := parseLevel(input)
	if err != nil {
		return err
	}

	fmt.Println("selected level", level)
	return nil
}

func parseLevel(s string) (int, error) {
	//(val, err)
	// if nil error -> sucess
	// err != nil -> failure

	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("level must be a number")
	}

	if n < 1 || n > 5 {
		return 0, fmt.Errorf("Level must be 1 and §")
	}

	return n, nil
}