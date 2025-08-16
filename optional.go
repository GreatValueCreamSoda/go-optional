// Package go_optional provides a generic Optional type for handling values
// that may or may not be present. It allows safe access to values with methods
// to check presence and retrieve values, avoiding nil pointer issues. The
// package is designed to be type-safe and works with any type via Go generics.
package go_optional

import "errors"

var (
	// ErrGotValueOfNone is returned when MustValue is called on an Optional
	// with no value.
	ErrGotValueOfNone = errors.New("called MustValue() on None")
)

// Optional represents a value of type T that may or may not be present.
//
// If a value is present, IsSome returns true and the value can be retrieved
// with Value or MustValue. If no value is present, IsNone returns true.
type Optional[T any] struct {
	value       T
	isSomething bool
}

// Some creates an Optional containing the given value. The returned Optional
// satisfies IsSome.
func Some[T any](value T) Optional[T] { return Optional[T]{value, true} }

// None creates an Optional with no value.  The returned Optional satisfies
// IsNone.
func None[T any]() Optional[T] { return Optional[T]{isSomething: false} }

// IsSome reports whether the Optional contains a value.
func (o Optional[T]) IsSome() bool { return o.isSomething }

// IsNone reports whether the Optional does not contain a value.
func (o Optional[T]) IsNone() bool { return !o.isSomething }

// Value returns the contained value if present, otherwise the provided default
// Value.
//
// Example:
//
//	opt := Some(42)
//	fmt.Println(opt.Value(0)) // Output: 42
//	empty := None[int]()
//	fmt.Println(empty.Value(0)) // Output: 0
func (o Optional[T]) Value(defaultValue T) T {
	if o.isSomething {
		return o.value
	} else {
		return defaultValue
	}
}

// MustValue returns the contained value if present, otherwise it returns an
// error.
//
// If no value is present, it returns the zero value of type T and
// ErrGotValueOfNone.
func (o Optional[T]) MustValue() (T, error) {
	if !o.isSomething {
		var ret T
		return ret, ErrGotValueOfNone
	}
	return o.value, nil
}
