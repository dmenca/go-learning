package main

import "fmt"

func main() {
	languages := make([]string, 5)
	languages[0] = "Ruby"
	languages[1] = "JavaScript"
	languages[2] = "Java"
	languages[3] = "Scala"
	languages[4] = "Go"

	// slice 支持通过 `slice[low:high]` 语法进行“切片”操作。
	dynamicLanguages := languages[0:2]
	fmt.Println(dynamicLanguages)

	staticTyped := languages[2:5]
	fmt.Println(staticTyped)

	// 切片返回的是原slice的视图
	staticTyped[0] = "Python"
	fmt.Println("staticTyped: ", staticTyped)
	fmt.Println("languages: ", languages)

	// 这个 slice 包含从 `s[0]` 到 `s[5]`（不包含 5）的元素。
	l := languages[:5]
	fmt.Println("sl2:", l)

	// 这个 slice 包含从 `s[2]`（包含 2）之后的元素。
	l = languages[2:]
	fmt.Println("sl3:", l)

}
