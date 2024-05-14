package main

import "fmt"

func main() {

	SumAndPrint(3, 6)
}

func SumAndPrint(first int, second int) {
	sum := first + second
	fmt.Println(sum)
}
