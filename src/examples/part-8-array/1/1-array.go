package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var languages [5]string
	languages[0] = "Go"
	languages[1] = "Ruby"
	languages[2] = "Pony"
	languages[3] = "Erlang"
	languages[4] = "Java"

	fmt.Println(languages)

	// json.Marshal方法返回json字符串的字节数组
	jsonString, err := json.Marshal(languages)
	if err != nil {
		fmt.Println("json error", err)
		return
	}
	// 通过string函数将字节数组转换为字符串
	fmt.Println(string(jsonString))
}
