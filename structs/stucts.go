package main

import "fmt"

func main() {
	type contactInfo struct {
		name    string
		age     int
		email   string
		address string
		zipcode int
		phone   string
	}

	james := contactInfo{
		name:    "James",
		age:     31,
		email:   "jrperkins@protonmail.com",
		address: "North Carolina",
		zipcode: 27529,
		phone:   "5555555555",
	}
	fmt.Println(james)
}
