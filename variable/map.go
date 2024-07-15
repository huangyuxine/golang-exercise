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

}
