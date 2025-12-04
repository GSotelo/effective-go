# iota

`iota` is a predeclared constant that automatically generates incrementing values within a constant declaration block. It starts at 0 and increments by 1 for each new constant.

## Simple Example

```go
package main

import "fmt"

const (
	Free   = iota // 0
	Basic         // 1 - iota is implicit
	Pro           // 2 - iota is implicit
	Custom        // 3 - iota is implicit
)

func main() {
	fmt.Println("Cost tiers:")
	fmt.Println("Free:", Free)
	fmt.Println("Basic:", Basic)
	fmt.Println("Pro:", Pro)
	fmt.Println("Custom:", Custom)
}
```

**Output:**
```
Cost tiers:
Free: 0
Basic: 1
Pro: 2
Custom: 3
```

## Key Points

- `iota` resets to 0 at the start of each `const` block
- After the first constant, you can omit `= iota` and it will automatically increment for subsequent constants
- `iota` is useful for creating enumerations and bit flags