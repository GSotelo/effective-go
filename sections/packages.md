# Packages in Go

A package is a collection of related code files in the same directory. All files in a package declare the same package name at the top. Go uses packages to organize code, control visibility, and enable code reuse. Don't use hyphens in package names.

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