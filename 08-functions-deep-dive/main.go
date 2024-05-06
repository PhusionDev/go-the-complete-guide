package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Printf("Numbers: %v\nDoubled: %v\n", numbers, doubled)
	fmt.Printf("Numbers: %v\nTripled: %v\n", numbers, tripled)
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

// transformFn is type alias for `func(int) int` in this case
func transformNumbers(numbers *[]int, transform transformFn) []int {
	transformed := []int{}

	for _, v := range *numbers {
		transformed = append(transformed, transform(v))
	}

	return transformed
}
