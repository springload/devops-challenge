## golang challenge

As a DevOps you are developing a small tool that makes use of sync.WaitGroup
to make some parallel API calls.

This quite simplified program is expected to print numbers from 0 to 4, but instead
it prints 5 multiple times! What's going on?

#The WaitGroup type of sync package is used to wait for the program to finish all goroutines launched from the main function. It uses a counter that specifies the number of goroutines, and Wait blocks the execution of the program until the WaitGroup counter is zero. For loop will not start untill the goroutines are executed & the WaitGroup counter is zeor
