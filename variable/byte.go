package main

import "fmt"

func main() {
	// byte可以表示数字 0-255，和uint8一样的效果
	var num1 byte = 255
	// 表示ASCII字符
	var num2 byte = 'a'
	fmt.Println(num1)
	fmt.Println(num2) // 97
}
