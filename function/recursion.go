package main

import "fmt"

func Fibonacci(number int) int {
	if number < 2 {
		return number
	}
	return Fibonacci(number-1) + Fibonacci(number-2)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(Fibonacci(i))
	}
}
