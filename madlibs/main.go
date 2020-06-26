package main

import "fmt"

func main() {
	var exclamation string = ""
	var adverb string = ""
	var noun string = ""
	var adjective string = ""

	fmt.Print("Enter a exclamation! ")
	fmt.Scanln(&exclamation)
	fmt.Print("Enter a adverb! ")
	fmt.Scanln(&adverb)
	fmt.Print("Enter a noun! ")
	fmt.Scanln(&noun)
	fmt.Print("Enter a Adjective")
	fmt.Scanln(&adjective)

	fmt.Println(exclamation + "! He said " + adverb + " as he jumped into his " + adverb + " covertible " + "and drove off with his " + adjective + "dog.")
}
