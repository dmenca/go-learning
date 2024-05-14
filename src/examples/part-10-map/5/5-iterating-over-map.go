package main

import "fmt"

func main() {

	languages := map[string]int{}
	languages["java"] = 5
	languages["ruby"] = 4
	languages["go"] = 2

	for key, value := range languages {
		fmt.Println("Key:", key, "- Value:", value)
	}

	for _, value := range languages {
		fmt.Println( " Value:", value)
	}

	for key := range languages {
		fmt.Println(key)
	}

}
