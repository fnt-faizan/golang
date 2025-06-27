package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
buff io vs scanln scan
*/
func main() {
	//fmt.scanln(&s)
	fmt.Printf("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	inp, _ := reader.ReadString('\n')
	fmt.Printf("%s\n", inp)

}
