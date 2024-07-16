package main

import "fmt"

func main() {
	// 例1
	number := 100
	switch number {
	case 0:
		fmt.Println(0)
	case 1, 2:
		fmt.Println(1, "/", 2)
	case 100:
		fmt.Println(100)
	default:
		fmt.Println("No match")
	}
	// 例2
	score := 30

	switch {
	case score >= 95 && score <= 100:
		fmt.Println("优秀")
	case score >= 80:
		fmt.Println("良好")
	case score >= 60:
		fmt.Println("合格")
	case score >= 0:
		fmt.Println("不合格")
	default:
		fmt.Println("有误")
	}
}
