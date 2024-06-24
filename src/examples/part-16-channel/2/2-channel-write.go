package main

import (
	"fmt"
	"time"
)

//有缓冲的queue，发送操作满了没有读取的话会导致写阻塞，直到有数据被读取后才能继续写入

var q1 = make(chan int, 2)
var val int

func main() {

	go func() {
		time.Sleep(10 * time.Second)
		val = <-q1
		val = <-q1
		fmt.Println("receive val:", val)
	}()

	// 通过channel 可以
	for i := 0; i < 3; i++ {
		select {
		// 写入两次到q1之后，第三次写入q1就会阻塞，因为q1的size是2两个，而读取的操作需要在10s后进行，没有读取就会导致写阻塞
		case q1 <- 1:
			fmt.Println("send val:", i)
		case <-time.After(2 * time.Second):
			fmt.Println("time out")
		}
	}

}
