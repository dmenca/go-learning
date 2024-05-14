package main

import "fmt"

func main() {
	fmt.Println("Executing inside the main function")

	AnotherFunction()
}

func AnotherFunction() {
	fmt.Println("Executing inside the another function")
}
