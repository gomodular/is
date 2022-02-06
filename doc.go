/*
Package is allows you to define struct fields which must differentiate between
their zero value and them not being assigned.

When building structs for JSON encoding/decoding, it's common to use pointers
for fields that are optional.

This pattern can make those structs annoying to use. For example:

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

This package provides a generic type `is.Maybe[T]` to make it easier to work with optional
fields.

Here is the same example leveraging `is.Maybe[T]`:

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
*/
package is

//go:generate ./scripts/generate_readme.sh
