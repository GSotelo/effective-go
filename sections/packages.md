# Packages in Go

Using hyphens in Go module namescan i use - when i create module name in go8:26 PMYes, but with a caveat.
In go.mod (module path): Hyphens are allowed.
module my-module-name
In import paths: Also fine — hyphens are valid URL characters, so packages hosted on GitHub etc. with hyphens work normally.
The catch — package directory names with hyphens: If a directory is named my-utils, Go will use it as the import path segment, but the package name inside must be a valid Go identifier (no hyphens). So you'd write:
go// directory: my-utils/
package myutils  // or utils, myUtils, etc. — no hyphen allowed
Then import it as:
goimport "my-module/my-utils"
And reference it in code using the package name you declared (myutils.Foo()), not the directory name.
Short answer: Hyphens are fine in module and directory names, but not in the package declaration itself.is it normal to use camel case for package name?8:31 PMNo, it's not idiomatic in Go. The convention is lowercase, single word — no camelCase, no underscores, no hyphens.
gopackage myutils   // ✅ idiomatic
package myUtils   // ❌ not idiomatic
package my_utils  // ❌ not idiomatic
The Go team explicitly recommends short, lowercase, single-word names. If you can't avoid combining words, just run them together lowercase (myutils, httputil, strconv). You can see this throughout the standard library: fmt, bufio, filepath, httputil, etc.paraphrase better: A package is a collection of related code files in the same directory. All files in a package declare the same package name at the top. Go uses packages to organize code, control visibility, and enable code reuse. Don't use hyphens in package names.The convention is lowercase, single word — no camelCase, no underscores, no hyphens.8:32 PMA package groups related Go files that share the same directory and declare the same package name. This is how Go organizes code, controls visibility, and enables reuse across projects. Package names should be short, lowercase, and a single word — avoid camelCase, underscores, or hyphens.

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