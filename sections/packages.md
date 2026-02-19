# Packages in Go

A package groups related Go files that share the same directory and declare the same package name. This is how Go organizes code, controls visibility, and enables reuse across projects. Package names should be short, lowercase, and a single word — avoid camelCase, underscores, or hyphens.

## Naming conventions

- **Packages** use nouns (what it contains): `database`, `http`, `user`
- **Functions** use verbs (what it does): `Connect()`, `Query()`, `Send()`

```go
package database  // noun - what it is

func Query(sql string) {  // verb - what it does
    // implementation
}
```

## Special case: package main

The `main` package is a special package used to create executable programs rather than reusable libraries.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```