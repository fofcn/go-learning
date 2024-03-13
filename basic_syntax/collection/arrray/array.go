package main

import "fmt"

func main() {
	var array [5]int

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	array = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	value := array[2]
	println("value: %d", value)

	// 查询数组长度
	println("value: %d", len(array))

	// 遍历for循环
	for i := 0; i < len(array); i++ {
		value = array[i]
	}

	// 遍历range
	for i, value := range array {
		println("index: %d, value: %d", i, value)
	}
}
