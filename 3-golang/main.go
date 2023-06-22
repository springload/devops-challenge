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
		go func(i int) {
			fmt.Println(i)
			time.Sleep(1 * time.Second)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
