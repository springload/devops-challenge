## golang challenge

### Question
As a DevOps you are developing a small tool that makes use of sync.WaitGroup
to make some parallel API calls.

This quite simplified program is expected to print numbers from 0 to 4, but instead
it prints 5 multiple times! What's going on?

### Answer
The purpose of running goroutines is to run processes in parallel. Because of this, it is not expected for operations to return in sequence. 
If order matters, you should be using synchronous processes, rather than parallel processes.

sync.Waitgroup is used to ensure goroutines have finished executing before continuing.

The reason why the program was not printing the correct values was because the variable `i` was actually a single variable and the closure is bound to that variable.

By adding `i` as a parameter to the closure, `i` is evaluated at each iteration and placed on the stack of the goroutine.

Common mistakes for go loop iteratior variables here:
https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables