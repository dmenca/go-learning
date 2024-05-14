package main

import "fmt"

func main() {
	languages := map[string]int{"java": 4, "ruby": 6, "go": 2}
	fmt.Println(languages)

	ele := languages["ruby"]
	fmt.Println(ele)

	// 当从一个 map 中取值时，还有可以选择是否接收的第二个返回值，该值表明了 map 中是否存在这个键。
	// 这可以用来消除 `键不存在` 和 `键的值为零值` 产生的歧义，
	_, exist := languages["python"]
	fmt.Println(exist)

	_, ex := languages["ruby"]
	fmt.Println(ex)
}
