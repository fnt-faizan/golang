package main

import "fmt"

func modify(ref *int) {
	(*ref)++
}
func main() {
	var num int = 5
	var ptr *int = &num
	fmt.Println(*ptr, "stored at", ptr)
	var ptr2 **int = &ptr
	if ptr2 != nil {
		fmt.Println(**ptr2)
	}
	modify(ptr)
	fmt.Println(num)
}
