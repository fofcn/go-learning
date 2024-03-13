package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	close(ch)
	// time.Sleep(10 * time.Second)
	ch = make(chan int)
	go func() {
		// 发送数据到通道
		for i := 0; i < 20; i++ {
			fmt.Printf("sleep one seconds: %d\n", i+1)
			time.Sleep(time.Second)
		}
		fmt.Println("prepare to send to channel")
		ch <- 42
		fmt.Println("send to channel")

	}()
	go func() {
		// 从通道接收数据
		select {
		case value := <-ch:
			fmt.Printf("case <- ch: %d\n", value) // 输出: 42
			break
		case <-time.After(30 * time.Second):
			println("case <- ch timeout") // 输出: 42
			break
		}
	}()

	// value := <-ch
	// println(value) // 输出: 42

	// // 从通道接收数据
	// value = <-ch
	// println(value) // 输出: 42

	time.Sleep(10000 * time.Second)
}
