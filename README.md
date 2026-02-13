# Effective Go

Writing idiomatic Go makes your code cleaner, more consistent, and way easier to maintain. Trying to copy patterns from other languages usually just leads to clunky, non-Go-like code. This repo is here to guide you through the Go ecosystem with simple, hands-on examples that show off its coolest features one step at a time.

<br>

<div align="center">
    <img src="assets/go-motorcycle.svg" alt="Go Programming Guide" width="200">
</div>

## Table of contents

- [Blocks](sections/blocks.md)
- [Declarations](sections/declarations.md)
- [Comments](sections/comments.md)
- [Shadowing](sections/shadowing.md)
- [Nil](sections/nil.md)
- [Naming](sections/naming/overview.md)
  - [Packages](sections/naming/packages.md)
- [Types](sections/types/types.md)
  - [Integer](sections/types/predeclared/integer.md)
  - [Floating-point](sections/types/predeclared/floating-point.md)
  - [Complex](sections/types/predeclared/complex.md)
  - [Strings](sections/types/predeclared/strings.md)
  - [Iota](sections/types/predeclared/iota.md)
  - [Arrays](sections/types/composite/arrays.md)
  - [Slices](sections/types/composite/slices.md)
  - [Maps](sections/types/composite/maps.md)
  - [Pointers](sections/types/pointers.md)
  - [Channels](sections/types/channels.md)
- [Functions](sections/functions.md)
- [Interfaces](sections/types/interfaces/overview.md)
  - [Implicit interfaces](sections/types/interfaces/implicit-interfaces.md)
  - [Dependency injection](sections/types/interfaces/dependency-injection.md)
- [Errors](sections/errors/sentinel-errors.md)
  - [Sentinel errors](sections/errors/sentinel-errors.md)
  - [Panic and recover](sections/errors/panic-recover.md)
- [Concurrency](sections/concurrency.md)
- [Packages](sections/packages.md)
- [Standard library](sections/stdlib/net-http.md)
  - [net/http](sections/stdlib/net-http.md)
- [Tooling](sections/tooling/overview.md)
  - [gofmt](sections/tooling/gofmt.md)
  - [govulncheck](sections/tooling/govulncheck.md)