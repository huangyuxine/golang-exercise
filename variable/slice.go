package main

import "fmt"

func main() {
	// 切片是引用类型的
	slice1 := []int{1, 2, 3}
	var slice2 []int = []int{1, 2, 3}
	var slice3 []int
	// 类型（Type），长度（size），容量（cap）
	slice3 = make([]int, 10, 10)
	slice3[0] = 1
	slice3[1] = 1
	slice3[9] = 1

	fmt.Println(slice1)
	fmt.Println(slice2)

	fmt.Printf("长度: %d, 容量: %d\n", len(slice3), cap(slice3)) // 长度: 10, 容量: 10
	slice3 = append(slice3, 1)

	fmt.Printf("长度: %d, 容量: %d\n", len(slice3), cap(slice3)) // 长度: 11, 容量: 20

	// 二维切片
	slice4 := [][]string{
		{"a", "b"},
		{"c", "d"},
	}

	fmt.Println(slice4)
}
