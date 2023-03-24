package main

import "fmt"

func main()  {
	var result = adder(1,2,3)
	fmt.Println("completed ", result)
}

func adder(values ...int) int {
	var total int;
	for i := range(values) {
		total += values[i]
	}

	return total
}