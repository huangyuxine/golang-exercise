package main

import (
	"fmt"
	"sync"
)

// 用来传输数据，多个goroutine之间传递
func main() {
	// 申明方式，并设置缓冲容量，如果不设置会报错fatal error: all goroutines are asleep - deadlock!
	ch := make(chan int, 1)
	// 往信道中发送数据
	ch <- 200
	// 往信道中接收数据
	data := <-ch
	// 关闭信道

	fmt.Printf("信道可缓冲 %d 个数据\n", cap(ch))
	fmt.Printf("信道有 %d 个数据\n", len(ch))
	close(ch)
	fmt.Printf("信道数据 %d", data)

	// 双向信道
	var wgs sync.WaitGroup

	sender := make(chan int)

	wgs.Add(2)
	go func() {
		fmt.Println("准备发送数据: 100")
		sender <- 100
		wgs.Done()
	}()

	go func() {
		receiver := <-sender
		fmt.Printf("接收到的数据是: %d", receiver)
		wgs.Done()
	}()
	// 主函数wait，使得上面两个goroutine有机会执行
	wgs.Wait()
}
