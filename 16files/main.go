package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main()  {
	fmt.Println("Welcome to files")

	content := "This needs to go in a file - savan"

	file, err := os.Create("./filego.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Length is: ", length)
	defer file.Close()

	readFileData(file.Name())
}

func readFileData(filename string) {
	dataByte , err := ioutil.ReadFile(filename)

	tmp, _ := os.ReadFile(filename)

	fmt.Println("Demo \n", string(tmp))

	if err != nil {
		panic(err)
	}

	fmt.Println("Dat in file is \n", string(dataByte))
}