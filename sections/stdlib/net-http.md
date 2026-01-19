# net/http package

The `net/http` package provides HTTP client and server implementations. It's one of Go's most powerful standard library packages for building web applications.

## HTTP server basics

The simplest way to start an HTTP server is with `http.ListenAndServe` and `http.HandleFunc`.

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })

    log.Fatal(http.ListenAndServe(":8080", nil)) // nil uses the default ServeMux
}
```

## ServeMux

A `ServeMux` is a request multiplexer (hence "mux") that matches incoming requests to handlers based on URL paths.

**What ServeMux provides:**

- Exact path matches: `/users` matches only `/users`
- Prefix matches with trailing slash: `/users/` matches `/users/anything`
- **Go 1.22+**: Method routing and path parameters

**What ServeMux lacks:**

- No regex matching
- No middleware chain support

```go
mux := http.NewServeMux()
mux.HandleFunc("/users", handleUsers)      // exact match
mux.HandleFunc("/users/", handleUserByID)  // prefix match: catches /users/*

log.Fatal(http.ListenAndServe(":8080", mux))
```

**Go 1.22+ enhancements:** ServeMux now supports method routing and path parameters:

```go
mux := http.NewServeMux()

// Method-specific routing
mux.HandleFunc("GET /users", listUsers)
mux.HandleFunc("POST /users", createUser)

// Path parameters with {name} syntax
mux.HandleFunc("GET /users/{id}", getUser)
mux.HandleFunc("DELETE /users/{id}", deleteUser)

log.Fatal(http.ListenAndServe(":8080", mux))
```

```go
func getUser(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")  // Go 1.22+: extract path parameter
    fmt.Fprintf(w, "User ID: %s", id)
}
```

For Go versions before 1.22, extract path parameters manually:

```go
func handleUserByID(w http.ResponseWriter, r *http.Request) {
    // /users/123 -> "123"
    id := strings.TrimPrefix(r.URL.Path, "/users/")
    fmt.Fprintf(w, "User ID: %s", id)
}
```

**Default vs custom ServeMux:**

```go
// Using the default (global) ServeMux
http.HandleFunc("/", homeHandler)            // registers to http.DefaultServeMux
log.Fatal(http.ListenAndServe(":8080", nil)) // nil = use DefaultServeMux

// Using your own ServeMux
mux := http.NewServeMux()
mux.HandleFunc("/", homeHandler)             // registers to your mux
log.Fatal(http.ListenAndServe(":8080", mux)) // pass your mux explicitly
```

The default `ServeMux` is a global variable (`http.DefaultServeMux`). This is fine for prototypes, but in larger projects create your own—any package (including third-party dependencies) can register routes on the global default without your knowledge.

**For real routing**, consider these popular alternatives:

- `chi` - lightweight, idiomatic (recommended)
- `httprouter` - high performance
- `gin` - full web framework

## Handler interface

The `http.Handler` interface is Go's core abstraction for handling HTTP requests.

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

Any type that has a `ServeHTTP(w http.ResponseWriter, r *http.Request)` method automatically becomes a Handler.

**Why it matters:**

- ServeMux expects Handlers
- Middleware wraps Handlers
- The HTTP server uses Handlers to process requests

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

// Always check status code
if resp.StatusCode != http.StatusOK {
    log.Fatalf("unexpected status: %s", resp.Status)
}

body, err := io.ReadAll(resp.Body)
if err != nil {
    log.Fatal(err)
}
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

if resp.StatusCode != http.StatusOK {
    log.Fatalf("unexpected status: %s", resp.Status)
}
```

POST request with JSON body:

```go
data := map[string]string{"name": "gopher"}
jsonData, err := json.Marshal(data)
if err != nil {
    log.Fatal(err)
}

resp, err := http.Post(
    "https://api.example.com/users",
    "application/json",
    bytes.NewBuffer(jsonData),
)
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

if resp.StatusCode != http.StatusCreated {
    log.Fatalf("unexpected status: %s", resp.Status)
}
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
        log.Printf("fetchData failed: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(data); err != nil {
        log.Printf("error encoding response: %v", err)
    }
}
```