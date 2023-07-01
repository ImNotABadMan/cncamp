package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	messages := make(chan int, 10)
	ctx := context.WithValue(context.Background(), "key", "value")
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	// 协程之间的上下文通信
	fmt.Println("协程之间的上下文通信")
	go func(inCtx context.Context) {
		fmt.Println(inCtx.Value("key"))
	}(ctx)

	// 子协程获取数据
	go func(ctx context.Context) {
		ticker := time.NewTicker(time.Second)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				// 超时就退出子协程
				fmt.Println("child goroutine interrupt", ctx.Err())
				return
			case value := <-messages:
				fmt.Printf("子协程获取到：%d\n", value)
			}
		}
	}(ctx)

	for index := 1; index < 10; index++ {
		messages <- rand.Intn(1000)
	}

	//
	time.Sleep(time.Second * 2)

	// 主协程也监听
	select {
	case <-ctx.Done():
		// 超时两秒，子协程超时上下文捕捉到了
		time.Sleep(time.Second * 2)
		fmt.Println("main done exit", ctx.Err())
	}
}
