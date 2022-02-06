# is.Maybe[T]: easy optional struct fields

Package is allows you to define struct fields which must differentiate between
their zero value and them not being assigned.

When building structs for JSON encoding/decoding, it's common to use pointers
for fields that are optional.

This pattern can make those structs annoying to use. For example:

```go
type User struct {
	Name            string
	AcceptedTermsAt *time.Time `json:",omitempty"`
}

u := User{
	Name: "Ernesto",
	AcceptedTermsAt: time.Now(),
}
// Will fail to compile:
// cannot use time.Now() (type time.Time) as type *time.Time in field value
```

This package provides a generic type `is.Maybe[T]` to make it easier to work with optional
fields.

Here is the same example leveraging `is.Maybe[T]`:

```go
type User struct {
	Name string
	AcceptedTermsAt *maybe.Is[time.Time]
}

u := User{
	Name: "Ernesto",
	AcceptedTermsAt: is.Value(time.Now()),
}

if t, ok := u.AcceptedTermsAt.ValueOk(); ok {
	fmt.Println(t)
}
```

## Types

### type [Maybe](/is.go#L9)

`type Maybe[T any] struct { ... }`

Maybe can hold a value of type T or nothing.

#### func [Nothing](/is.go#L23)

`func Nothing[T any]() *Maybe[T]`

Nothing returns a Maybe[T] that could hold a value of type T but is empty.

#### func [Value](/is.go#L15)

`func Value[T any](value T) *Maybe[T]`

Value returns a Maybe that wholds the given value.

#### func (*Maybe[T]) [HasValue](/is.go#L28)

`func (m *Maybe[T]) HasValue() bool`

HasValue returns false if Maybe is empty or true if it holds a value.

#### func (*Maybe[T]) [MarshalJSON](/is.go#L60)

`func (m *Maybe[T]) MarshalJSON() ([]byte, error)`

MarshalJSON marshals the underlying value. If m is empty, it will encode
null.

#### func (*Maybe[T]) [String](/is.go#L79)

`func (m *Maybe[T]) String() string`

String returns an empty string when empty. Otherwise, it
returns fmt.Sprint of the underlying value.

#### func (*Maybe[T]) [UnmarshalJSON](/is.go#L68)

`func (m *Maybe[T]) UnmarshalJSON(b []byte) error`

UnmarshalJSON unmarshals into the underlying value.

#### func (*Maybe[T]) [Value](/is.go#L42)

`func (m *Maybe[T]) Value() T`

Value held by the Maybe. If it's empty, it will return the zero value of
type T.

If m is nil, it will return the zero value of type T.

#### func (*Maybe[T]) [ValueOk](/is.go#L54)

`func (m *Maybe[T]) ValueOk() (T, bool)`

ValueOk is a shorthand to return (m.Value(), m.HasValue()).

## Examples

```golang
package main

import (
	"fmt"
	"github.com/gomodular/is"
	"time"
)

func main() {
	type User struct {
		Name            string
		AcceptedTermsAt *is.Maybe[time.Time]
	}

	u := User{
		Name:            "Ernesto",
		AcceptedTermsAt: is.Value(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)),
	}

	if t, ok := u.AcceptedTermsAt.ValueOk(); ok {
		fmt.Println(t)
	}

}

```

 Output:

```
2009-11-10 23:00:00 +0000 UTC
```

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
