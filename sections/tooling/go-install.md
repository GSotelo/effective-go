# go install

The `go install` command compiles and installs Go packages and their dependencies. It builds the named packages and places the resulting binaries in `$GOBIN` (or `$GOPATH/bin` if `$GOBIN` is not set), making them available for use from the command line. It is the standard way to install Go-based tools and executables.

# Configuring the installation path

To set a default for `GOBIN`, use `go env -w`:

```go
go env -w GOBIN=path/to/your/bin
```