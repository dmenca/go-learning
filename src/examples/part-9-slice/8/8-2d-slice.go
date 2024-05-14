package main

import "fmt"

func main() {
	// Slice 可以组成多维数据结构。内部的 slice 长度可以不一致，这一点和多维数组不同。
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
