package main

import "fmt"

func main()  {
	var username string = "Savan"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)
}