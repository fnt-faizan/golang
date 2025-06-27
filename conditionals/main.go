package main

import "fmt"

func main() {
	x := 100
	if x < 100 {
		fmt.Printf("smaller")
	} else if x > 200 { //in same line
		fmt.Printf("greater")
	} else {
		fmt.Printf("Greater than 100 and less than 200")
	}
	if true {
		fmt.Println("Hello")
	} else if !true {
		fmt.Println("second")
	} else {
		fmt.Println("thrird")
	}
}
