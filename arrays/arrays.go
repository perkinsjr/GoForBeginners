package main

import "fmt"

func main() {
	var intArray [5]int

	intArray[0] = 1
	intArray[1] = 2
	intArray[2] = 3
	intArray[3] = 4
	intArray[4] = 5
	fmt.Println(intArray)
	fmt.Println(intArray[3])
}
