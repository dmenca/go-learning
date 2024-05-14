package main

import "fmt"

func main() {
	fmt.Println("Calling the Function: main")

	anotherFunction()
}

func anotherFunction()  {
	fmt.Print("Calling the function: anotherFunction()")
}
