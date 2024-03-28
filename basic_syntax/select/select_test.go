package select_test

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

func TestSelect_ShouldRecvZeroValue_WhenSelectFromClosedChannel(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	close(ch1)
	close(ch2)
	close(ch3)

	select {
	case value := <-ch1:
		if value == 0 {
			println("Recieved zero value from ch1")
		} else {
			println("Recieved ch1")
		}
	case value := <-ch2:
		if value == 0 {
			println("Recieved zero value from ch2")
		} else {
			println("Recieved ch2")
		}
	case value := <-ch3:
		if value == 0 {
			println("Recieved zero value from ch3")
		} else {
			println("Recieved ch3")
		}
	case <-time.After(10 * time.Second):
		println("Timeout")
	default:
		println("Default")
	}
}

func TestSelect_ShouldPanic_WhenSendToClosedChannel(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	close(ch1)
	close(ch2)
	close(ch3)

	assert.Panics(t, func() {
		select {
		case ch1 <- 1:
			println("send ch1")
		case ch2 <- 1:
			println("send ch2")
		case ch3 <- 1:
			println("send ch3")
		case <-time.After(10 * time.Second):
			println("Timeout")
		default:
			println("Default")
		}
	})

}
