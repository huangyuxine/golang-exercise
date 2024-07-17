package main

import "fmt"

func getmap(map1 map[string]string) {
	map1["name"] = "h"
}

func main() {
	// map引用类型, 无序
	map1 := map[string]string{
		"name": "a",
	}
	getmap(map1)
	fmt.Println(map1) // map[name:h]

	// 分配初始化容量
	var map2 map[string]string
	map2 = make(map[string]string, 2)
	map2["name"] = "a"
	map2["age"] = "1"
	map2["age1"] = "1"

	// 空map等于nil
	var map3 map[string]string
	fmt.Println(map3 == nil)

	// 新增
	map4 := map[string]int{}
	map4["water"] = 2
	map4["iphone"] = 5000
	fmt.Println("map4:", map4)
	// 删
	delete(map4, "water")
	fmt.Println("map4:", map4)
	// 改
	map4["iphone"] = 1
	fmt.Println("map4:", map4)
	// 查
	fmt.Println("iphone:", map4["iphone"])

	// 长度
	fmt.Println(len(map4))

	// 判断值是否存在
	res, ok := map4["iphone"]
	fmt.Println(res, ok)
}
