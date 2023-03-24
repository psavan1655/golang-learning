package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://dummyjson.com/users"

func main()  {
	fmt.Println("Welcome to Web Requests")

	response, err := http.Get(url)
	
	checkForNilError(err)

	fmt.Printf("Type of response is: %T\n", response)

	defer response.Body.Close() // Never forget to close the connection.

	dataBytes,err := ioutil.ReadAll(response.Body)

	checkForNilError(err)

	content := string(dataBytes)
	fmt.Println("Data is : ", content)
	fmt.Println("Data is : ", response)

}

func checkForNilError(err error) {
	if err != nil {
		panic(err)
	}
}