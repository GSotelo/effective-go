// Package main is the entry point that runs all snippets.
package main

import (
	"fmt"

	"github.com/gsotelo/effective-go/snippets/interfaces"
	"github.com/gsotelo/effective-go/snippets/variadic"
)

func main() {
	variadic.Run()
	fmt.Println()
	interfaces.Run()
}
