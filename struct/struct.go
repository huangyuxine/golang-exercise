package main

import "fmt"

// 结构体
type Person struct {
	name string
	age  int
}

// 方法
func (person *Person) SetName(name string) {
	person.name = name
}

func main() {
	person := Person{
		name: "a",
		age:  18,
	}
	fmt.Println(person.name)

	person.SetName("b")
	fmt.Println(person.name)
}
