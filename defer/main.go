package main

import "fmt"

func main() {
	defer fmt.Println("deferred statement call")
	//not executed until main is about to exit
	fmt.Println("not deferred")
}
