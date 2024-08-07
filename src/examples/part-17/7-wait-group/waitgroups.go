package main

import (
	"fmt"
	"sync"
	"time"
)

// 每个协程都会运行该函数。
// 注意，WaitGroup 必须通过指针传递给函数。
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", id)

}

func main() {
	// WaitGroup 用于等待该函数启动的所有协程。
	var wg sync.WaitGroup
	// 启动几个协程，并为其递增 WaitGroup 的计数器。
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	// 阻塞，直到 WaitGroup 计数器恢复为 0；
	// 即所有协程的工作都已经完成。
	wg.Wait()
}
