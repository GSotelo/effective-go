// Package variadic demonstrates variadic functions in Go.
// Variadic functions accept a variable number of arguments of the same type.
package variadic

import "fmt"

// Sum accepts zero or more integers and returns their sum.
func Sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

// PrintItems demonstrates that variadic parameters are treated as slices.
func PrintItems(label string, items ...string) {
	fmt.Printf("%s: [", label)
	for i, item := range items {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(item)
	}
	fmt.Println("]")
}

// Run executes all variadic examples.
func Run() {
	fmt.Println("=== Variadic Functions ===")
	fmt.Println()

	// Basic usage
	fmt.Println("Sum():", Sum())
	fmt.Println("Sum(1):", Sum(1))
	fmt.Println("Sum(1, 2, 3):", Sum(1, 2, 3))
	fmt.Println("Sum(1, 2, 3, 4, 5):", Sum(1, 2, 3, 4, 5))
	fmt.Println()

	// Spreading a slice
	values := []int{10, 20, 30, 40}
	fmt.Println("Sum(values...):", Sum(values...))
	fmt.Println()

	// Combining regular and variadic parameters
	PrintItems("Fruits", "apple", "banana", "cherry")
}
