package main

import "fmt"

func main() {
	age := 20

	if age < 15 {
		fmt.Println("少年")
	} else if age > 15 && age < 40 {
		fmt.Println("青年")
	} else if age > 40 && age < 60 {
		fmt.Println("中年")
	} else {
		fmt.Println("老年")
	}

	if age := 20; age > 18 {
		fmt.Println("已经成年了")
	}
}
