package main

import "fmt"

func main() {
	fmt.Println("Structures in Golang")

	savan := User{"Savan", "savan@go.dev", true, 18}
	fmt.Println(savan)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
