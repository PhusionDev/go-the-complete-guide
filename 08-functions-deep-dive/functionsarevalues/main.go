package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Printf("Numbers: %v\nDoubled: %v\n", numbers, doubled)
	fmt.Printf("Numbers: %v\nTripled: %v\n", numbers, tripled)

	transformFn1 := getTransformer(&numbers)
	transformFn2 := getTransformer(&moreNumbers)

	transformed1 := transformNumbers(&numbers, transformFn1)
	transformed2 := transformNumbers(&moreNumbers, transformFn2)

	fmt.Println(transformed1)
	fmt.Println(transformed2)
}

// function that returns a function
func getTransformer(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
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
