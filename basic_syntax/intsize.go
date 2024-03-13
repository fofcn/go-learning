package main

import "unsafe"

func main() {
	var a int8 = 0
	println("sizeof(int): %u", unsafe.Sizeof(a))
}
