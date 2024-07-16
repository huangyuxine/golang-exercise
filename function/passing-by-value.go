package main

import "fmt"

func setSlice(slice []int) {
	slice[0] = 10
}

func setMap(map1 map[string]string) {
	map1["name"] = "b"
}

func setInt(number *int) {
	*number = 100
}

func main() {

	// 切片引用类型
	slice := []int{1, 2, 3}
	setSlice(slice)
	fmt.Println(slice) // [10 2 3]

	// map引用类型
	map1 := map[string]string{
		"name": "a",
	}
	setMap(map1)
	fmt.Println(map1) // map[name:b]

	// 传地址
	number := 1
	setInt(&number)
	fmt.Println(number)
}
