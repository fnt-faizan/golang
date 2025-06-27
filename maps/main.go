package main

import "fmt"

// unordered map
// delete(mapName, key)
func main() {
	students := make(map[string]int)
	students["faizan"] = 80
	students["tapan"] = 10
	students["abcs"] = 20
	for k, v := range students {
		fmt.Println(k, ":", v)
	}
	grade, exists := students["faiza"]
	fmt.Println(grade, " : ", exists)
	tester := make(map[string]string)
	tester["faizan"] = "nabi"
	tester["sahil"] = "jaiswal"
	for k, v := range tester {
		fmt.Println(k, v)
	}
}
