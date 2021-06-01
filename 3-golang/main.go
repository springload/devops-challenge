package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup

func main() {
	//var wg sync.WaitGroup
	wg.Add(5) //Counter starts from 5 which controls to wait until myPrint finishes.
	go myPrint() //Execute the gorutine 
	/*for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}*/
	
	wg.Wait()//Block the execution of main() function until the goroutines in the waitgroup have successfully completed. (Execute next line until wg = 0.)

	fmt.Println("Program finished.")	
}

func myPrint() {
	for i := 0; i < 5; i++ { //This for loop is to control the printing from 0 to 4. 
		fmt.Println(i) //print the i value in the for loop.
		wg.Done() //Counter-1 in one run of the for loop. Control the main program not to continue the next line until the counter equals to 0       
    	}
}

/* 
In order to print 0 to 4, it requires 2 variables to:
1. Print the numbers, starting from 0 and stopping at 4.
2. The WaitGroup variables blocks the main function until the printing is done.
3. The WaitGroup counter should be aligned with the loop range to avoid deadlock.

Problems about the previous version of the code:
1. The goroutine 'func()' does not receive any variables as a parameter 'i' to control the printing. So, the for loop doesn't really affect the value for the printing line. 
2. The goroutine 'func()' can only get the value 'i' from the memory/cache which is the last value from the for loop range instead of the starting value. 
3. While the wg variables allows to wait until counter equals to 0, thereby allowing five times execution of the for loop before continuing the main function. 
*/  

