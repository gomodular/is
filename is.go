package is

import (
	"encoding/json"
	"fmt"
)

// Maybe can hold a value of type T or nothing.
type Maybe[T any] struct {
	val T
	hasVal bool
}

// Value returns a Maybe that wholds the given value.
func Value[T any](value T) *Maybe[T] {
	return &Maybe[T]{
		val: value,
		hasVal: true,
	}
}

// Nothing returns a Maybe[T] that could hold a value of type T but is empty.
func Nothing[T any]() *Maybe[T] {
	return &Maybe[T]{hasVal: false}
}

// HasValue returns false if Maybe is empty or true if it holds a value.
func (m *Maybe[T]) HasValue() bool {
	if m == nil {
		return false
	}
	if !m.hasVal{
		return false
	}
	return true
}

// Value held by the Maybe. If it's empty, it will return the zero value of
// type T.
//
// If m is nil, it will return the zero value of type T.
func (m *Maybe[T]) Value() T {
	var nothing T
	if m == nil {
		return nothing
	}
	if !m.hasVal{
		return nothing
	}
	return m.val
}

// ValueOk is a shorthand to return (m.Value(), m.HasValue()).
func (m *Maybe[T]) ValueOk() (T, bool) {
	return m.Value(), m.HasValue()
}

// MarshalJSON marshals the underlying value. If m is empty, it will encode
// null.
func (m *Maybe[T]) MarshalJSON() ([]byte, error) {
	if !m.hasVal {
		return []byte("null"), nil
	}
	return json.Marshal(m.val)
}

// UnmarshalJSON unmarshals into the underlying value.
func (m *Maybe[T]) UnmarshalJSON(b []byte) error {
	m.hasVal = true
	return json.Unmarshal(b, &m.val)
}

type stringer interface {
	String() string
}

// String returns an empty string when empty. Otherwise, it
// returns fmt.Sprint of the underlying value.
func (m *Maybe[T]) String() string {
	if m == nil || !m.HasValue() {
		return ""
	}
	return fmt.Sprint(m.Value())
}
