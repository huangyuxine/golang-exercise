package main

import "fmt"

func intSum(numbers []int) int {
	sums := 0
	for _, number := range numbers {
		sums += number
	}
	return sums
}

func float64Sum(numbers []float64) float64 {
	sums := 0.0
	for _, number := range numbers {
		sums += number
	}
	return sums
}

// 使用泛型
func sum[T int | float64](numbers []T) T {
	var sums T
	for _, number := range numbers {
		sums += number
	}
	return sums
}

func main() {
	num1 := []int{1, 2, 3, 4}
	num2 := []float64{1, 2, 3, 4}
	fmt.Println(intSum(num1))
	fmt.Println(float64Sum(num2))

	fmt.Println(sum(num1))
	fmt.Println(sum[int](num1))
	fmt.Println(sum(num2))
	fmt.Println(sum[float64](num2))

}
