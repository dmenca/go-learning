package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 500; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 10000; j++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("ops:", ops)
}
