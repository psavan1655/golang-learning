package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"courseName"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`              // "-" removes the key in JSON
	Tags     []string `json:"tags,omitempty"` // Help to remove the key, if value is nil
}

func main() {
	fmt.Println("Welcome to JSON world of golang")
	// EncodeJson()
	DecodeJson()
}

func checkForNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func EncodeJson() {
	srpCourses := []course{
		{"ReactJS Bootcamp", 299, "srp.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "srp.in", "bcd123", []string{"web-dev", "jsx"}},
		{"Go Bootcamp", 599, "srp.in", "xyz123", nil},
	}

	// Package this data as JSON data
	finalJson, err := json.MarshalIndent(srpCourses, "", "\t")
	checkForNilError(err)

	fmt.Printf("Final JSON is: %s", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"courseName": "ReactJS Bootcamp",
		"price": 299,
		"website": "srp.in",
		"tags": ["web-dev","js"],
		"message": "hey buddy"
	}
	`)

	var srpCourse course

	isvalid := json.Valid(jsonDataFromWeb)

	if isvalid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &srpCourse)
		fmt.Printf("%#v\n", srpCourse)
	} else {
		fmt.Println("invalid JSON")
	}

}
