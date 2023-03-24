package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Welcome to is else")

	var age = 20

	if age < 18 {
		fmt.Println("You are under age")
	} else if age < 60{
		fmt.Println("You are adult")
	} else {
		fmt.Println("You are senior citizen")
	}

	// Assign value to variable first, and then check it
	if num := 3; num > 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is greater than 10")
	}


}