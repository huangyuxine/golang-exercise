package main

import "fmt"

func main() {
	// 一维数组
	arr := [3]int{1, 2, 3}

	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3

	var arr2 [3]int = [3]int{1, 2, 3}
	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(arr2)

	// 二维数组
	arr3 := [3][4]string{
		{"a", "b", "c", "d"},
		{"a", "b", "c", "d"},
		{"a", "b", "c", "d"},
	}
	fmt.Println(arr3)

}
