package main

import "fmt"

// 只读
func get(ch <-chan int) {
	fmt.Println(<-ch)
}

// 只写
func put(ch chan<- int, input int) {
	ch <- input
}

func main() {
	rwCh := make(chan int, 1)
	put(rwCh, 1)
	get(rwCh)

}
