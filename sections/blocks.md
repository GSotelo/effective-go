# Blocks in Go

A block is a set of statements enclosed by matching braces `{}`. Blocks create scopes — regions where variables exist and can be accessed. Understanding blocks is essential to understanding how variables interact in Go.

## File block
Import declarations are file-scoped — each file must import the packages it needs.

```go
// file1.go
package myapp

import "fmt"  // fmt is only accessible in file1.go

func PrintMessage(msg string) {
    fmt.Println(msg)  // can use fmt here
}
```

```go
// file2.go
package myapp

import "fmt"  // must import separately; file1.go's import doesn't apply here

func AnotherFunction() {
    fmt.Println("Hello")  // can use fmt here because of the import above
}
```

## Package block
Variables declared at the package level (outside functions) are accessible from any file in that package. First, let's create a library package with exported and unexported variables:

```go
// config/settings.go
package config

import "fmt"

var defaultTimeout = 30  // lowercase - unexported, only accessible within config package

// Timeout is the request timeout in seconds
var Timeout = defaultTimeout  // uppercase - exported, accessible from other packages

func LogSettings() {
    fmt.Println("Timeout:", defaultTimeout)  // can use package-level variable
}
```

Now, another file in the same `config` package can access both:

```go
// config/server.go
package config

import "fmt"

func InitServer() {
    fmt.Println("Using timeout:", defaultTimeout)  // can access unexported defaultTimeout - same package
    fmt.Println("Server ready with timeout:", Timeout)
}
```

From a different package (main), you can only access exported variables:

```go
// main.go
package main

import (
    "fmt"
    "config"
)

func main() {
    fmt.Println("App timeout:", config.Timeout)  // can access exported Timeout from config package
    // fmt.Println(config.defaultTimeout)  // ERROR - unexported, not accessible across packages
}
```

**Note:** Package-level variables should be avoided because they introduce global mutable state, making code harder to test, reason about, and vulnerable to unexpected mutations across your package.

## Function block
Scope: Inside a function body. Variables declared here are local to that function only and cannot be accessed from outside.

```go
package main

import "fmt"

func main() {
    message := "Hello"  // local to main() - not accessible elsewhere
    fmt.Println("From main:", message)

    greet()
}

func greet() {
    // message is not accessible here - different function scope
    greeting := "Hi there"  // local to greet() function
    fmt.Println("From greet:", greeting)
}
```

## Control flow blocks
Scope: Inside control structures. Variables declared here exist only within that specific block.

**if/else block:**
```go
package main

import "fmt"

func main() {
    age := 25

    if age >= 18 {
        status := "adult"  // only exists inside this if block
        fmt.Println("Status:", status)
    }
    // status is not accessible here
    fmt.Println("Age check complete")
}
```

**for loop block:**
```go
package main

import "fmt"

func main() {
    for i := 1; i <= 3; i++ {
        result := i * 2  // only exists inside this loop iteration
        fmt.Println("i =", i, "result =", result)
    }
    // result is not accessible here - loop scope ended
    fmt.Println("Loop finished")
}
```

**switch/case block:**
```go
package main

import "fmt"

func main() {
    day := 1

    switch day {
    case 1:
        dayName := "Monday"  // only exists in this case block
        fmt.Println("Day:", dayName)
    case 2:
        dayName := "Tuesday"  // different block, can reuse the name
        fmt.Println("Day:", dayName)
    default:
        fmt.Println("Unknown day")
    }
    // dayName is not accessible here
    fmt.Println("Switch complete")
}
```

## Explicit blocks
Scope: Standalone `{ ... }` blocks create a new inner scope without any condition. Useful for limiting variable lifetime.

```go
package main

import "fmt"

func main() {
    x := 100  // outer scope

    {
        x := 50  // inner explicit block - shadows outer x
        fmt.Println("Inside block:", x)  // prints 50
    }

    fmt.Println("Outside block:", x)  // prints 100 - inner x no longer exists
}
```

## Key concepts

### Scope and visibility
- **Inner blocks can access outer blocks**: Variables declared in outer blocks are visible and accessible in inner blocks.
- **Outer blocks cannot access inner blocks**: Once you leave an inner block, variables declared there are no longer accessible.

### Variable lifetime
- Variables exist only from their declaration until the end of their block.
- Once a block closes, all variables declared in that block are destroyed.
- This helps manage memory and prevents accidental variable reuse.

### Block hierarchy
Blocks follow a clear hierarchy from largest to smallest scope:
1. **Package block** - Widest scope, shared across all files in a package (for declarations: variables, constants, functions, types)
2. **File block** - Specific to one `.go` file (imports are file-scoped)
3. **Function block** - Local to a specific function
4. **Control flow blocks** - Limited to if/for/switch structures
5. **Explicit blocks** - Smallest scope, created with `{ }`

Key distinction: Package-level declarations are shared across all files in the package, but imports are per-file (file-scoped).

Variables declared in inner blocks override those in outer blocks within that scope.

### Shadowing risk
You can declare a variable with the same name in an inner block. This **shadows** (hides) the outer variable within that inner scope. While technically valid, shadowing can cause confusion and bugs. Be careful when reusing variable names in nested scopes.