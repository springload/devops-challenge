## golang challenge

As a DevOps you are developing a small tool that makes use of sync.WaitGroup
to make some parallel API calls.

This quite simplified program is expected to print numbers from 0 to 4, but instead
it prints 5 multiple times! What's going on?

In order to print 0 to 4, it requires 2 variables to:
1. Print the numbers, starting from 0 and stopping at 4.
2. The WaitGroup variables blocks the main function until the printing is done.
3. The WaitGroup counter should be aligned with the loop range to avoid deadlock.

Problems about the previous version of the code:
1. The goroutine 'func()' does not receive any variables as a parameter 'i' to control the printing. So, the for loop doesn't really affect the value for the printing line. 
2. The goroutine 'func()' can only get the value 'i' from the memory/cache which is the last value from the for loop range instead of the starting value. 
3. While the wg variables allows to wait until counter equals to 0, thereby allowing five times execution of the for loop before continuing the main function. 

