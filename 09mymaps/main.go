package main

import "fmt"

func main() {
	fmt.Println("Welcome to the world of Maps")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Printf("Type of languages %T\n", languages)
	fmt.Println(languages)

	delete(languages, "RB")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Println(key, " for ", value)
	}
}
