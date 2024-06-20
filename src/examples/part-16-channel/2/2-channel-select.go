package main

import (
	"fmt"
	"time"
)

//queue是一个无缓冲（unbuffered）的channel。在Go语言中，无缓冲的channel要求发送和接收必须同时准备好，才能进行数据传递。也就是说，发送操作会阻塞，直到有接收者来读取数据

var queue = make(chan int)

func main() {

	select {
	// 在这里，queue <- 1操作会阻塞，因为没有任何接收者读取这个值。由于没有其他接收者，5秒后time.After case会被触发，输出“time out”。
	case queue <- 1:
		fmt.Println("queue 1")
	case <-time.After(5 * time.Second):
		fmt.Println("time out")
	}

	//为了让发送操作不阻塞，有两种方法：

	//添加接收者：在另一处代码中添加一个接收者，以便在发送时可以读取数据。

	//启动了一个goroutine来从queue中读取数据。这样，发送操作queue <- 1可以成功完成，输出“queue 1”。
	go func() {
		val := <-queue
		fmt.Println("Received:", val)
	}()

	select {
	case queue <- 2:
		fmt.Println("queue 2 ")
	case <-time.After(5 * time.Second):
		fmt.Println("queue 2 time out")
	}
	// 使用带缓冲的channel：使用带缓冲的channel可以在不立即需要接收者的情况下允许一些数据发送。

	var queue_buffer = make(chan int, 1) // 创建一个带缓冲的channel
	select {
	case queue_buffer <- 1:
		fmt.Println("queue_buffer 1")
	case <-time.After(5 * time.Second):
		fmt.Println("queue_buffer time out")
	}

}
