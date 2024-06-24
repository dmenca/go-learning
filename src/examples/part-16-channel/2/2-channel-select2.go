package main

import (
	"fmt"
	"time"
)

//queue是一个无缓冲（unbuffered）的channel。在Go语言中，无缓冲的channel要求发送和接收必须同时准备好，才能进行数据传递。也就是说，发送操作会阻塞，直到有接收者来读取数据

var q1 = make(chan int)
var val int

func main() {
	go func() {
		time.Sleep(3 * time.Second)
		q1 <- 1
		fmt.Println("send q1 1")
	}()

	go func() {
		time.Sleep(2 * time.Second)
		q1 <- 3
		fmt.Println("send q1 3")

	}()
	go func() {
		for {
			s := "22"
			if s == "33" {
				break
			}
		}

	}()
	for {
		select {
		// 在这里，queue <- 1操作会阻塞，因为没有任何接收者读取这个值。由于没有其他接收者，5秒后time.After case会被触发，输出“time out”。
		case val = <-q1:
			fmt.Println("val:", val)
			// case <-time.After(5 * time.Second):
			// 	fmt.Println("time out")
		}
	}

}
