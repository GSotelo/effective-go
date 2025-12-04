# Sentinel Errors

Sentinel errors are package-level, predefined error values that represent specific, anticipated failure conditions important enough to expose as part of your public API. They serve as identifiable markers that enable callers to make explicit control-flow decisions through identity comparison (using `==`).

By declaring sentinel errors at package scope, you create an explicit contract with consumers: these error values become stable identifiers that users will depend on in their code, committing you to maintain them across future versions of your package.

## Simple Example

```go
package main

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound  = errors.New("user not found")
	ErrInvalidID = errors.New("invalid user ID")
)

func FindUser(id int) (string, error) {
	if id <= 0 {
		return "", ErrInvalidID
	}
	if id > 100 {
		return "", ErrNotFound
	}
	return fmt.Sprintf("User_%d", id), nil
}

func main() {
	_, err := FindUser(999)
	if err == ErrNotFound {
		fmt.Println("We don't have that user in our system")
	}

	_, err = FindUser(-5)
	if err == ErrInvalidID {
		fmt.Println("Please provide a positive ID number")
	}
}
```

## Key Points

- Sentinel errors are named with the prefix `Err` by convention
- They're declared using `errors.New()` at the package level
- You compare them by identity using `==`
- They enable programmatic decisions—callers decide whether to continue, retry, use defaults, or fail
- Declaring a sentinel error is an API contract—removing or renaming it in future versions is a breaking change
- Document all sentinel errors in your package's public API godoc