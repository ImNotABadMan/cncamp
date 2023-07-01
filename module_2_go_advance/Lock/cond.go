package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	messages []string
	cond     sync.Cond
}

func main() {
	queue := Queue{
		messages: []string{},
		cond:     sync.Cond{L: &sync.Mutex{}},
	}
	fmt.Println("条件锁的使用Cond, 本质是锁，只是带条件的锁，什么时候等待，处理边界问题")

	go func() {
		for {
			producer(&queue)
			time.Sleep(time.Second)
		}
	}()

	for {
		consumer(&queue)
		time.Sleep(time.Second)
	}
}

func producer(queue *Queue) {
	queue.cond.L.Lock()
	defer queue.cond.L.Unlock()
	queue.messages = append(queue.messages, "sss")
	fmt.Println("producer", queue.messages)
	// 有数据就通知消费者，条件唤醒
	queue.cond.Broadcast()
	// 唤醒一个
	//queue.cond.Signal()
}

func consumer(queue *Queue) {
	queue.cond.L.Lock()
	defer queue.cond.L.Unlock()
	// 没有数据就等待，条件等待
	if len(queue.messages) == 0 {
		queue.cond.Wait()
	}
	fmt.Println("consumer", queue.messages)
	queue.messages = queue.messages[1:]
	queue.cond.Signal()
}
