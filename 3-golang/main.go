package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		// Avoid re-use of var - ref: https://go.dev/doc/faq#closures_and_goroutines
		go func(t int) {
			fmt.Println("Start thread: ", t)
			// Pretend workload here
			time.Sleep(1 * time.Second)
			fmt.Println("End thread: ", t)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
