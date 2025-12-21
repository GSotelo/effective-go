# Functions

## Variadic functions

Variadic functions accept a variable number of arguments of the same type. They provide a flexible way to write functions that can handle zero or more arguments without requiring the caller to create a slice explicitly.

### Basic usage

A variadic parameter is declared using three dots `...` before the type. The variadic parameter must be the last parameter in the function signature.

```go
package main

import "fmt"

// sum accepts zero or more integers and returns their sum
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}

func main() {
    fmt.Println(sum())           // 0 - no arguments
    fmt.Println(sum(1))          // 1 - single argument
    fmt.Println(sum(1, 2, 3))    // 6 - multiple arguments
    fmt.Println(sum(1, 2, 3, 4, 5)) // 15 - many arguments
}
```

Inside the function, the variadic parameter is treated as a slice of the specified type:

```go
package main

import "fmt"

func printItems(label string, items ...string) {
    fmt.Printf("%s: [", label)
    for i, item := range items {
        if i > 0 {
            fmt.Print(", ")
        }
        fmt.Print(item)
    }
    fmt.Println("]")
    fmt.Printf("Type of items: %T\n", items) // []string
}

func main() {
    printItems("Fruits", "apple", "banana", "cherry")
    // Output:
    // Fruits: [apple, banana, cherry]
    // Type of items: []string
}
```

### Spreading a slice

When you already have a slice, you can pass it to a variadic function using the spread operator `...` after the slice:

```go
package main

import "fmt"

func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}

func main() {
    values := []int{10, 20, 30, 40}

    // Spread the slice using ...
    result := sum(values...)
    fmt.Println("Sum:", result) // 100

    // This is equivalent to calling:
    // sum(10, 20, 30, 40)
}
```

You can also combine individual arguments with a spread slice:

```go
package main

import "fmt"

func concat(separator string, parts ...string) string {
    result := ""
    for i, part := range parts {
        if i > 0 {
            result += separator
        }
        result += part
    }
    return result
}

func main() {
    words := []string{"Go", "is", "awesome"}

    // Mix spread slice with additional arguments
    sentence := concat(" ", words...)
    fmt.Println(sentence) // "Go is awesome"

    // Can also add more arguments after spreading
    moreParts := []string{"simple", "and"}
    sentence2 := concat(" ", moreParts...)
    fmt.Println(sentence2) // "simple and"
}
```

### Combining regular and variadic parameters

A function can have both regular parameters and a variadic parameter. The variadic parameter must come last:

```go
package main

import "fmt"

// Valid: regular parameters followed by variadic parameter
func greetAll(greeting string, names ...string) {
    for _, name := range names {
        fmt.Printf("%s, %s!\n", greeting, name)
    }
}

func main() {
    greetAll("Hello", "Alice", "Bob", "Charlie")
    // Output:
    // Hello, Alice!
    // Hello, Bob!
    // Hello, Charlie!
}
```

```go
// Invalid: variadic parameter must be last
func invalid(numbers ...int, multiplier int) int {  // compilation error
    // ...
}

// Valid: regular parameter before variadic
func multiply(multiplier int, numbers ...int) []int {
    results := make([]int, len(numbers))
    for i, n := range numbers {
        results[i] = n * multiplier
    }
    return results
}
```

### Empty variadic parameters

A variadic function can be called with zero arguments for the variadic parameter. Inside the function, this results in a nil slice:

```go
package main

import "fmt"

func printCount(items ...string) {
    if items == nil {
        fmt.Println("No items provided (nil slice)")
    } else {
        fmt.Printf("Received %d items\n", len(items))
    }
}

func main() {
    printCount()                    // No items provided (nil slice)
    printCount("apple")             // Received 1 items
    printCount("apple", "banana")   // Received 2 items
}
```

Be aware of the difference between no arguments and an empty slice:

```go
package main

import "fmt"

func inspect(values ...int) {
    if values == nil {
        fmt.Println("nil slice")
    } else if len(values) == 0 {
        fmt.Println("empty slice")
    } else {
        fmt.Printf("slice with %d elements\n", len(values))
    }
}

func main() {
    inspect()                          // nil slice - no arguments passed

    emptySlice := []int{}
    inspect(emptySlice...)             // empty slice - explicit empty slice spread

    inspect(1, 2, 3)                   // slice with 3 elements
}
```

### Key points

- Variadic parameters use `...` before the type: `func f(args ...Type)`
- The variadic parameter must be the last parameter in the function signature
- Inside the function, the variadic parameter is a slice: `[]Type`
- Variadic functions can be called with zero or more arguments
- Use the spread operator `...` to pass a slice to a variadic function: `f(mySlice...)`
- A nil slice is passed when no arguments are provided for the variadic parameter
- Only one variadic parameter is allowed per function, and it must be last

## Named return values

Named return values are pre-declared variables in the function signature that can be referenced and assigned within the function body, then explicitly returned by specifying them in the return statement.

```go
func divide(a, b float64) (result float64, err error) {
    if b == 0 {
        err = fmt.Errorf("division by zero")
        return result, err
    }
    result = a / b
    return result, err
}
```

> **Warning**: Avoid using bare return statements with named return values. Bare returns make it difficult to track how data flows through your function. When using a bare return, the last value assigned to each named return variable is what gets returned. If the function receives invalid input and returns early without assignment, the zero values of those return types are returned instead.

```go
func divide(a, b float64) (result float64, err error) {
    if b == 0 {
        err = fmt.Errorf("division by zero")
        return // returns (0.0, error) - result is zero value
    }
    result = a / b
    return // returns (calculated result, nil) - err is zero value
}
```

## Functions as values

Functions are first-class values in Go, meaning they can be assigned to variables, passed as arguments, and returned from other functions. This enables powerful functional programming patterns and flexible code design.

```go
package main

import "fmt"

// Define a function type that takes two integers and returns an integer
type Operation func(int, int) int

// add is a function that matches the Operation signature
func add(a, b int) int {
    return a + b
}

// multiply is another function that matches the Operation signature
func multiply(a, b int) int {
    return a * b
}

func main() {
    // Assign functions to variables of type Operation
    var op Operation

    op = add
    fmt.Println("5 + 3 =", op(5, 3))    // 8

    op = multiply
    fmt.Println("5 * 3 =", op(5, 3))    // 15

    // Functions can also be passed directly without declaring the type
    calculate := func(x, y int, operation Operation) int {
        return operation(x, y)
    }

    fmt.Println("Calculate with add:", calculate(10, 2, add))        // 12
    fmt.Println("Calculate with multiply:", calculate(10, 2, multiply)) // 20
}
```

## Anonymous functions

Anonymous functions are function literals defined without a name. They can be created inline and used in several ways: assigned to variables for later invocation, passed directly as arguments to other functions, or executed immediately at the point of definition. Anonymous functions have access to variables from their enclosing scope, forming closures.

```go
func main() {
    square := func(n int) int { return n * n }
    fmt.Println(square(5)) // 25
}
```

> **Note**: Anonymous functions are commonly used with `defer` and goroutines because they can capture variables from the surrounding scope. This lets you write cleanup logic or start concurrent operations inline without creating separate named functions.