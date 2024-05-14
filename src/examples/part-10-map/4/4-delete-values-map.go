package main

import "fmt"

func main() {
	languages := map[string]int{}
	languages["java"] = 5
	languages["ruby"] = 4
	languages["go"] = 2

	delete(languages, "ruby")
	delete(languages, "java")
	// no-op
	delete(languages, "python")

	fmt.Println(languages)
}
