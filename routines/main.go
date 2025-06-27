package main

import (
	"fmt"
	"learning/errors"
	"time"
)

func f1() {
	fmt.Println("Function 1 called")
	fmt.Println("Function 1 ended")
}
func f2() {
	fmt.Println("Function 2 called")
	fmt.Println("Function 2 ended")
}
func main() {
	go f1()
	go f2()

	time.Sleep(1 * time.Second)
	fmt.Printf("type of time is %T\n", time.Second)
	fmt.Println(errors.Div(12, 2))
}
