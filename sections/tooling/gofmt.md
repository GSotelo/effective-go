# gofmt

The `gofmt` tool automatically formats Go source code to follow the standard Go formatting conventions. It ensures consistent code style across all Go projects, eliminating debates about formatting and making code easier to read and maintain.

## gofmt vs go fmt

`gofmt` is the standalone formatting tool that operates on individual files or directories. `go fmt` is a convenience wrapper that runs `gofmt -l -w` on the packages named by the import paths, integrating seamlessly with Go's module system.

## Usage

Format a single file:

```bash
gofmt -w main.go
```

Format all Go files in the current directory:

```bash
gofmt -w .
```

Using the `go fmt` wrapper on packages:

```bash
go fmt ./...
```

## Common flags

| Flag | Description |
|------|-------------|
| `-w` | Write result to source file instead of stdout |
| `-l` | List files whose formatting differs from gofmt's |
| `-d` | Display diffs instead of rewriting files |
| `-s` | Simplify code in addition to formatting |

## Example: Before and after formatting

Before running `gofmt`:

```go
package main

import "fmt"

func main(){
x:=1
if x>0{
fmt.Println("positive")
}
}
```

After running `gofmt -w main.go`:

```go
package main

import "fmt"

func main() {
	x := 1
	if x > 0 {
		fmt.Println("positive")
	}
}
```

## Best practices

- Run `gofmt` or `go fmt` before committing code
- Configure your editor to format on save
- Add formatting checks to your CI pipeline with `gofmt -l`
- Use `gofmt -s` to apply additional simplifications like removing unnecessary type declarations from composite literals