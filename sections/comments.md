# Comments

Go supports C-style (`/* */`) and C++-style (`//`) comments.

## Line comments

Line comments are the recommended form. They start with `//` and extend to the end of the line.

```go
// This is a line comment
x := 10 // inline comment
```

## Block comments

Block comments use `/* */` syntax. Use them for package documentation or to disable large blocks of code during debugging.

```go
/*
Package math provides basic mathematical functions.
This is the package documentation.
*/
package math

/*
func oldImplementation() {
    // temporarily disabled
}
*/
```