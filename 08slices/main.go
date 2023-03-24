package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Welcome to the world of slices")

	var slice = []string{"Apple", "Tomato", "Peach"}

	fmt.Println(slice)
	
	slice = append(slice, "Mango")
	fmt.Println(slice)

		
	slice = slice[1:2]
	fmt.Println(slice)
	
	newSlice := make([]int, 1)
	newSlice[0] = 123
	fmt.Println(newSlice)
	
	newSlice = append(newSlice, 234,264,654,100)
	fmt.Println(newSlice)
	
	sort.Ints(newSlice)
	fmt.Println(newSlice)

	// Remove value from slices based on index
	var courses = []string{"reactjs", "javascript","swift", "python","ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}