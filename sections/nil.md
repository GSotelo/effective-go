# What is nil?

In Go, `nil` is a predeclared identifier available in the universal block. It is frequently used to denote the zero value of several types, including pointers, channels, functions, interfaces, maps, and slices.

# When to use nil?

## Error handling
The standard practice in Go is for functions to return an error type as their terminal return value. A nil value denotes successful execution, whereas a non-nil value indicates an error. This explicit mechanism for error reporting distinguishes Go from languages that rely on exceptions or specific error codes. More details are provided in the example below.

```go
package main

import "fmt"

func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Division by zero")
	}
	return a / b, nil
}

func main() {
	result, err := safeDivide(0, 0)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", result)
	}
}
```
