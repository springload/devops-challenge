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
		go func() {
			time.Sleep(1 * time.Second)
			wg.Done()
		}()
		// wait for all goroutines to complete before printing output
		fmt.Println(i)
	}
	wg.Wait()
}
