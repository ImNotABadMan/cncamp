package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan int, 10)
	done := make(chan bool)
	go func() {
		t := time.NewTicker(time.Second)
		for _ = range t.C {
			select {
			case value, ok := <-done:
				// 接受退出信号
				fmt.Println("child goroutine interept...", value, ok)
				return
			default:
				fmt.Printf("get %d\n", <-messages)
			}
		}
	}()
	for index := 0; index < 10; index++ {
		messages <- index
	}
	time.Sleep(time.Second * 5)
	// 关闭通道，由发送方关闭，告知接受方，最后关闭的状态，所有接收方都不阻塞了，拿到空值，和ok = false
	close(done)
	time.Sleep(time.Second)
	fmt.Println("main exit")
}
