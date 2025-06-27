package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error occured")
	} else {
		data, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error parsing")
		} else {
			fmt.Println(string(data))
		}
	}
}
