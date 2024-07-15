package main

import "fmt"

func main() {
	// 本质和int32差不多
	var num1 rune = 100000
	// 表示unicode字符 可省略类型
	var s = '中'
	var num2 int32 = '中'

	fmt.Println(num1)
	fmt.Println(s)    // 20013
	fmt.Println(num2) // 20013
}
