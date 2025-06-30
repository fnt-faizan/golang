package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// This program fetches a TODO item from a public API and prints it to the console.
func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error fetching data:", err)
	}
	defer res.Body.Close()
	// if res.StatusCode != http.StatusOK {
	// 	fmt.Println(res.Status)
	// } else if res.StatusCode == http.StatusOK {
	// 	data, err := io.ReadAll(res.Body)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	} else {
	// 		fmt.Println(string(data))
	// 	}
	// }
	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	} else {
		fmt.Println(todo)
		fmt.Printf("User ID: %d, ID: %d, Title: %s, Completed: %t\n", todo.UserID, todo.ID, todo.Title, todo.Completed)
	}
}
