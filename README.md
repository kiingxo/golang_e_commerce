# golang_e_commerce

An e-commerce REST API in Go, in early development.

## Current state
- `cmd/main.go`: builds the application config (listen address `:8080`, database DSN read from the environment).
- `cmd/api.go`: `Application` and `Config` types, and a [chi](https://github.com/go-chi/chi) router with request-logging middleware.

The HTTP server, database layer and product/order/cart endpoints aren't built yet.

## Stack
Go · chi v5

## Run
```bash
go run ./cmd
```

## Roadmap
- [ ] Start the HTTP server from `mount()`
- [ ] PostgreSQL connection and migrations
- [ ] Products, cart and orders endpoints
- [ ] Tests and CI
