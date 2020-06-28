package main

import "fmt"

func main() {
	foods := []string{
		"popcorn",
		"pizza",
		"hotdogs",
		"apples",
	}
	for int, value := range foods {
		fmt.Println(int, value)
	}

}
