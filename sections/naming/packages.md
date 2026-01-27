# Package naming

## Keep it short and simple

Package names should be short, lowercase, and single-word. Avoid underscores or camelCase. Brevity matters because everyone types it.

```go
// Good
package json
package http
package bytes

// Bad
package jsonParser
package http_utils
package ByteBuffer
```

## Package name matches directory name

If your package lives in `src/encoding/json`, the package name is `json`, not `encoding_json`.

```go
// Directory: src/compress/gzip
package gzip // not compress_gzip
```

## Avoid stuttering

Don't repeat the package name in exported identifiers. Users already see the package prefix when calling your code.

```go
// Good: users call http.Client
type Client struct{}

// Bad: users would call http.HttpClient (redundant)
type HttpClient struct{}
```

```go
// Good: json.Encoder
type Encoder struct{}

// Bad: json.JSONEncoder (stutters)
type JSONEncoder struct{}
```

## Leverage the package prefix

The package name provides context. Use it to keep exported names concise.

```go
// Good
list.New()      // not list.NewList()
context.WithCancel() // not context.WithContextCancel()
bytes.Buffer    // not bytes.ByteBuffer
```

## Import collisions are okay

If two packages share the same name, alias one at import time. Don't preemptively make names longer to avoid hypothetical conflicts.

```go
import (
    "crypto/rand"
    mrand "math/rand"
)
```

The package name is part of how users reference your code. Use `pkg.Thing`, not `pkg.PkgThing`.