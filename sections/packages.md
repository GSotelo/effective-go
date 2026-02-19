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

## Special case: Package main

The `main` package is special — it tells Go to build an executable rather than a reusable library. Within it, the runtime looks for `func main()` as the program's entry point. It stays lowercase by design since it's not meant to be exported or called by other packages; it's a reserved convention used internally by the runtime.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```