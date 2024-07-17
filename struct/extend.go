package main

import "fmt"

type People struct {
	Name string
	Age  int
}

func (people People) Eat() {
	fmt.Println("eat")
}

type Xiaoming struct {
	People
	Class  string
	Gender string
}

// 重写
func (xiaoming Xiaoming) Eat(food string) {
	fmt.Println("eat" + food)
}

func main() {
	// 实例化
	xm := Xiaoming{
		People: People{Name: "h"},
		Class:  "三班",
		Gender: "男",
	}
	xm.Eat("meat")
	// 使用new
	xm1 := new(Xiaoming)
	xm1.People.Age = 1
	xm1.People.Name = "b"
	xm1.Class = "四班"
	xm1.Gender = "女"
	xm1.Eat("shit")
}
