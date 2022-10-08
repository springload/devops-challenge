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
		i := i //This is based on the go documentation here -> https://go.dev/doc/faq#closures_and_goroutines and here -> https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		//the above solution does print out values between 0-4 but not in order
		go func() {
			fmt.Println(i)
			time.Sleep(1 * time.Second)
			wg.Done()
		}()
		/*fmt.Println(i)
		if you MOVE the print statement here then the values will print in order
		however that defeats the purpose of running the go routine concurrently*/
	}
	wg.Wait()
}

//Running go vet gives the result "loop variable i captured by func literal", googled that statement to figure out the solution above
