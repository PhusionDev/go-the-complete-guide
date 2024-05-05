package main

import "fmt"

func main() {
	// exercise()

	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99
	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println("All Prices w/Discounts:", prices)
}

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	fmt.Println(prices)

// 	fmt.Println(prices[0])
// 	fmt.Println(productNames)

// 	featuredPrices := prices[1:3]
//   featuredPrices2 := prices[1:]
// 	fmt.Println(featuredPrices)
//   fmt.Println(featuredPrices2)
// }

func exercise() {
	// 1
	hobbies := []string{"Programming", "Video Games", "Audio Production"}
	fmt.Println(hobbies)

	// 2
	fmt.Println(hobbies[0])
	combined := []string{hobbies[1], hobbies[2]}
	fmt.Println(combined)

	// 3
	sliced1 := hobbies[:2]
	// sliced2 := hobbies[0:2]

	// 4
	resliced := sliced1[1:]
	fmt.Println(resliced)

	// 5
	courseGoals := []string{"Get a job", "Make more money"}
	fmt.Println(courseGoals)

	// 6
	courseGoals[1] = "Buy a sandwich"
	courseGoals = append(courseGoals, "Level Up")
	fmt.Println(courseGoals)

	// 7
	type Product struct {
		title string
		id    int
		price float64
	}

	products := []Product{
		{
			"My First Product",
			1,
			13.98,
		},
		{
			"My Second Product",
			2,
			2.98,
		},
	}
	newProduct := Product{"My Third Product", 3, 8.98}
	products = append(products, newProduct)
	fmt.Println(products)

}
