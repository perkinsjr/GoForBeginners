package main

import "fmt"

func main() {
	number := 1

	for number < 5 {
		number *= 2
		fmt.Println(number)
	}
}
