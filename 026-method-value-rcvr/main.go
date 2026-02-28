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

	fmt.Println(u.Intro())
}

// val receiver receives a copy of the user
func (u User) Intro() string {
	return fmt.Sprintf("Hi, I am %s", u.Name)
}
