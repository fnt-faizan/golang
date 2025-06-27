package main

import "fmt"

// fallthrough statenment
func main() {
	x := 5
	switch x {
	case 1:
		fmt.Println("mon")
	case 2:
		fmt.Println("tue")
	case 3:
		fmt.Println("wed")
	default:
		fmt.Println("sat")

	}
}
