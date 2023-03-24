package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const myurl = "http://localhost:8000"

func main() {
	fmt.Println("Welcome to my web server")

	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func checkForNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func PerformGetRequest() {
	response, err := http.Get(myurl + "/get")
	checkForNilError(err)

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	var responseString strings.Builder // Using "strings" library here, much powerful

	content, err := ioutil.ReadAll(response.Body)
	checkForNilError(err)

	byteCount, _ := responseString.Write(content)

	fmt.Println(byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest() {
	// JSON paylaod
	requestBody := strings.NewReader(`
		{
			"name": "Savan",
			"email": "savan@go.dev"
		}
	`)

	var readerString strings.Builder

	response, err := http.Post(myurl+"/post", "application/json", requestBody)
	checkForNilError(err)
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	readerString.Write(content)

	fmt.Println("Response is: ", readerString.String())
}

func PerformPostFormRequest() {

	// get data of txt file
	var textFileData = getDataFromFile()

	// form data
	data := url.Values{}
	data.Add("firstName", "Savan")
	data.Add("email", "savan@go.dev")
	data.Add("message", textFileData)

	response, err := http.PostForm(myurl+"/postform", data)
	checkForNilError(err)
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	checkForNilError(err)

	var readerString strings.Builder
	readerString.Write(content)

	fmt.Println(readerString.String())

}

func getDataFromFile() string {
	data, err := os.ReadFile("../16files/filego.txt")
	checkForNilError(err)

	return string(data)
}
