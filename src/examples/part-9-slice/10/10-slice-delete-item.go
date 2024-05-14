package main

import "fmt"

func main() {
	nums := make([]int, 5)
	nums[0] = 0
	nums[1] = 1
	nums[2] = 2
	nums[3] = 3
	nums[4] = 4

	// 删除切片中第三个元素 需要使用append 切片替代删除
	delete_index := 2
	if delete_index < len(nums) {
		// 使用...来变成...Type类型
		nums = append(nums[:delete_index], nums[delete_index+1:]...)
		fmt.Println("Slice after removing element,", nums)
	} else {
		fmt.Println("Index out of range")
	}
}
