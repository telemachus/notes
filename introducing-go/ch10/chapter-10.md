# Introducing Go, Chapter 10: Concurrency

Go supports concurrency, which Doxsey defines as “[m]aking progress on more than one task simultaneously,” via goroutines and channels.

## Goroutines

Goroutines are invoked as `go <function call>`. Normally, all the statements in a function are executed before the next line is run, but in the case of a goroutine, execution of the next lines does not wait. According to Doxsey, goroutines are “lightweight and we can easily create thousands of them.”

## Channels

Channels allow goroutines to communicate and to synchronize execution. You can pass messages to and from channels using the `<-` operator.

```go
// You can send to a channel.
c <- "pong"

// You can receive from a channel.
fmt.Println(<-c)

// You can also assign from a channel.
msg := <-c
```

### Channel Direction

You can specify a direction for a channel type in order to restrict it either to sending or to receiving.

```go
// You can only send to c; you cannot receive from it.
// If you try to receive from c, the program will not compile.
func pinger(c chan<- string)

// You can only receive from c; you cannot send to it.
// If you try to send to c, the program will not compile.
func printer(c <-chan string)
```

If you can send and receive from a channel, then the channel is said to be *bidirectional*. You can send a bidirectional channel to a function that accepts send-only or receive-only channels. But you cannot send a one-way channel to a function that expects a bidirectional channel.

### Select

Go provides a `switch`-like statement, but for channels: `select`.

```go
c1 := make(chan string)
c2 := make(chan string)

go func() {
    for {
        c1 <- "from 1"
        time.Sleep(time.Second *2)
    }
}()

go func() {
    for {
        c2 <- "from 2"
        time.Sleep(time.Second * 3)
    }
}()

go func() {
    for {
        select {
        case msg1 := <-c1:
            fmt.Println(msg1)
        case msg2 := <-c2:
            fmt.Println(msg2)
        }
    }
}()
```

### Buffered Channels

When you create a channel with `make`, you can add a second value for the capacity of the channel. E.g., `c := make(chan int, 1)`. This creates a buffered channel with a capacity of 1. A buffered channel will not send or wait until the channel is full. (I don’t understand what this means.)
