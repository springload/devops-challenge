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

