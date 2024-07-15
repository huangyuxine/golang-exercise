package main

import "fmt"

func main() {
	// 二进制
	num1 := 0b1100
	// 八进制
	num2 := 0o14
	// 十六进制
	num3 := 0x114
	// 十进制
	num4 := 999
	fmt.Printf("2进制数 %b 表示的是: %d \n", num1, num1)
	fmt.Printf("8进制数 %o 表示的是: %d \n", num2, num2)
	fmt.Printf("16进制数 %X 表示的是: %d \n", num3, num3)
	fmt.Println(num4)
}
