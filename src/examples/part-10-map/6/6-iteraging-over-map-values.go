package main

import "fmt"

func main() {

	languages := map[string]int{}
	languages["java"] = 5
	languages["scala"] = 4
	languages["go"] = 3

	for _, number := range languages {
		fmt.Println("Value:", number)
	}
}
