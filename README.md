# Go Distributed Token Bucket Rate Limiter

[![Go CI](https://github.com/FelipeMadson/go-distributed-token-bucket/actions/workflows/ci.yml/badge.svg)](https://github.com/FelipeMadson/go-distributed-token-bucket/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Release](https://img.shields.io/badge/release-v1.0.0-green.svg)](https://github.com/FelipeMadson/go-distributed-token-bucket/releases)

Thread-safe, high-concurrency Token Bucket and Sliding Window rate limiting engine in Go. Guarantees zero data races across thousands of parallel goroutines.

## Features

- **Mutex Isolation**: Granular `sync.RWMutex` and `sync.Mutex` protection per tenant key.
- **Continuous Refill Calculus**: High-precision time delta evaluation eliminating discrete tick overhead.
- **Race Condition Free**: Rigorously validated under `go test -race`.

## Testing

```bash
go test -v -race ./...
```

## Author

Developed by **[Felipe Madison](https://github.com/FelipeMadson)**.
