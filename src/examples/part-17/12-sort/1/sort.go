package main

import (
	"fmt"
	"sort"
)

func main() {

	// 它是原地排序的，所以他会直接改变给定的切片，而不是返回一个新切片
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)

	//使用 sort 来检查一个切片是否为有序的
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

}
