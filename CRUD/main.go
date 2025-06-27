package main

import (
	"fmt"
	"io"
	"net/http"
)

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	UserId    int    `json:"userId"`
}

func main() {
	req, err := http.Get("https://jsonplaceholder.typicode.com/todos/2")
	if err != nil {
		return
	}
	// var t Todo
	// fmt.Println()
	// dat := json.NewDecoder(req.Body).Decode(&t)
	// if dat != nil {
	// 	return
	// }
	// fmt.Println(t)
	defer req.Body.Close() //close the body to prevent resource leak
	//consumed once used
	// fmt.Println(req)
	// fmt.Println(req.Body)
	// data, _ := io.ReadAll(req.Body)
	// fmt.Println(data)
	// fmt.Println(string(data))
	// var t Todo
	// _ = json.Unmarshal(data, &t)
	// fmt.Println(t)
	// var t Todo
	// dat := json.NewDecoder(req.Body).Decode(&t)
	// if dat != nil {
	// 	fmt.Println("Error decoding JSON:", dat)
	// }
	// fmt.Println(t)
	dat, _ := io.ReadAll(req.Body) //read the body
	fmt.Println(string(dat))       //convert to string and print
}
