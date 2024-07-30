package main

import "fmt"

func foo() {
	for i := 1; i < 20; i++ {
		fmt.Println("foo:", i)
	}
}

func bar() {
	for i := 20; i < 40; i++ {
		fmt.Println("bar:", i)
	}
}

// main 的地位相当于主线程，当 main 函数执行完成后，这俩线程也就终结了
// 协程还没来得及并执行
func main() {
	go foo()
	go bar()
}
