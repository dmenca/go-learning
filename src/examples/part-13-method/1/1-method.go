package main

import (
	"fmt"
	"strings"
)

type Description string

// 方法声明 声明一个接收器(receiver)为Description类型的方法。允许方法在被调用时操作接收器所属的实例
func (d Description) Upper() string {
	return strings.ToUpper(string(d))
}

func main() {

	description := Description("My Go special description")
	upper := description.Upper()
	fmt.Println(upper)

}
