package main

import "fmt"

func main() {
	value := Sum(9, 7)
	fmt.Println("Total value:", value)
}

func Sum(first, second int) int {
	totalValue := first + second
	return totalValue
}
