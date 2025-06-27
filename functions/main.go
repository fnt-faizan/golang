package main

import "fmt"

func simple() {
	fmt.Println("Hello this is a function")
}
func add(a int, b int) int { //curly braces must start on the same line
	return a + b

}

// go mod init myprojecr
func tester(a *int, b *int) int {
	(*a)++
	return *a + *b
}
func main() {
	res := add(5, 6)
	fmt.Println(res)
	a := 10
	b := 12
	fmt.Println(tester(&a, &b))
	simple()
}
