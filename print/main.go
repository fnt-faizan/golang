package main

import "fmt"

// printf vs println
func main() {
	age := 24
	name := "faizan"
	height := 5.10

	fmt.Printf("Name is: %s Age is: %d Height is: %.2f \n", name, age, height)
}
