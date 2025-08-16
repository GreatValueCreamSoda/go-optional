// Package go_optional_test contains unit tests for the go_optional package.
package go_optional_test

import (
	"errors"
	"testing"

	go_optional "github.com/GreatValueCreamSoda/go-optional"
)

// TestSome verifies that Some creates an Optional with a value and satisfies IsSome.
func TestSome(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
	}{
		{"int", 42},
		{"string", "hello"},
		{"zero int", 0},
		{"empty string", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.value.(type) {
			case int:
				opt := go_optional.Some(v)
				if !opt.IsSome() {
					t.Error("Some(int) should return IsSome() true")
				}
				if opt.IsNone() {
					t.Error("Some(int) should return IsNone() false")
				}
				if got := opt.Value(0); got != v {
					t.Errorf("Some(%v).Value(0) = %v; want %v", v, got, v)
				}
			case string:
				opt := go_optional.Some(v)
				if !opt.IsSome() {
					t.Error("Some(string) should return IsSome() true")
				}
				if opt.IsNone() {
					t.Error("Some(string) should return IsNone() false")
				}
				if got := opt.Value(""); got != v {
					t.Errorf("Some(%q).Value(\"\") = %q; want %q", v, got, v)
				}
			}
		})
	}
}

// TestNone verifies that None creates an Optional with no value and satisfies IsNone.
func TestNone(t *testing.T) {
	tests := []struct {
		name       string
		makeNone   func() interface{}
		defaultVal interface{}
		wantValue  interface{}
	}{
		{"int", func() interface{} { return go_optional.None[int]() }, 0, 0},
		{"string", func() interface{} { return go_optional.None[string]() }, "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt := tt.makeNone()
			switch opt := opt.(type) {
			case go_optional.Optional[int]:
				if opt.IsSome() {
					t.Error("None[int]() should return IsSome() false")
				}
				if !opt.IsNone() {
					t.Error("None[int]() should return IsNone() true")
				}
				if got := opt.Value(tt.defaultVal.(int)); got != tt.wantValue {
					t.Errorf("None[int]().Value(%v) = %v; want %v", tt.defaultVal, got, tt.wantValue)
				}
			case go_optional.Optional[string]:
				if opt.IsSome() {
					t.Error("None[string]() should return IsSome() false")
				}
				if !opt.IsNone() {
					t.Error("None[string]() should return IsNone() true")
				}
				if got := opt.Value(tt.defaultVal.(string)); got != tt.wantValue {
					t.Errorf("None[string]().Value(%q) = %q; want %q", tt.defaultVal, got, tt.wantValue)
				}
			}
		})
	}
}

// TestMustValue verifies that MustValue returns the value for Some and an error for None.
func TestMustValue(t *testing.T) {
	tests := []struct {
		name      string
		opt       interface{}
		wantValue interface{}
		wantErr   error
	}{
		{"Some int", go_optional.Some(42), 42, nil},
		{"Some string", go_optional.Some("test"), "test", nil},
		{"None int", go_optional.None[int](), 0, go_optional.ErrGotValueOfNone},
		{"None string", go_optional.None[string](), "", go_optional.ErrGotValueOfNone},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch opt := tt.opt.(type) {
			case go_optional.Optional[int]:
				got, err := opt.MustValue()
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("MustValue() error = %v; want %v", err, tt.wantErr)
				}
				if err == nil && got != tt.wantValue {
					t.Errorf("MustValue() = %v; want %v", got, tt.wantValue)
				}
			case go_optional.Optional[string]:
				got, err := opt.MustValue()
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("MustValue() error = %v; want %v", err, tt.wantErr)
				}
				if err == nil && got != tt.wantValue {
					t.Errorf("MustValue() = %q; want %q", got, tt.wantValue)
				}
			}
		})
	}
}

// TestValue verifies that Value returns the contained value for Some or the default for None.
func TestValue(t *testing.T) {
	tests := []struct {
		name       string
		opt        interface{}
		defaultVal interface{}
		wantValue  interface{}
	}{
		{"Some int", go_optional.Some(42), 0, 42},
		{"None int", go_optional.None[int](), 0, 0},
		{"Some string", go_optional.Some("hello"), "default", "hello"},
		{"None string", go_optional.None[string](), "default", "default"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch opt := tt.opt.(type) {
			case go_optional.Optional[int]:
				if got := opt.Value(tt.defaultVal.(int)); got != tt.wantValue {
					t.Errorf("Value(%v) = %v; want %v", tt.defaultVal, got, tt.wantValue)
				}
			case go_optional.Optional[string]:
				if got := opt.Value(tt.defaultVal.(string)); got != tt.wantValue {
					t.Errorf("Value(%q) = %q; want %q", tt.defaultVal, got, tt.wantValue)
				}
			}
		})
	}
}
