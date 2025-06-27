package main

import (
	"fmt"
	"strings"
)

// Split
// count
// trimspace
// join
func main() {
	data := "apple, orange, banana"
	splitter := strings.Split(data, ",")
	fmt.Println(splitter)
	str := "one two threee four two two"
	cust := strings.Count(str, "two")
	fmt.Println(str, cust)
	joined := strings.Join([]string{data, str}, ",")
	fmt.Println(joined)
}
