# Panic and Recover

**panic** is a built-in function that stops the normal execution flow of the current goroutine and begins unwinding the stack, executing any deferred functions along the way. When a panic reaches the top of a goroutine's call stack without being recovered, the program terminates and prints a stack trace.

**recover** is a built-in function that regains control of a panicking goroutine, but only when called inside a deferred function. It stops the panic sequence, returns the value passed to panic, and allows the program to continue execution.

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

**Panic and recover are NOT exceptions.** While they may superficially resemble try/catch blocks from other languages, that's not their intended purpose in Go. The Go philosophy strongly favors explicit error handling through return values rather than control-flow-by-panic.

**Panic indicates fatal, unrecoverable situations.** You typically don't want to continue execution after a panic occurs. Use panic for programmer errors (like nil pointer dereferences, out-of-bounds array access) or truly exceptional circumstances that represent broken invariants—situations where continuing would lead to undefined behavior or data corruption.

**Recover enables graceful shutdown and logging.** The primary legitimate use of recover is to prevent a single panic from bringing down an entire service. When you recover, you should log the panic details to a monitoring system, perform any necessary cleanup, and usually exit the program gracefully rather than continuing normal operation.

**Library authors must not leak panics.** If you're building a library, it's your responsibility to catch panics within your code and convert them to errors. Never let a panic escape your package boundary—recover from internal panics and return proper error values instead. The library consumer should make decisions about error handling, not have panics forced upon them.

**Only deferred functions execute after panic.** When a panic occurs, the normal execution path stops immediately. The only code that runs during stack unwinding is deferred functions, which execute in LIFO (last-in, first-out) order. This is why recover must be called from within a defer—it's the only opportunity to intercept the panic.

**Panics bubble up through call stacks, not across goroutines.** A panic propagates up the call stack of its goroutine until it's recovered or reaches the top. If you're using goroutines, a panic only affects that specific goroutine's stack—it bubbles up to the function that launched the goroutine, not beyond. An unrecovered panic in a goroutine will crash the entire program, so always consider adding recovery mechanisms in goroutine entry points.