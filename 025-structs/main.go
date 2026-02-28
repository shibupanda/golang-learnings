package main

import "fmt"

type User struct {
	ID int
	Name string
	Email string
	Age int
}

func main() {
	u1 := User {
		ID: 123,
		Name: "Shibu",
		Email: "a@mail.com",
		Age: 36,
	}

	fmt.Println(u1, u1.Age)
}
