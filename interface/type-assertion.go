package main

import "fmt"

// any = interface{}

// 类型断言，判断类型

func Validate(params any) {
	//isString, ok := params.(string)
	switch params.(type) {
	case string:
		fmt.Printf("string类型，值为：%s", params)
	case int:
		fmt.Printf("int类型，值为：%d", params)
	case bool:
		fmt.Printf("bool类型，值为：%t", params)
	default:
		fmt.Println("no match")
	}

}

func main() {
	Validate("name")
}
