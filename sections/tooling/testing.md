# Testing conventions

Go recognizes a few special directories and file patterns for testing. The `go` tool treats them differently from regular source files.

## testdata/

A directory named `testdata` is ignored by `go build`, `go vet`, and the Go toolchain in general — it's never treated as package source. Use it to store fixture files your tests read from disk (JSON, golden files, sample inputs, etc.).

```
myapp/
├── parser.go
├── parser_test.go
└── testdata/
    └── input.json
```

```go
func TestParse(t *testing.T) {
    data, err := os.ReadFile("testdata/input.json")
    if err != nil {
        t.Fatal(err)
    }
    // ...
}
```

## Example functions

A function named `ExampleXxx` in a `_test.go` file is both documentation and a test. It's shown in godoc next to the thing it documents, and if it has an `// Output:` comment, `go test` runs it and checks the output matches.

```go
func ExampleAdd() {
    fmt.Println(Add(2, 3))
    // Output: 5
}
```

If the printed output doesn't match the `// Output:` comment, the test fails.

## Table-driven tests

A common pattern is a slice of test cases, each with a `wantErr bool` field to check if an error was expected.

```go
tests := []struct {
    name    string
    input   string
    wantErr bool
}{
    {"valid input", "123", false},
    {"invalid input", "abc", true},
}

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        _, err := Parse(tt.input)
        if (err != nil) != tt.wantErr {
            t.Errorf("got err = %v, wantErr %v", err, tt.wantErr)
        }
    })
}
```