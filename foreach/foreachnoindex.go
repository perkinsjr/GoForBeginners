package main

import "fmt"

func main() {
	foods := []string{
		"popcorn",
		"pizza",
		"hotdogs",
		"apples",
	}
	for index, value := range foods {
		fmt.Println(index, value)
	}

}
