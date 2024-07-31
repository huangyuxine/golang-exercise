package main

import (
	"fmt"
	"sync"
)

type Receiver = <-chan int // 只读

type Sender = chan<- int // 只写

func main() {
	// 单向信道
	ch := make(chan int)
	var wg3 sync.WaitGroup
	wg3.Add(2)
	go func() {
		var sender Sender = ch
		fmt.Println("准备发送数据: 100")
		sender <- 100
		wg3.Done()
	}()

	go func() {
		var receiver Receiver = ch
		num := <-receiver
		fmt.Printf("接收到的数据是: %d", num)
		wg3.Done()
	}()
	wg3.Wait()
}
