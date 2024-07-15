package main

import "fmt"

func IsAdult(age int) bool {
	if age >= 16 {
		return true
	}
	return false
}
func main() {

	age := 28
	result := IsAdult(age)

	fmt.Println(result)
}
