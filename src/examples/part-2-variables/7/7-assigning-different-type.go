package main

import "fmt"

func main() {

	var num int = 10
	var price float64 = 15.10

	fmt.Println(num, price)
	// 不能像其他语言一样隐式转化, 只能显式转化
	// price = num
	price = float64(num)

	fmt.Println(price)
}
