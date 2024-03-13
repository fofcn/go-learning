package main

import "fmt"

func modifyPointer(ptr *int) {
	// 直接修改原始值
	*ptr *= 2
}

func main() {
	initial := 10
	// 发送内存地址
	modifyPointer(&initial)
	// initial被修改了
	fmt.Println("initial after modifyPointer:", initial)
}
