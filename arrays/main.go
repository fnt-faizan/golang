package main

import "fmt"

func main() {
	var arr [10](int) = [10](int){1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(arr)
	fmt.Println(len(arr))
	var faizan [5]int = [5](int){289, 2938, 283, 2938, 29}
	//default value is zero
	var arr2 [10](int)
	fmt.Println(arr2)
	fmt.Println(faizan)
	faizan[0] = 28
	fmt.Println(faizan)
}
