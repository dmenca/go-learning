package main

import "fmt"

// 隐式返回值
func main() {
	total := sum(8, 12)
	
	fmt.Println("Total value:", total)

	fmt.Println()

}

func sum(first, second int) (amount int) {
	amount = first + second
	return
}
