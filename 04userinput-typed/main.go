package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println("Welcome to user input typed program")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your age => ")
	input, _ := reader.ReadString('\n')

	numRating, error := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if error != nil {
		fmt.Println("Error => ",error)
	} else {
		fmt.Println("New Age is ", numRating + 1)
	}

}