package main

import "fmt"

func main() {
	// 3.7E+1 表示 37
	// 3.7E-2 表示 0.037

	// 有效位7位, 8位开始损失精度
	var num1 float32 = 100000291
	fmt.Println(num1) // 1.0000029e+08

	// 有效位15位, 16位就开始损失精度
	var num2 float64 = 1000000000000092
	fmt.Println(num2) // 1.000000000000092e+15
}
