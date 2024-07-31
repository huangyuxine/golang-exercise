package main

import (
	"fmt"
	"sync"
)

func add(count *int, wg *sync.WaitGroup, lock *sync.Mutex) {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		*count = *count + 1
		lock.Unlock()
	}
	wg.Done()
}

// 互斥锁是为了来保护一个资源不会因为并发操作而引起冲突导致数据不准确
// 此案例不加互斥锁会导致每次的结果不等于3000
func main() {
	var wg3 sync.WaitGroup
	lock := sync.Mutex{}
	num := 0
	wg3.Add(3)

	go add(&num, &wg3, &lock)
	go add(&num, &wg3, &lock)
	go add(&num, &wg3, &lock)

	wg3.Wait()

	fmt.Println("count 的值为：", num)
}
