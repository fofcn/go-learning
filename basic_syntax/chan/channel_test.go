package chan_test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChan_ShouldSuccesss_WhenUseUnbuffferedChannelAndOneSendRoutinAnotherRecvRoutine(t *testing.T) {

	// 创建一个无缓冲通道
	ch := make(chan int)

	// 创建一个goroutine
	go func() {

		// 向通道发送1
		// ch <- 1
		log.Println("完成")
	}()

	// 从通道接收
	var val int = <-ch

	// // 断言收到的是1
	assert.Equal(t, 1, val)

}
