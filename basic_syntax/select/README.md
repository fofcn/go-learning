# 概述

这里又有多路复用，但是Go中的这个多路复用不同于网络中的多路复用。在Go里，select用于同时等待多个通信操作（即多个channel的发送或接收操作）。Go中的channel可以参考我的文章：[逐步学习Go-并发通道chan(channel)](https://fofcn.tech/2024/03/27/%e9%80%90%e6%ad%a5%e5%ad%a6%e4%b9%a0go-%e5%b9%b6%e5%8f%91%e9%80%9a%e9%81%93chanchannel/)

## 拆字解释

- 多路：指的是多个channel操作路径。你可以在select块中定义多个case，每个case对应一个channel上的I/O操作（发送或接收）。


- 复用：指的是select的功能，它可以监听多个channel上的事件，并且仅当其中一个channel准备就绪时才会执行相关操作。这样，单个goroutine可以高效地等待多个并发事件而不是单个事件。

**复用的是goroutine，一个goroutine使用select可以监听多个信道。**


整体来讲：Select就是为channel设计的。

![这张图是参考大佬Dravenss画的](https://fofcn.tech:443/wp-content/uploads/2024/03/image-1711546607411.png)

# select语法
Go语言中的select关键字功能在概念上与操作系统的select类似，区别在于Go的select是用于goroutine监听多个channel的可读或可写状态。

Go的select允许在channel上进行非阻塞收发，同时当多个channel同时响应时，select会随机执行其中的一个case。

Go的select语句可以包含一个default分支，使得在没有channel准备好时，不会阻塞goroutine，而是执行default分支。

```go

	select {
	case <-ch:
		println("recieved")
	case <-time.After(10 * time.Second):
		println("Timeout")
	default:
		printStr = "Hello Select"
	}


```

接下来我们来看场景用例。

# select只有一个case条件满足

```go

func TestSelect_ShouldRecvChan1_WhenChan1CaseWasFullfilled(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	chans := []chan int{ch1, ch2, ch3}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		chans[0] <- 1
		wg.Done()
	}()

	wg.Wait()

	select {
	case <-ch1:
		println("Recieved ch1")
	case <-ch2:
		println("Recieved ch2")
	case <-ch3:
		println("Recieved ch3")
	case <-time.After(10 * time.Second):
		println("Timeout")
	default:
	}

}

```


# select有多个case条件满足

```go

func TestSelect_ShouldRandomEnterCaseBranch_WhenAllChannelsCaseWereFullfilled(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	chans := []chan int{ch1, ch2, ch3}

	var wg sync.WaitGroup
	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func(i int) {
			chans[i] <- 1
			wg.Done()
		}(i)
	}

	wg.Wait()

	select {
	case <-ch1:
		println("Recieved ch1")
	case <-ch2:
		println("Recieved ch2")
	case <-ch3:
		println("Recieved ch3")
	case <-time.After(10 * time.Second):
		println("Timeout")
	default:
	}

}

```

# select没有条件满足-阻塞

```go

func TestSelect_ShouldBlock_WhenNoCaseWasFullfilled(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	select {
	case <-ch1:
		println("Recieved ch1")
	case <-ch2:
		println("Recieved ch2")
	case <-ch3:
		println("Recieved ch3")
	}

}

```


# select没有条件满足-超时

```go

func TestSelect_ShouldTimeout_WhenNoCaseWasFullfilled(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	select {
	case <-ch1:
		println("Recieved ch1")
	case <-ch2:
		println("Recieved ch2")
	case <-ch3:
		println("Recieved ch3")
	case <-time.After(10 * time.Second):
		println("Timeout")
	}

}

```

# select没有条件满足-default

```go

func TestSelect_ShouldRunDefaultBranch_WhenNoCaseWasFullfilledAndHasDefaultBranch(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	select {
	case <-ch1:
		println("Recieved ch1")
	case <-ch2:
		println("Recieved ch2")
	case <-ch3:
		println("Recieved ch3")
	case <-time.After(10 * time.Second):
		println("Timeout")
	default:
		println("Default")
	}

}


```
