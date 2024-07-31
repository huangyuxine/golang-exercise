package main

import (
	"fmt"
	"sync"
)

var data map[string]string
var rwMutex sync.RWMutex

func init() {
	data = make(map[string]string)
	data["key"] = "value"
}

func readData(key string, wg *sync.WaitGroup) {
	defer wg.Done() // 确保在函数退出前调用Done()

	rwMutex.RLock()
	value, exists := data[key]
	if exists {
		fmt.Printf("Read: %s -> %s\n", key, value)
	}
	rwMutex.RUnlock()
}

func writeData(key, value string, wg *sync.WaitGroup) {
	defer wg.Done() // 确保在函数退出前调用Done()

	rwMutex.Lock()
	data[key] = value
	fmt.Printf("Write: %s -> %s\n", key, value)
	rwMutex.Unlock()
}

func main() {
	var wg5 sync.WaitGroup
	wg5.Add(4) // 初始化WaitGroup的计数器为5

	go readData("key", &wg5)
	go readData("key", &wg5)
	go writeData("key", "new_value", &wg5)
	go readData("key", &wg5)

	wg5.Wait() // 等待所有goroutines完成
}
