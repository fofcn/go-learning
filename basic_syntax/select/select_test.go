package select_test

import (
	"sync"
	"testing"
	"time"
)

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

func TestSelect_ShouldSuccess_WhenUsingCorrectSyntax(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	chans := []chan int{ch1, ch2, ch3}

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)

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
