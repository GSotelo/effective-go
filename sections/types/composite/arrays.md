# Arrays in Go

Arrays in Go are fixed-size, ordered sequences of elements of the same type. They're a fundamental data structure, though often less common in everyday Go programming than slices due to their fixed size.

## Declaring Arrays

In Go, you can declare arrays in several ways:

```go
// Declare an array of 5 integers (zero-initialized)
var arr [5]int

// Declare and initialize with values
var arr2 = [5]int{1, 2, 3, 4, 5}

// Short declaration syntax
arr3 := [5]int{1, 2, 3, 4, 5}
```
