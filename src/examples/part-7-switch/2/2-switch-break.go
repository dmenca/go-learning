package main

import "fmt"

func main() {
	number := 10
	switch number {
	case 5:
		fmt.Println("The number is 5")
		// 找到一个匹配的条件会自动break
	case 10:
		fmt.Println("The number is 10")
	case 15:
		fmt.Println("This will not be checked!")
	}
}
