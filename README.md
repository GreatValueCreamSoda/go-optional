# go-optional

[![Go Reference](https://pkg.go.dev/badge/github.com/GreatValueCreamSoda/go-optional.svg)](https://pkg.go.dev/github.com/GreatValueCreamSoda/go-optional)
[![Go Report Card](https://goreportcard.com/badge/github.com/GreatValueCreamSoda/go-optional)](https://goreportcard.com/report/github.com/GreatValueCreamSoda/go-optional)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

`go-optional` is a lightweight Go package that provides a generic `Optional[T]` type for handling values that may or may not be present. It offers a type-safe way to avoid nil pointer issues, with methods to check for value presence and retrieve values safely. Built with Go generics, it supports any type and follows idiomatic Go design principles.

## Features

- Generic `Optional[T]` type for any Go type.
- Methods to check value presence (`IsSome`, `IsNone`).
- Safe value retrieval with `Value` (using a default) or `MustValue` (error handling).

## Installation

To use `go-optional` in your Go project, ensure you have Go 1.18 or later (required for generics). Add the package to your project:

```bash
go get github.com/GreatValueCreamSoda/go-optional
```

## Usage

Optional types can be created using the `Some()` and `None()` containing and not containing a value respectively. Below are examples demonstrating common use of them and what functions optional types support.

```go
package main

import (
    "fmt"
    go_optional "github.com/GreatValueCreamSoda/go-optional"
)

func main() {
    // Create an Optional with a value
    some := go_optional.Some(42)
    if some.IsSome() {
        fmt.Println(some.Value(0)) // Output: 42
    }

    // Create an Optional with no value
    none := go_optional.None[string]()
    if none.IsNone() {
        fmt.Println(none.Value("default")) // Output: default
    }

    // Use MustValue for error handling
    someStr := go_optional.Some("hello")
    val, err := someStr.MustValue()
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Value:", val) // Output: Value: hello
    }

    // Handle None with MustValue
    noneInt := go_optional.None[int]()
    _, err = noneInt.MustValue()
    fmt.Println(err) // Output: called MustValue() on None
}
```

## Testing

The package includes comprehensive unit tests covering all methods and edge cases, including zero values and error handling. The following command is used to run all tests.

```bash
go test -v
```
