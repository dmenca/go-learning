package main

import "fmt"

type Language string

type Status string

func main() {

	language := Language("Java")
	fmt.Println(language)

	status := Status("OK")
	fmt.Println(status)

}
