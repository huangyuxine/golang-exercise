package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	var event []int
	for _, number := range numbers {
		if callback(number) {
			event = append(event, number)
		}
	}
	return event
}

func main() {
	event := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(event) // [2 4]
}
