package main

import "fmt"

func root(x int) int {
	i := 0
	j := x
	for i < j {
		mid := i + (j-i)/2
		if mid*mid < x {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return i
}
func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	x := []int{239, 329, 92, 902, 01, 30}
	for i, v := range x {
		fmt.Println(i, ":", v)
	}
	tester := make([]int, 50, 50)
	for i, v := range tester {
		fmt.Printf("Index %d , Value %d\n", i, v)
	}
	for i := 2; i < 100; i++ {
		var x = true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				x = false
				continue
			}
		}
		if !x {
			continue
		}
		fmt.Println(i, "is a prime number")
	}
	fmt.Println(root(49))
}
