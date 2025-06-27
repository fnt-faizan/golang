package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {
	//myurl := "https://www.instagram.com/fntfaizan?viewId=5"
	//create  a areader
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n') //to pause the program
	s = strings.TrimSpace(s)        //to remove the trailing newline character
	fmt.Printf("You entered: %s\n", s)
	parsedurl, error := url.Parse(s)
	fmt.Printf("Type of parsedurl: %T\n", parsedurl)
	if error != nil {
		fmt.Println("Error ocxcured!")
	} else {
		fmt.Println(parsedurl)
		fmt.Println(parsedurl.Host, parsedurl.Path, parsedurl.RawQuery, parsedurl.Port())
	}
}
