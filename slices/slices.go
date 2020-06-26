package main

import "fmt"

func main() {
	foods := []string{
		"popcorn",
		"pizza",
		"hot dogs",
		"apples",
	}

	fmt.Println(foods)
	foods = append(foods, "burrito")
	fmt.Println(foods)
}
