# Go Tooling

Go's tooling ecosystem is one of its greatest strengths, providing a comprehensive suite of built-in tools that standardize development workflows across teams and projects. These tools handle everything from code formatting and testing to security scanning and performance profiling, making Go development productive and consistent.

## govulncheck

The `govulncheck` tool scans your Go code and dependencies for known security vulnerabilities from the Go vulnerability database.

### Installation

Install govulncheck using `go install`:

```bash
go install golang.org/x/vuln/cmd/govulncheck@latest
```

### Usage

Run govulncheck in your project directory:

```bash
govulncheck ./...
```

### Example: Detecting vulnerable dependencies

Consider a simple program that uses an older version of the `golang.org/x/text` package with a known vulnerability:

```go
package main

import (
    "fmt"
    "golang.org/x/text/language"
    "golang.org/x/text/message"
)

func main() {
    p := message.NewPrinter(language.English)
    p.Printf("Hello, %s!\n", "world")
    fmt.Println("Application running...")
}
```

```go
// go.mod
module example.com/myapp

go 1.25.5

require golang.org/x/text v0.3.0
```

When you run `govulncheck ./...` on this code, it detects the vulnerability:

```
=== Symbol Results ===

No vulnerabilities found.

Your code is affected by 0 vulnerabilities.
This scan also found 2 vulnerabilities in packages you import and 1
vulnerability in modules you require, but your code doesn't appear to call these
vulnerabilities.
Use '-show verbose' for more details.
```

To fix the vulnerability, update the dependency to the patched version:

```bash
go get -u=patch ./...
go mod tidy
```

After updating, running `govulncheck ./...` again should report no vulnerabilities:

```
No vulnerabilities found.
```

### Best practices

- Run `govulncheck` regularly as part of your CI/CD pipeline
- Integrate it into pre-deployment security checks
- Keep dependencies updated to minimize vulnerability exposure
- Use `go mod tidy` to remove unused dependencies that might contain vulnerabilities