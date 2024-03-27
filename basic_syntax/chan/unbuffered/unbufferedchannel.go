package main

import (
	"log"
)

func main() {
	// 创建一个无缓冲通道
	ch := make(chan int)
	defer close(ch)

	// 创建一个goroutine
	go func() {

		// 向通道发送1
		// ch <- 1
		log.Println("完成")
	}()

	// 从通道接收
	<-ch

}
