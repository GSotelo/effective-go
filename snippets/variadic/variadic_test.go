package variadic

import "fmt"

func ExampleSum() {
	fmt.Println(Sum())
	fmt.Println(Sum(1))
	fmt.Println(Sum(1, 2, 3))
	// Output:
	// 0
	// 1
	// 6
}

func ExampleSum_spread() {
	values := []int{10, 20, 30, 40}
	fmt.Println(Sum(values...))
	// Output: 100
}

func ExamplePrintItems() {
	PrintItems("Fruits", "apple", "banana", "cherry")
	// Output: Fruits: [apple, banana, cherry]
}
