package main

import (
	"fmt"
	"sync"
)

func worker(id int, w *sync.WaitGroup) {
	defer w.Done()
	println("Worker", id, "started")
}
func main() {
	var wg sync.WaitGroup
	for i := 1; i < 5; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go worker(i, &wg)

	}
	wg.Wait() // Wait for all workers to finish
	fmt.Println("All workers ended")
}
