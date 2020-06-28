package main

import "fmt"

func main() {
	var operator string
	var number1, number2 int

	fmt.Print("Please enter the first number: ")
	fmt.Scanln(&number1)
	fmt.Print("Please enter Second number: ")
	fmt.Scanln(&number2)
	fmt.Print("Please enter the operator you would like: ")
	fmt.Scanln(&operator)

	output := 0

	if operator == "+" {
		output = number1 + number2
	}
	if operator == "-" {
		output = number1 - number2
	}
	if operator == "*" {
		output = number1 * number2
	}
	if operator == "/" {
		output = number1 / number2
	}

	fmt.Printf("%d %s %d = %d", number1, operator, number2, output)
}
