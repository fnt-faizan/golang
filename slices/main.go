package main

import "fmt"

// slices double size when overfliow
func main() {
	arr := []int{3489, 294892, 928, 29, 91, 91}
	fmt.Println(arr)
	fmt.Printf("Types or arr is : %T\n", arr)
	arr = arr[1:3]
	nsl := make([]int, 1, 2) //using make function to create a slice
	newslice := make([]int, 2, 5)
	tester := [4]byte{'a', 'b', 'c', 'd'}
	fmt.Printf("Type of tester is %T\n", tester)
	fmt.Println(newslice)
	fmt.Println(nsl)
	fmt.Println("Capacity is: ", cap(arr), "Length is: ", len(arr))
}
