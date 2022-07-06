## golang challenge

As a DevOps you are developing a small tool that makes use of sync.WaitGroup
to make some parallel API calls.

This quite simplified program is expected to print numbers from 0 to 4, but instead
it prints 5 multiple times! What's going on?

## Answers

1. It happend because each iteration of the loop uses the same instance of the variable `i`, so each closure shares that single variable. When the closure runs, it prints the value of `i` at the time `fmt.Println` is executed, but `i` may have been modified since the goroutine was launched. We can solve it thru two ways (first one is preferable):

```Golang
...
go func(j int) {
    fmt.Println(j)
    time.Sleep(1 * time.Second)
    wg.Done()
...
```

```Golang
...
i := i
go func() {
    fmt.Println(i)
    time.Sleep(1 * time.Second)
    wg.Done()
...
}
```

