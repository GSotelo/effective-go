# Concurrency in Go

Concurrency is the composition of independently executing computations. It's about structuring your program to handle multiple tasks that can make progress independently, even if they don't necessarily run at the same time.

## Concurrency is NOT Parallelism

This is a crucial distinction that often confuses newcomers to Go:

- **Concurrency** is about dealing with lots of things at once. It's a way to structure your program by breaking it into pieces that can execute independently. A concurrent program can run on a single processor by interleaving the execution of different tasks.

- **Parallelism** is about doing lots of things at once. It requires multiple processors or cores actually executing tasks simultaneously at the same moment in time.

Think of it this way: concurrency is about the structure and design of your program, while parallelism is about the execution. A well-designed concurrent program can run in parallel when multiple processors are available, but it doesn't require parallelism to be concurrent.