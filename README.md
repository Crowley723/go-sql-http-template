# Go HTTP Server Template

A minimal Go HTTP server template with structured logging, graceful shutdown, and a custom middleware system for rapid project initialization.

## Features

- Clean middleware architecture with dependency injection
- Structured logging using Go's `slog` package
- Graceful shutdown with signal handling
- Hot reload development workflow
- Built-in JSON and text response utilities
- Debug support with Delve debugger

## Quick Start

```bash
# Install dependencies
make install

# Start development server with hot reload
make dev

# Server runs at http://localhost:8080
curl http://localhost:8080/health
```

## Development

```bash
make dev        # Hot reload development
make debug      # Debug mode with delve on :2345
make test       # Run tests
make coverage   # Run tests with coverage
```

## Architecture

The template uses a custom `AppContext` middleware system that provides:

- Request/response access
- Structured logger
- JSON/text response utilities
- Error handling helpers

Add new routes in `api/server.go`:
```go
mux.HandleFunc("GET /api/users", middlewares.Wrap(handleUsersGET))
```

Create handlers in `api/handlers.go`:
```go
func handleUsersGET(ctx *middlewares.AppContext) {
    ctx.WriteJSON(http.StatusOK, map[string]string{"users": "data"})
}
```

## Use as Template

This template provides a solid foundation for HTTP services requiring authentication, APIs, or web applications. The middleware system makes it easy to add JWT auth, CORS, rate limiting, and other common HTTP service features.