# Variable Shadowing

## Definition

Variable shadowing happens when you create a new variable with the same name inside a smaller scope (like inside a loop or function). This new variable "hides" or "shadows" the outer variable with the same name, making it invisible within that inner scope. Once you exit the inner scope, the original outer variable is accessible again.

```go
package main

import "fmt"

func main() {
    x := 100  // outer scope
    fmt.Println("Before loop, x =", x)

    i := 0
    for i < 1 {
        x := 50  // inner scope - shadows outer x
        fmt.Println("Inside loop, x =", x)
        i++
    }

    fmt.Println("After loop, x =", x)  // x is still 100
}
```

Output:
```
Before loop, x = 100
Inside loop, x = 50
After loop, x = 100
```

The shadowed `x` inside the loop only exists within the loop scope. Outside the loop, the original `x` (100) is visible.
