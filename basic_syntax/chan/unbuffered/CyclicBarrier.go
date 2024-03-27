package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// CyclicBarrier 让一组goroutine在到达某个点之后才能继续执行
type CyclicBarrier struct {
	// 总goroutine数量
	participant int
	// 用于等待所有goroutine准备好
	waitGroup sync.WaitGroup
	// 无缓冲channel，用于goroutine间同步
	barrierChan chan struct{}
	running     int32
}

// NewCyclicBarrier 创建一个新的CyclicBarrier
func NewCyclicBarrier(participant int) *CyclicBarrier {
	b := &CyclicBarrier{
		participant: participant,
		barrierChan: make(chan struct{}),
		running:     int32(participant),
	}
	// 设置等待的goroutine数
	b.waitGroup.Add(participant)
	return b
}

// 当一个goroutine调用Wait时，
// 它将在屏障处等待，
// 直到所有goroutine都到达这里
func (b *CyclicBarrier) Wait() {
	// 一个goroutine准备好了
	b.waitGroup.Done()

	// 等待所有goroutine都准备好
	b.waitGroup.Wait()

	// 当所有goroutine都准备好了，关闭channel进行广播通知
	if atomic.AddInt32(&b.running, -1) == 0 {
		close(b.barrierChan)
	} else {
		// 等待通知
		<-b.barrierChan
	}

}

// 阻塞调用goroutine直到所有goroutine都调用了Wait方法，
// 屏障开放后，重新置为待关闭状态
func (b *CyclicBarrier) Await() {
	// 等待屏障开放的信号
	<-b.barrierChan

	// 重置屏障状态
	b.barrierChan = make(chan struct{})
	b.waitGroup.Add(b.participant)
}

func (b *CyclicBarrier) Close() {
	close(b.barrierChan)
}

func main() {
	// 这里我们设置3个goroutine参与
	barrier := NewCyclicBarrier(100)

	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Printf("Goroutine %d is working...\n", i)
			// 模拟工作
			time.Sleep(time.Duration(1) * time.Second)
			fmt.Printf("Goroutine %d reached the barrier.\n", i)
			barrier.Wait()

			fmt.Printf("Goroutine %d passed the barrier.\n", i)
		}(i)
	}

	// 主goroutine等待所有goroutine都到达屏障
	barrier.Await()
	fmt.Println("All goroutines have passed the barrier")
}
