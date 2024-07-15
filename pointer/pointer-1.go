package main

import "fmt"

func main() {
	var a = 100

	// 指向a的内存地址
	var b = &a

	*b = 200

	fmt.Println(a)

}
