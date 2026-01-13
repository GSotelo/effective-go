# Blocks in Go

A block is a set of statements enclosed by matching braces `{}`. Blocks create scopes — regions where variables exist and can be accessed. Understanding blocks is essential to understanding how variables interact in Go.

## Universal block

The universal block contains Go's **predeclared** identifiers — built-in constants, types, and functions that are available everywhere in your program without needing to import anything. These are the foundational building blocks of the language.

**Predeclared constants:**
```go
true   // boolean constant
false  // boolean constant
iota   // auto-incrementing constant (used in const blocks)
nil    // zero value for pointers, slices, maps, channels, functions, and interfaces
```

**Predeclared types:**
```go
// Basic types
bool
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64, uintptr
float32, float64
complex64, complex128
string
byte  // alias for uint8
rune  // alias for int32

// Aggregate types (built-in)
error  // interface for error handling
```

**Predeclared functions:**
```go
make()    // create slices, maps, channels
new()     // allocate memory
len()     // length of array, slice, string, map
cap()     // capacity of array, slice
append()  // add elements to slice
copy()    // copy slice elements
delete()  // remove map element
clear()   // clear all elements from map or slice
max()     // return maximum value from arguments
min()     // return minimum value from arguments
complex() // create complex number
real()    // get real part of complex
imag()    // get imaginary part of complex
close()   // close a channel
print()   // print to stderr (for bootstrapping/debugging)
println() // print to stderr with newline (for bootstrapping/debugging)
panic()   // trigger panic
recover() // recover from panic
```

Here's a practical example showing some universal block identifiers:

```go
package main

import "fmt"

func main() {
    // Using predeclared constants
    isValid := true
    emptyValue := nil

    // Using predeclared types
    var age int = 30
    var name string = "Alice"
    var ok bool = true

    // Using predeclared functions
    slice := make([]int, 0, 10)  // create a slice with capacity 10
    slice = append(slice, 1, 2, 3)  // add elements
    sliceLength := len(slice)  // get length

    fmt.Println(isValid, age, name, ok, sliceLength)
}
```

**Key point:** You don't need to import anything to use predeclared identifiers. They're automatically available in every Go program, making them the most universally accessible scope in Go.

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

import "fmt"  // imports are file-scoped; each file needs its own

func AnotherFunction() {
    fmt.Println("Hello")  // can use fmt here because of the import above
}
```

## Function block
Inside a function body. Variables declared here are local to that function only and cannot be accessed from outside.

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
Inside control structures. Variables declared here exist only within that specific block.

**if/else:**
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

**for loop:**
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

**switch:**
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
- Variables declared in outer blocks are visible and accessible in inner blocks.
- Once you leave an inner block, variables declared there are no longer accessible.

### Variable lifetime
- Variables exist only from their declaration until the end of their block.
- Once a block closes, all variables declared in that block are destroyed.
- This helps manage memory and prevents accidental variable reuse.

### Block hierarchy
Blocks follow a clear hierarchy from largest to smallest scope:
- **Universal block** - Widest scope, predeclared identifiers available everywhere in your program
- **Package block** - Shared across all files in a package (for declarations: variables, constants, functions, types)
- **File block** - Specific to one `.go` file (imports are file-scoped)
- **Function block** - Local to a specific function
- **Control flow blocks** - Limited to if/for/switch structures
- **Explicit blocks** - Smallest scope, created with `{ }`

Variables declared in inner blocks override those in outer blocks within that scope.

### Shadowing risk
You can declare a variable with the same name in an inner block. This **shadows** (hides) the outer variable within that inner scope. While technically valid, shadowing can cause confusion and bugs. Be careful when reusing variable names in nested scopes.