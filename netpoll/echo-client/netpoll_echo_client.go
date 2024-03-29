package main

import (
	"fmt"
	"time"

	"github.com/cloudwego/netpoll"
)

func main() {
	network, address, timeout := "tcp", "192.168.2.10:8080", 50*time.Millisecond

	// use default
	conn, _ := netpoll.DialConnection(network, address, timeout)
	conn.Close()

	// use dialer
	dialer := netpoll.NewDialer()
	conn, _ = dialer.DialConnection(network, address, timeout)

	conn.AddCloseCallback(func(connection netpoll.Connection) error {
		fmt.Printf("[%v] connection closed\n", connection.RemoteAddr())
		return nil
	})

	// write & send message
	writer := conn.Writer()
	message := "hello world"
	writer.WriteString(message)
	writer.Flush()

	reader := conn.Reader()
	defer reader.Release()
	echoMsg, _ := reader.ReadString(len(message))
	fmt.Printf("[recv msg] %v\n", echoMsg)
}
