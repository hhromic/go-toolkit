# go-toolkit

[![Go Reference](https://pkg.go.dev/badge/github.com/hhromic/go-toolkit.svg)](https://pkg.go.dev/github.com/hhromic/go-toolkit)

Simple toolkit with reusable utilities for Go programs.

## Testing

To run package testing with code coverage profiling and reporting:
```
go test -coverprofile cover.out ./...
go tool cover -html=cover.out -o cover.html
```

## License

This project is licensed under the [Apache License Version 2.0](LICENSE).
