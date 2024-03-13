package main

import "fmt"

func modifyValue(val int) {
	// 修改的是副本
	val *= 2
}

func main() {
	initial := 10
	// 发送一个原始值的副本
	modifyValue(initial)
	// initial保持不变
	fmt.Println("initial after modifyValue:", initial)
}
