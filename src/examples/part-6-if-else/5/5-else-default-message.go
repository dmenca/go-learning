package main

import "fmt"

func main() {
	// Go 没有三目运算符
	// if statement
	if number := 10; number > 5 {
		fmt.Println("Will not be printed!")
	} else if number < 7 {
		fmt.Println("Will not be printed!")
	} else if number < 9 {
		fmt.Println("Will not be printed!")
	} else {
		fmt.Println("Message Default!")
	}
}
