package main

import "fmt"

func main()  {
	fmt.Println("Welcome to the world of pointers\n")

	var b = "hey"

	c := &b
	fmt.Println(b)
	fmt.Println(c)
	
	*c = "demo"
	fmt.Println(b)
	fmt.Println(c)

}