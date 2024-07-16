package main

import "fmt"

// 例2
func greet9(number int) func() int {
	return func() int {
		return number + 1
	}
}

// 例3
func greet10(number1 int) func(number int) int {
	return func(number int) int {
		return number + number1
	}
}

func main() {
	// 例1
	greet8 := func() {
		fmt.Println("Hello world!")
	}
	greet8()

	fmt.Println(greet9(1)())
	fmt.Println(greet10(2)(3))

}
