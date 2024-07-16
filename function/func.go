package main

import (
	"fmt"
)

// 无参
func greet() {
	fmt.Println("greet")
}

// 一个参数
func greet1(name string) {
	fmt.Println(name)
}

// 两个参数
func greet2(name string, age int) {
	fmt.Printf("name:%s; age:%d", name, age)
}

// 两个类型一致的参数
func greet3(name, address string) {
	fmt.Printf("name:%s; address:%s", name, address)
}

// 一个返回值
func greet4(name string) string {
	return name
}

// 两个返回值
func greet5(name string, age int) (string, int) {
	return name, age
}

// 返回一个有名的值
func greet6(name string) (address string) {
	address = name + " address"
	return address
}

// 返回两个个有名的值
func greet7(name1 string) (name2, address string) {
	address = name1 + " address"
	return name1, address
}

// 语法糖
func sum(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

// 语法糖
func avg(numbers ...int) float64 {
	sums := 0
	length := len(numbers)
	for _, number := range numbers {
		sums += number
	}
	result := float64(sums / length)
	return result
}

func main() {
	result := sum(1, 2)
	fmt.Println(result)

	result1 := avg([]int{1, 2}...)
	fmt.Println(result1)

	name, address := greet7("h")
	fmt.Println(name, address)
}
