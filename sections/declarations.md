# Declarations

## var vs :=

Go provides two ways to declare variables: the `var` keyword and the short variable declaration `:=`.

### Short variable declaration (:=)

The `:=` operator declares and initializes a variable with type inference. It can only be used inside functions and blocks, never at package level. The `var` keyword is used for package-level declarations and when explicit type specification is needed.

```go
package main

import (
    "fmt"
    "time"
)

// Package level: must use var
var (
    MaxConnections = 100
    Timeout = 30 * time.Second
    DefaultName = "unnamed"
)

// name := "John"  // compilation error: := not allowed at package level

func example() {
    // Valid: := inside a function with type inference
    name := "John"
    count := 42

    if count > 0 {
        // Valid: := inside a block within a function
        message := "Count is positive"
        fmt.Println(message)
    }

    // Can also use var inside functions if explicit typing is needed
    var port int = 8080
    fmt.Printf("Server running on port %d\n", port)
}
```