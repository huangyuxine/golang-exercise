package main

import "fmt"

func main() {
	str := "哈哈\r\n嘿嘿"
	str1 := str + "哼哼"
	str2 := `哈哈哈哈哈哈\r\n`
	fmt.Println(str)
	fmt.Println(str1)
	fmt.Println(str2)
}
