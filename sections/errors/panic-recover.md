# Panic and Recover

**Panic** is a built-in function that stops normal execution of the current goroutine. When a function calls `panic`, normal execution stops, all deferred functions are executed, and then the function returns to its caller. This process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes.

**Recover** is a built-in function that regains control of a panicking goroutine. It only works when called directly inside a deferred function—calling `recover` outside a defer or when the goroutine isn't panicking returns `nil` and has no effect.

## When Panics Occur

Panics happen either explicitly via `panic()` or implicitly when the runtime detects unrecoverable errors:

```go
// Explicit panic for broken invariants
func MustCompile(pattern string) *Regexp {
	re, err := Compile(pattern)
	if err != nil {
		panic("regexp: Compile(" + pattern + "): " + err.Error())
	}
	return re
}

// Runtime panics (implicit)
var s []int
_ = s[0]          // panic: index out of range

var m map[string]int
m["key"] = 1      // panic: assignment to entry in nil map

var p *int
_ = *p            // panic: nil pointer dereference

var i interface{} = "string"
_ = i.(int)       // panic: interface conversion
```

## Basic Recover Pattern

Recover must be called directly within a deferred function:

```go
func safeCall(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}
	}()
	fn()
	return nil
}

func main() {
	err := safeCall(func() {
		panic("something went wrong")
	})
	fmt.Println(err) // panic recovered: something went wrong
}
```

## HTTP Server Recovery Middleware

A common real-world pattern—prevent one bad request from crashing your entire server:

```go
func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Log the panic with stack trace for debugging
				log.Printf("panic: %v\n%s", err, debug.Stack())

				// Return 500 to client
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
```

## Library Pattern: Convert Panic to Error

Libraries should never let panics escape to callers. Convert internal panics to errors:

```go
package parser

// Parse returns an error instead of panicking
func Parse(input string) (result *AST, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("parse error: %v", r)
		}
	}()
	return parse(input), nil
}

// Internal function may panic for simplicity
func parse(input string) *AST {
	if input == "" {
		panic("empty input")
	}
	// ... parsing logic that might panic on malformed input
	return &AST{}
}
```

## Goroutine Recovery

Panics don't cross goroutine boundaries. Each goroutine must handle its own panics:

```go
func worker(jobs <-chan int, results chan<- int) {
	for job := range jobs {
		// Each job gets its own recovery
		func() {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("worker panic on job %d: %v", job, r)
					results <- -1 // Signal failure
				}
			}()
			results <- process(job)
		}()
	}
}
```

## Key Points

- **Panic is not for normal error handling.** Use `error` return values for expected failures. Panic is for truly exceptional situations: programmer errors, broken invariants, or impossible conditions.

- **Recover only works in deferred functions.** Calling `recover()` in normal code does nothing—it must be inside a `defer` to catch panics.

- **Libraries must not leak panics.** Always recover at package boundaries and convert panics to errors. Let callers decide how to handle failures.

- **Each goroutine needs its own recovery.** A panic in one goroutine cannot be recovered by another. An unrecovered panic in any goroutine crashes the entire program.

- **Use `debug.Stack()` for diagnostics.** When recovering from panics, log the stack trace to help debug the root cause.

- **Prefer `Must` prefix for panic-on-error helpers.** Functions like `regexp.MustCompile` or `template.Must` signal that they panic instead of returning errors—use them only with compile-time constants.