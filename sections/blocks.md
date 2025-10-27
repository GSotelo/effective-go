# Blocks in Go

A block is a set of statements enclosed by matching braces `{}`. Blocks create scopes—regions where variables exist and can be accessed. Understanding blocks is essential to understanding how variables interact in Go.

## Basic Block Types in Go

### 1. File Block
Scope: The file-level declarations within a single `.go` file. Variables declared here are visible throughout the file.

```go
package main

// File block: visible throughout this file
const maxSize = 100
var counter = 0

func init() {
    counter = 1  // can access file-level variables
}
```

### 2. Package Block
Scope: All files in the same package. Variables declared at the package level (outside functions) are accessible from any file in that package.

**Within Same Package (main):**
```go
// file1.go
package main

var config = "settings"  // lowercase - unexported, accessible within package main

func Setup() {
    fmt.Println(config)  // can use package-level variable
}
```

```go
// file2.go
package main

func Start() {
    fmt.Println(config)  // can access config from file1.go - same package
}
```

**Across Different Packages:**
To share variables across packages, define in a non-main package with uppercase (exported):
```go
// mylib/config.go
package mylib

var Config = "shared"  // uppercase = exported, accessible from other packages
```

```go
// main.go
package main

import "mylib"

func main() {
    fmt.Println(mylib.Config)  // access exported variable from mylib package
}
```

### 3. Function Block
Scope: Inside a function body. Variables declared here are local to that function only and cannot be accessed from outside.

```go
func main() {
    message := "Hello"  // local to main() - not accessible elsewhere
    fmt.Println(message)
}

func other() {
    // message is not accessible here - different function scope
}
```

### 4. Control Flow Blocks
Scope: Inside control structures. Variables declared here exist only within that specific block.

**if/else block:**
```go
age := 25

if age >= 18 {
    status := "adult"  // only exists inside this if block
    fmt.Println(status)
}
// status is not accessible here
```

**for loop block:**
```go
for i := 1; i <= 3; i++ {
    result := i * 2  // only exists inside this loop iteration
    fmt.Println(result)
}
// result is not accessible here - loop scope ended
```

**switch/case block:**
```go
day := 1

switch day {
case 1:
    dayName := "Monday"  // only exists in this case block
    fmt.Println(dayName)
case 2:
    dayName := "Tuesday"  // different block, can reuse the name
    fmt.Println(dayName)
}
// dayName is not accessible here
```

### 5. Explicit Blocks
Scope: Standalone `{ ... }` blocks create a new inner scope without any condition. Useful for limiting variable lifetime.

```go
x := 100  // outer scope

{
    x := 50  // inner explicit block - shadows outer x
    fmt.Println("Inside block:", x)  // prints 50
}

fmt.Println("Outside block:", x)  // prints 100 - inner x no longer exists
```

## Key Concepts

1. **Nested Scopes**: Inner blocks can see and use variables from outer blocks, but outer blocks cannot see variables declared in inner blocks.

2. **Variable Lifetime**: A variable exists only from its declaration to the end of its block. Once a block closes, its variables are gone.

3. **Hierarchy**: Blocks follow a hierarchy:
   - Package ← File ← Function ← Control Flow ← Explicit Blocks
   - Variables are accessible from inner to outer, but not the other way around.

4. **Shadowing Risk**: You can declare a variable with the same name in an inner block, which hides (shadows) the outer variable.