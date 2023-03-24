package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=123abc"

func main()  {
	fmt.Println("Welcome to URLS in golang")

	result, err := url.Parse(myurl)
	checkForNilError(err)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.RawQuery)

	qParams := result.Query()
	fmt.Println(qParams["coursename"])

	for _ ,val := range qParams {
		fmt.Println(val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=savan",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)
}

func checkForNilError(err error) {
	if err != nil {
		panic(err)
	}
}