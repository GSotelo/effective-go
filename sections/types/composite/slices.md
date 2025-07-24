# Slices

Slices are flexible, dynamically sized data structures whose length isn’t included in their type. While working with slices is similar to working with arrays, slices offer greater versatility, such as allowing functions to accept slices of any size. Although slices and arrays share some similarities, there are also important differences. Let’s start by looking at the similarities and create some example slices:

```go
var numbers = []int{100, 200, 300}
	fmt.Printf("Slice with literal: %v\n", numbers) // [100 200 300]
```

You can create a slice literal with specific indexes initialized to predefined values, while the remaining elements are set to their zero values. For example:

```go
var numbers = []int{0: 100, 2: 200}
	fmt.Printf("Slice with predefined index values: %v\n", numbers) // [100 0 200]
```

Now, let’s begin by examining the differences.

Core Ideas:
- Slices are not comparable. The only exception is when comparing a slice to `nil`.
