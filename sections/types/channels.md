# Channels in Go

Channels are Go's primary mechanism for communication between goroutines. They allow you to pass data safely between concurrent operations, embodying Go's philosophy: "Don't communicate by sharing memory; share memory by communicating."

This philosophy means that instead of having multiple goroutines access shared variables (which requires complex locking), you pass data through channels. One goroutine "owns" the data at a time, eliminating race conditions naturally.

Channels are reference types with a zero value of `nil`. Like slices and maps, they must be created using `make` before use.

## Creating Channels

Channels are created using the `make` function with the `chan` keyword:

```go
// Create an unbuffered channel for integers
ch := make(chan int)

// Create a buffered channel with capacity 5
bufferedCh := make(chan int, 5)

// Create a channel for custom types
type Message struct {
    Text string
    ID   int
}
msgCh := make(chan Message)
```

## Channel Operations

Channels support three main operations: send, receive, and close.

```go
ch := make(chan int)

// Send a value to the channel
ch <- 42

// Receive a value from the channel
value := <-ch

// Receive with ok idiom to check if channel is closed
value, ok := <-ch
if !ok {
    fmt.Println("Channel is closed")
}

// Close a channel
close(ch)
```

## Channel Types

### Unbuffered Channels
Unbuffered channels provide synchronous communication - the sender blocks until a receiver is ready.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan string)
    
    go func() {
        time.Sleep(2 * time.Second)
        ch <- "Hello from goroutine!"
    }()
    
    fmt.Println("Waiting for message...")
    msg := <-ch
    fmt.Println(msg)
}
```

### Buffered Channels
Buffered channels allow asynchronous communication up to their capacity.

```go
package main

import "fmt"

func main() {
    ch := make(chan int, 3)
    
    // These sends won't block because buffer has capacity
    ch <- 1
    ch <- 2
    ch <- 3
    
    // Receive the values
    fmt.Println(<-ch) // 1
    fmt.Println(<-ch) // 2
    fmt.Println(<-ch) // 3
}
```

## Directional Channels

You can restrict channels to be send-only or receive-only:

```go
// Send-only channel
func sender(ch chan<- int) {
    ch <- 42
}

// Receive-only channel
func receiver(ch <-chan int) {
    value := <-ch
    fmt.Println(value)
}

func main() {
    ch := make(chan int)
    
    go sender(ch)
    receiver(ch)
}
```

## Channel Patterns

### Fan-out Pattern
Distribute work among multiple goroutines:

```go
func fanOut(input <-chan int, workers int) []<-chan int {
    outputs := make([]<-chan int, workers)
    
    for i := 0; i < workers; i++ {
        output := make(chan int)
        outputs[i] = output
        
        go func(out chan<- int) {
            defer close(out)
            for n := range input {
                out <- n * n // Process the data
            }
        }(output)
    }
    
    return outputs
}
```

### Select Statement
Handle multiple channel operations:

```go
select {
case msg1 := <-ch1:
    fmt.Println("Received from ch1:", msg1)
case msg2 := <-ch2:
    fmt.Println("Received from ch2:", msg2)
case <-time.After(1 * time.Second):
    fmt.Println("Timeout")
default:
    fmt.Println("No channels ready")
}
```

## Go Channels: What Can Go Wrong

Here are the key pitfalls to watch out for when working with channels in Go:

### Deadlocks
The most common issue is creating deadlocks by blocking on channel operations:

```go
package main

func main() {
    ch := make(chan int)
    
    // This will deadlock - no goroutine to receive
    ch <- 42
    
    // This line will never be reached
    fmt.Println("Done")
}
```

**Solution**: Always ensure there's a corresponding receiver for every sender:

```go
package main

import "fmt"

func main() {
    ch := make(chan int)
    
    go func() {
        ch <- 42
    }()
    
    value := <-ch
    fmt.Println("Received:", value)
}
```

### Sending on Closed Channels
Sending to a closed channel causes a panic:

```go
package main

func main() {
    ch := make(chan int)
    close(ch)
    
    // This will panic
    ch <- 42
}
```

**Solution**: Use the comma ok idiom and coordinate channel closing:

```go
package main

import "fmt"

func main() {
    ch := make(chan int, 1)
    
    // Safe way to send with checking
    select {
    case ch <- 42:
        fmt.Println("Sent successfully")
    default:
        fmt.Println("Channel full or closed")
    }
    
    close(ch)
}
```

### Goroutine Leaks
Forgetting to close channels or leaving goroutines blocked can cause memory leaks:

```go
// Bad: goroutine may leak if context is cancelled
func worker(ctx context.Context, ch chan int) {
    for {
        select {
        case work := <-ch:
            // Process work
            fmt.Println("Processing:", work)
        case <-ctx.Done():
            return // Good: respect context cancellation
        }
    }
}
```

### Range Over Channels
Remember that ranging over a channel requires the channel to be closed:

```go
package main

import "fmt"

func main() {
    ch := make(chan int, 3)
    
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch) // Important: close the channel
    
    for value := range ch {
        fmt.Println(value)
    }
}
```

## Best Practices

1. **Close channels from the sender side** to avoid panic
2. **Use buffered channels carefully** - they can hide timing issues
3. **Always handle the closed channel case** when receiving
4. **Use context for cancellation** rather than closing channels
5. **Prefer channel direction restrictions** for API clarity
6. **Test concurrent code thoroughly** with race detection: `go test -race`