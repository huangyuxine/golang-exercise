package main

import (
	"fmt"
	"sync"
)

// 全局变量
var wg sync.WaitGroup

func foo1() {
	for i := 1; i < 20; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar1() {
	for i := 20; i < 40; i++ {
		fmt.Println("bar:", i)
	}
	wg.Done() // 计数器 - 1
}

// 一主多子的协程协作
func main() {
	// concurrency.go
	wg.Add(2) // 携程数量计数器
	go foo1()
	go bar1()
	wg.Wait() // 阻塞当前协程，直到计数器为0停止
}
