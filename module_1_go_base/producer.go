package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("课后习题1.2")
	fmt.Println("基于 Channel 编写一个简单的单线程生产者消费者模型：\n\n队列：\n队列长度 10，队列元素类型为 int\n生产者：\n每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞\n消费者：\n每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞")

	ch := make(chan int, 10)
	go func(ch chan int) {
		ticker := time.NewTicker(time.Second)
		for _ = range ticker.C {
			input := rand.Intn(10000)
			ch <- input
			fmt.Printf("生产者接发送：%d\n", input)
		}
	}(ch)

	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		fmt.Printf("消费者接收到：%d\n", <-ch)
	}
}
