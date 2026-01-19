# net/http Package

The `net/http` package provides HTTP client and server implementations. It's one of Go's most powerful standard library packages for building web applications.

## HTTP Server Basics

The simplest way to start an HTTP server is with `http.ListenAndServe` and `http.HandleFunc`.

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })

    http.ListenAndServe(":8080", nil)
}
```

The `http.Handler` interface is the foundation of HTTP handling in Go:

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

Any type that implements `ServeHTTP` can handle HTTP requests.

## Request Handling

The `http.Request` struct contains all information about an incoming request.

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Method: GET, POST, etc.
    method := r.Method

    // URL path
    path := r.URL.Path

    // Query parameters: /search?q=golang
    query := r.URL.Query().Get("q")

    // Headers
    contentType := r.Header.Get("Content-Type")

    fmt.Fprintf(w, "Method: %s, Path: %s", method, path)
}
```

Parsing JSON from request body:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    var data struct {
        Name string `json:"name"`
    }

    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    fmt.Fprintf(w, "Hello, %s!", data.Name)
}
```

## Response Writing

The `http.ResponseWriter` interface is used to construct an HTTP response.

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Set headers before writing body
    w.Header().Set("Content-Type", "application/json")

    // Set status code (default is 200 OK)
    w.WriteHeader(http.StatusCreated)

    // Write response body
    w.Write([]byte(`{"message": "created"}`))
}
```

Writing JSON responses:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    data := map[string]string{"message": "Hello, World!"}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}
```

## HTTP Client

Making HTTP requests with the client:

```go
// Simple GET request
resp, err := http.Get("https://api.example.com/data")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

body, err := io.ReadAll(resp.Body)
```

Using `http.Client` for more control:

```go
client := &http.Client{
    Timeout: 10 * time.Second,
}

req, err := http.NewRequest("GET", "https://api.example.com/data", nil)
if err != nil {
    log.Fatal(err)
}

req.Header.Set("Authorization", "Bearer token")

resp, err := client.Do(req)
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()
```

POST request with JSON body:

```go
data := map[string]string{"name": "gopher"}
jsonData, _ := json.Marshal(data)

resp, err := http.Post(
    "https://api.example.com/users",
    "application/json",
    bytes.NewBuffer(jsonData),
)
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()
```

## Common Patterns

### Graceful Shutdown

```go
func main() {
    srv := &http.Server{Addr: ":8080"}

    go func() {
        if err := srv.ListenAndServe(); err != http.ErrServerClosed {
            log.Fatal(err)
        }
    }()

    // Wait for interrupt signal
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt)
    <-quit

    // Graceful shutdown with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    srv.Shutdown(ctx)
}
```

### Context with Requests

```go
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    // Use context for cancellation
    select {
    case <-time.After(2 * time.Second):
        fmt.Fprintf(w, "completed")
    case <-ctx.Done():
        log.Println("request cancelled")
        return
    }
}
```

### Error Handling

```go
func handler(w http.ResponseWriter, r *http.Request) {
    data, err := fetchData()
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}
```