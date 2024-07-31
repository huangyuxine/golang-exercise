package main

import "fmt"

func getChan(ch chan int) {
	// 获取信道容量
	n := cap(ch)
	for i := 0; i < n; i++ {
		ch <- i
	}
	// 关闭信道
	close(ch)
}
func main() {
	ch := make(chan int, 10)

	go getChan(ch)

	for i := range ch {
		fmt.Print(i)
	}
}
