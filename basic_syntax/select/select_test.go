package select_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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
