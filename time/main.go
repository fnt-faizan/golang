package main

import (
	"fmt"
	"time"
)

// time.parse
func main() {
	currtime := time.Now()
	fmt.Println(currtime)
	fmt.Printf("Type of time is %T\n", time.Now())
	ctime := "2025-06-20"
	atime, _ := time.Parse("2006-01-02", ctime)
	atime = time.Now()
	fmt.Println(atime)
}
