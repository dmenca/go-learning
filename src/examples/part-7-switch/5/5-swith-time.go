package main

import (
	"fmt"
	"time"
)

func main() {
	// 在同一个 `case` 语句中，你可以使用逗号来分隔多个表达式。
	// 在这个例子中，我们还使用了可选的 `default` 分支。
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// 不带表达式的 `switch` 是实现 if/else 逻辑的另一种方式。
	// 这里还展示了 `case` 表达式也可以不使用常量。
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}
