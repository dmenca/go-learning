package main

import "fmt"

type Minute int
type Hour int

func main() {
	minutes := Minute(70)
	hour := Hour(10)

	// 报错 不同类型不能进行比较
	if minutes > hour {
		fmt.Println("This will never be executed")
	}

}
