package main

import "fmt"

func array_test() {
	var array1 [5]string
	array1[0] = "ee"
	array1[1] = "pkg"
	fmt.Println(array1)

	var arr2 = make([]int, 5)
	arr2[0] = 1
	fmt.Println(arr2)
	fmt.Println(len(arr2))
	arr2 = append(arr2, 5)
	fmt.Println(arr2)
	fmt.Println(len(arr2))

	var arr3 = make([]int, 6)
	copy(arr3, arr2)
	fmt.Println(arr3)

	arr4 := arr3[0:2]
	fmt.Println(arr4)

	for i := 0; i < len(arr4); i++ {
		fmt.Println(arr4[i])
	}
}
func map_test() {
	map1 := map[string]string{}
	map1["1"] = "zhangsan"
	map1["2"] = "lisi"
	for key, value := range map1 {
		fmt.Printf("id: %s ,name: %s\n", key, value)
	}
	fmt.Println(len(map1))
	delete(map1, "1")
	fmt.Println(len(map1))

	key1 := "1"
	value, exists := map1[key1]
	fmt.Println(exists)
	if exists {
		fmt.Println(value)
	}

	key2 := "2"
	if key2 == "2" {
		fmt.Println("key == 2")
	}

}
func main() {
	array_test()
	map_test()
}
