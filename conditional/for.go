package main

import "fmt"

func main() {

	// 常用
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 无条件
	number := 100
	for {
		if number < 90 {
			break
		}
		number--
	}
	// 一个条件
	a := 10
	for a > 0 {
		fmt.Println(a)
		a--
	}

	map1 := map[string]string{
		"name":  "h",
		"class": "six",
	}
	for key, val := range map1 {
		fmt.Printf("key:%s, val:%s", key, val)
	}

	slice1 := []string{"a", "b", "c"}
	for _, val := range slice1 {
		fmt.Println(val)
	}
}
