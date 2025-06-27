package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// This is a placeholder main function.
	// You can add your code here or import other packages to use their functionality.
	// For example, you can import "fmt" and print something to the console.
	type bird struct {
		Canfly    bool    `json:"canfly"`
		Wingspan  float32 `json:"wingspan"`
		Migratory bool    `json:"migratory"`
		Color     string  `json:"color"`
	}
	var bulbul bird = bird{Canfly: true, Wingspan: 0.102, Migratory: false, Color: "green"}
	a, b := json.Marshal(bulbul)
	if b != nil {

	} else {
		println(string(a))
		var bulbul2 bird
		err := json.Unmarshal(a, &bulbul2)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
		}
		fmt.Println(bulbul2)
	}

}
