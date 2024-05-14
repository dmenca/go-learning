package main

import "fmt"

func main() {
	printNumbers(1, 21, 3)
	printStr("1", "2", "3")
}

func printNumbers(numbers ...int) {
	fmt.Println(numbers)
}

func printStr(str ...string) {
	fmt.Println(str)
}
