package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Queue struct {
	message []string
	cond    *sync.Cond
}

func main() {
	fmt.Println("课题2.1，多个生产者和多个消费者")
	queue := Queue{
		message: []string{},
		cond:    sync.NewCond(&sync.Mutex{}),
	}

	done, canel := context.WithTimeout(context.Background(), time.Second*10)
	defer canel()

	for i := 0; i < 2; i++ {
		go producer(&queue, i)
	}

	for i := 0; i < 2; i++ {
		go consumer(&queue, i)
	}

	for {
		select {
		case <-done.Done():
			// 10秒超时就结束
			fmt.Println("main exit")
			return
		default:
		}
	}
}

func producer(queue *Queue, index int) {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		queue.cond.L.Lock()
		str := fmt.Sprintf("test_%d", rand.Int())
		queue.message = append(queue.message, str)
		queue.cond.Broadcast()
		fmt.Printf("producer[%d]: %s\n", index, str)
		queue.cond.L.Unlock()
	}

}

func consumer(queue *Queue, index int) {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		queue.cond.L.Lock()
		if len(queue.message) == 0 {
			queue.cond.Wait()
		}
		message := queue.message[0]
		queue.message = queue.message[1:]
		fmt.Printf("consumer[%d] - %s\n", index, message)
		queue.cond.L.Unlock()
	}
}
