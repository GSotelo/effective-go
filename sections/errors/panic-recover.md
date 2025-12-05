# Panic and Recover

**panic** is a runtime condition that represents an unrecoverable error state where normal execution cannot continue. It occurs either when the Go runtime encounters critical errors (nil pointer dereferences, out-of-bounds array access, type assertion failures) or when explicitly triggered by calling the panic() built-in function.

**recover** is the mechanism to intercept and handle a panic, preventing it from crashing the program. It only functions when called directly from within a deferred function, where it captures the panic value and halts the stack unwinding process, allowing the program to regain control.

## Panic Example

```go
func processData(data []int) {
	if len(data) == 0 {
		panic("cannot process empty data")
	}
}
```

## Recover Example

```go
func safeProcess() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered from panic: %v\n", r)
		}
	}()
	panic("something went wrong")
}
```

## Philosophy and Best Practices

- **Panic and recover are NOT exceptions.** Unlike try/catch in other languages, Go favors explicit error handling through return values, not control-flow-by-panic.

- **Panic indicates fatal, unrecoverable situations.** Use panic only for programmer errors (nil dereferences, out-of-bounds access) or broken invariants where continuing would cause undefined behavior or data corruption.

- **Recover enables graceful shutdown and logging.** The primary use of recover is preventing a single panic from crashing your service. Log panic details to monitoring systems, clean up resources, and typically exit gracefully rather than resuming normal operation.

- **Library authors must not leak panics.** Libraries must recover from internal panics and return errors instead. Never let panics escape package boundaries—let library consumers decide how to handle errors.

- **Only deferred functions execute after panic.** Normal execution stops immediately on panic. Deferred functions execute in LIFO order during stack unwinding, which is why recover only works inside defer.

- **Panics bubble up through call stacks, not across goroutines.** A panic propagates up its goroutine's stack until recovered or it reaches the top. Panics don't cross goroutine boundaries, but an unrecovered panic in any goroutine crashes the entire program.