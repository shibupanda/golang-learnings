package main

import (
	"fmt"
)

func main() {
	isAdmin := true
	isLogged := false


	canOpenDashboard := isLogged && isAdmin
	canDelete := isAdmin || isLogged

	fmt.Println(canOpenDashboard, canDelete)

	age := 25
	isAdult := age > 18
	fmt.Println((isAdult))
}

