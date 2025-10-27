# Blocks in Go

A block is a set of statements enclosed by matching braces `{}`. Blocks create scopes—regions where variables exist and can be accessed. Understanding blocks is essential to understanding how variables interact in Go.

## Basic block types in Go

### 1. File block
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

### 2. Package block
Scope: All files in the same package. Variables declared at the package level (outside functions) are accessible from any file in that package.

**Within same package (main):**
```go
// database.go
package main

import "fmt"

var databaseURL = "postgresql://localhost:5432/myapp"  // lowercase - unexported, only accessible within package main

func ConnectDB() {
    fmt.Println("Connecting to:", databaseURL)  // can use package-level variable
}
```

```go
// server.go
package main

import "fmt"

func StartServer() {
    fmt.Println("Database URL:", databaseURL)  // can access databaseURL from database.go - same package
}
```

**Across different packages:**
To share variables across packages, define them in a non-main package with uppercase (exported):

```go
// mylib/settings.go
package mylib

import "fmt"

var privateVar = "only-for-mylib"  // lowercase - unexported, only accessible within mylib package
var PublicVar = "accessible-everywhere"  // uppercase - exported, accessible from other packages

func Setup() {
    fmt.Println(privateVar)  // can use package-level variable
}
```

```go
// main.go
package main

import (
    "fmt"
    "mylib"
)

func main() {
    fmt.Println(mylib.PublicVar)  // can access exported PublicVar from mylib
    // fmt.Println(mylib.privateVar)  // ERROR - privateVar is unexported, not accessible
}
```

### 3. Function block
Scope: Inside a function body. Variables declared here are local to that function only and cannot be accessed from outside.

```go
package main

import "fmt"

func main() {
    message := "Hello"  // local to main() - not accessible elsewhere
    fmt.Println(message)
}

func other() {
    // message is not accessible here - different function scope
}
```

### 4. Control flow blocks
Scope: Inside control structures. Variables declared here exist only within that specific block.

**if/else block:**
```go
package main

import "fmt"

func main() {
    age := 25

    if age >= 18 {
        status := "adult"  // only exists inside this if block
        fmt.Println(status)
    }
    // status is not accessible here
}
```

**for loop block:**
```go
package main

import "fmt"

func main() {
    for i := 1; i <= 3; i++ {
        result := i * 2  // only exists inside this loop iteration
        fmt.Println(result)
    }
    // result is not accessible here - loop scope ended
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
        fmt.Println(dayName)
    case 2:
        dayName := "Tuesday"  // different block, can reuse the name
        fmt.Println(dayName)
    }
    // dayName is not accessible here
}
```

### 5. Explicit blocks
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

1. **Nested scopes**: Inner blocks can see and use variables from outer blocks, but outer blocks cannot see variables declared in inner blocks.

2. **Variable lifetime**: A variable exists only from its declaration to the end of its block. Once a block closes, its variables are gone.

3. **Hierarchy**: Blocks follow a hierarchy:
   - Package ← File ← Function ← Control Flow ← Explicit Blocks
   - Variables are accessible from inner to outer, but not the other way around.

4. **Shadowing risk**: You can declare a variable with the same name in an inner block, which hides (shadows) the outer variable.