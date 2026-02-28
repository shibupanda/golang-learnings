package main

import "fmt"


type User struct {
	ID int
	Name string
	Email string
	Age int
}

func main() {
	u := User {
		Name: "Shibu",
		Age: 36,
	}

	fmt.Println("Before: ", u)
	u.Birthday()
	fmt.Println("After ", u)
}

func (u *User) Birthday() {
	u.Age++
}