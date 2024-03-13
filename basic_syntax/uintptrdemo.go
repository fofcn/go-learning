package main

import (
	"fmt"
	"unsafe"
)

// func main() {
	var x []int = make([]int, 2)
	x[0] = 42
	x[1] = 43

	// 获取第一个元素:42的地址
	firstElementPtr := unsafe.Pointer(&x[0])

	// 计算第二个元素的地址
	// 确保我们根据 int 类型的大小来移动指针
	secondElementPtr := unsafe.Pointer(uintptr(firstElementPtr) + unsafe.Sizeof(x[0]))

	// 通过 secondElementPtr 获取第二个元素的值
	secondElementValue := *(*int)(secondElementPtr)

	fmt.Printf("The value of the second element is: %d\n", secondElementValue)
}
