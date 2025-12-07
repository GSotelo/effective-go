# Declarations

## var vs :=

Go provides two ways to declare variables: the `var` keyword and the short variable declaration `:=`.

### var

The `var` keyword can be used at any scope level:

```go
// Package level
var packageVar = "available everywhere"

func example() {
    // Function level
    var localVar = "available in function"

    if true {
        // Block level
        var blockVar = "available in block"
    }
}
```

### Short Variable Declaration (:=)

The `:=` operator declares and initializes a variable with type inference. **It cannot be used outside of functions.**

```go
func example() {
    // Valid: inside a function
    name := "John"
    count := 42

    if count > 0 {
        // Valid: inside a block within a function
        message := "Count is positive"
        fmt.Println(message)
    }
}
```

```go
// Invalid: cannot use := at package level
package main

// name := "John"  // compilation error
var name = "John"  // correct
```

At the package level, use `var` for variable declarations:

```go
var (
    MaxConnections = 100
    Timeout = 30 * time.Second
    DefaultName = "unnamed"
)
```