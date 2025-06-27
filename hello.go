package main

import (
	"fmt"
	"learning/errors"
)

func main() {
	fmt.Println("Hello World!")
	var faizan string = "my name is faizan"
	var tester = " this is tester"
	fmt.Println(faizan + tester)
	var checked = false
	fmt.Println("Value of boolean is: ", checked)
	person := 727.92389 //shortcut
	fmt.Println(person)
	var csis string = "hello this is for practice"
	fmt.Println(csis)
	fmt.Println(errors.Div(12, 2)) //imported from error package
	//Capitalization - want to export then start with capital varibles or functions
}
