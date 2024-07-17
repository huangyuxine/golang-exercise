package main

import "fmt"

type Good interface {
	info()
}

type Phone struct {
	Price int
	Name  string
}

func (phone Phone) info() {
	fmt.Printf("name:%s , price:%d", phone.Name, phone.Price)
}

type Water struct {
	Price int
	Name  string
}

func (water Water) info() {
	fmt.Printf("name:%s , price:%d", water.Name, water.Price)
}

func getGoodInfo(good Good) {
	good.info()
}
func main() {
	phone := Phone{
		Price: 5000,
		Name:  "iphone",
	}
	phone.info()

	water := Water{
		Price: 2,
		Name:  "农夫",
	}
	phone.info()

	// 实现接口里的所有方法，就代表实现了接口

	var good Good
	good = Phone{
		Price: 5000,
		Name:  "iphone",
	}
	good.info()

	// 多态
	getGoodInfo(phone)
	getGoodInfo(water)
}
