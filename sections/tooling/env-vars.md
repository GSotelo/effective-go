# Environment variables

Go uses environment variables to control the behavior of its toolchain. Two key variables, `GOBIN` and `GOPATH`, determine where `go install` places compiled binaries.

## Install directory resolution

```mermaid
flowchart TD
    A[go install] --> B{GOBIN set?}
    B -- Yes --> C[Install to $GOBIN]
    B -- No --> D{GOPATH set?}
    D -- Yes --> E["Install to bin subdirectory of\nfirst directory in $GOPATH list"]
    D -- No --> F["Install to default GOPATH\n$HOME/go/bin · %USERPROFILE%\\go\\bin"]
```

## Variables

| Variable | Description |
|----------|-------------|
| `GOBIN` | Directory where `go install` places compiled binaries. Takes precedence over `GOPATH`. |
| `GOPATH` | Workspace root. Binaries are installed to the `bin` subdirectory of its first listed directory. Defaults to `$HOME/go` on Unix or `%USERPROFILE%\go` on Windows. |