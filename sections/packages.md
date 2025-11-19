# Packages in Go

A package is a collection of related code files in the same directory. All files in a package declare the same package name at the top. Go uses packages to organize code, control visibility, and enable code reuse.

## Naming conventions

**Package names should be nouns** — they describe what the package contains, not what it does.

**Function names should be verbs** — they describe the action being performed.

```go
package database  // noun - what it is

func Connect(dsn string) {  // verb - what it does
    // implementation
}

func Query(sql string) {  // verb - what it does
    // implementation
}
```