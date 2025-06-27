package main

import "fmt"

type person struct {
	age    int
	name   string
	height float32
}
type bird struct {
	canfly    bool
	wingspan  float32
	migratory bool
}

func main() {
	var faizan person = person{age: 24, name: "faizan", height: 5.10}
	var person1 = new(person)
	fmt.Println(faizan)
	fmt.Println(person1)
	bulbul := bird{canfly: true, wingspan: 0.102, migratory: false}
	fmt.Println(bulbul)
	// for i := 0; i < 10; i++ {
	// 	for j := 0; j < 10-i; j++ {
	// 		fmt.Printf(" ")
	// 	}
	// 	for k := 0; k < i; k++ {
	// 		fmt.Printf("*")
	// 	}
	// 	fmt.Println("")
	// }
}
