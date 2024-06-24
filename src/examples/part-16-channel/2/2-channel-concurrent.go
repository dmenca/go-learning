package main

import (
	"fmt"
	"time"
)

//模拟java的10线程执行命令，等待10个线程结果返回后再继续执行命令

func main() {
	// 创建一个用于通知的通道
	done := make(chan struct{})
	// 创建一个带缓冲的通道，用于传递结果
	var resps = make(chan int, 10)
	// 启动一个 goroutine 来处理结果
	go func() {
		// 接收并打印结果
		for resp := range resps {
			fmt.Printf("resp:%d\n", resp)

		}
		close(done)
	}()

	// 向 resps 通道发送结果
	for i := 0; i < 10; i++ {
		go func() {
			resps <- i
		}()
		time.Sleep(100 * time.Millisecond)
	}

	// 需要执行close操作 因为 resps 通道还没有被关闭，for resp := range resps 循环在接收完所有数据后仍然会等待新的数据，从而导致死锁。
	// 通知接收 goroutine 没有更多数据要发送。 等已有的resps数据被处理完后就会关闭。
	close(resps)
	// 等待接收 goroutine 完成
	<-done
}
