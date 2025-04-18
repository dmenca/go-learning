package main

import "fmt"

func main() {
	languages := []string{"Ruby", "Java", "Pony", "Go", "Erlang"}

	for i, language := range languages {
		fmt.Println("index: ", i, " language: ", language)
	}

	for index := range languages {
		fmt.Println("language: ", languages[index])
	}
	
	

	fmt.Println(languages)
}
