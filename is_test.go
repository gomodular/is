package is_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/gomodular/is"
)

type testUser struct {
	ID string
}

type testStruct struct {
	Name string
	Pronouns *is.Maybe[string] `json:",omitempty"`
	IsActive *is.Maybe[bool]
	CreatedAt *is.Maybe[time.Time] `json:",omitempty"`
	User *is.Maybe[*testUser] `json:",omitempty"`
	Visits *is.Maybe[int]
}

func TestEncode(t *testing.T) {
	t.Run("emtpy", func(t *testing.T) {
		s := testStruct{
			Name: "Ernesto",
		}
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(s)
		require.NoError(t, err)
		require.Equal(t, `{"Name":"Ernesto","IsActive":null,"Visits":null}`, strings.TrimSpace(buf.String()))
	})

	t.Run("with values", func(t *testing.T) {
		s := testStruct{
			Name: "Ernesto",
			Pronouns: is.Value("he/him"),
			IsActive: is.Value(true),
			CreatedAt: is.Value(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)),
			User: is.Value(&testUser{ID: "123"}),
			Visits: is.Value(0),
		}
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(s)
		require.NoError(t, err)
		require.Equal(
			t,
			`{"Name":"Ernesto","Pronouns":"he/him","IsActive":true,"CreatedAt":"2009-11-10T23:00:00Z","User":{"ID":"123"},"Visits":0}`,
			strings.TrimSpace(buf.String()),
		)
	})
}

func TestDecode(t *testing.T) {
	t.Run("with values", func(t *testing.T) {
		var s testStruct
		err := json.Unmarshal(
			[]byte(`{"Name":"Ernesto","Pronouns":"he/him","IsActive":true,"CreatedAt":"2009-11-10T23:00:00Z","User":{"ID":"123"},"Visits":0}`),
			&s,
		)
		require.NoError(t, err)
		require.Equal(t, "Ernesto", s.Name)
		require.Equal(t, "he/him", s.Pronouns.Value())
		require.Equal(t, true, s.IsActive.Value())
		require.Equal(t, time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), s.CreatedAt.Value())
		require.Equal(t, "123", s.User.Value().ID)
		require.Equal(t, 0, s.Visits.Value())
	})

	t.Run("without values", func(t *testing.T) {
		var s testStruct
		err := json.Unmarshal(
			[]byte(`{"Name":"Ernesto"}`),
			&s,
		)
		require.NoError(t, err)
		require.Equal(t, "Ernesto", s.Name)
		require.Equal(t, "", s.Pronouns.Value())
		require.False(t, s.Pronouns.HasValue())
		require.Equal(t, false, s.IsActive.Value())
		require.False(t, s.IsActive.HasValue())
		require.Equal(t, time.Time{}, s.CreatedAt.Value())
		require.False(t, s.CreatedAt.HasValue())
		require.Equal(t, 0, s.Visits.Value())
		require.False(t, s.Visits.HasValue())
	})
}

func TestMaybe_Value(t *testing.T) {
	t.Run("nil *is.Maybe[string] returns empty string", func(t *testing.T) {
		var v *is.Maybe[string]
		require.Equal(t, "", v.Value())
	})

	t.Run("*is.Maybe[string] with value returns the value", func(t *testing.T) {
		v := is.Value("something")
		require.Equal(t, "something", v.Value())
	})

	t.Run("empty *is.Maybe[string] returns zero value", func(t *testing.T) {
		v := is.Nothing[string]()
		require.Equal(t, "", v.Value())
	})
}

func ExampleMaybe_Value() {
	var (
		nilValue *is.Maybe[string]
		empty = is.Nothing[string]()
		withValue = is.Value("something")
	)
	fmt.Printf("nilValue: %#v\n", nilValue.Value())
	fmt.Printf("empty: %#v\n", empty.Value())
	fmt.Printf("withValue: %#v\n", withValue.Value())
	// Output:
	// nilValue: ""
	// empty: ""
	// withValue: "something"
}

func TestMaybe_HasValue(t *testing.T) {
	t.Run("nil *is.Maybe[string] returns false", func(t *testing.T) {
		var v *is.Maybe[string]
		require.Equal(t, false, v.HasValue())
	})

	t.Run("*is.Maybe[string] with value returns true", func(t *testing.T) {
		v := is.Value("something")
		require.Equal(t, true, v.HasValue())
	})

	t.Run("empty *is.Maybe[string] returns false", func(t *testing.T) {
		v := is.Nothing[string]()
		require.Equal(t, false, v.HasValue())
	})
}

func TestMaybe_ValueOk(t *testing.T) {
	t.Run("nil *is.Maybe[string] returns zero value, false", func(t *testing.T) {
		var v *is.Maybe[string]
		val, ok := v.ValueOk()
		require.Equal(t, "", val)
		require.Equal(t, false, ok)
	})

	t.Run("*is.Maybe[string] with value returns true", func(t *testing.T) {
		v := is.Value("something")
		val, ok := v.ValueOk()
		require.Equal(t, "something", val)
		require.Equal(t, true, ok)
	})

	t.Run("empty *is.Maybe[string] returns false", func(t *testing.T) {
		v := is.Nothing[string]()
		val, ok := v.ValueOk()
		require.Equal(t, "", val)
		require.Equal(t, false, ok)
	})
}

func ExampleMaybe_HasValue() {
	var (
		nilValue *is.Maybe[string]
		empty = is.Nothing[string]()
		withValue = is.Value("something")
	)
	fmt.Printf("nilValue: %#v\n", nilValue.HasValue())
	fmt.Printf("empty: %#v\n", empty.HasValue())
	fmt.Printf("withValue: %#v\n", withValue.HasValue())
	// Output:
	// nilValue: false
	// empty: false
	// withValue: true
}

func TestMaybe_String(t *testing.T) {
	t.Run("nil *is.Maybe[testStruct] returns false", func(t *testing.T) {
		var v *is.Maybe[testStruct]
		require.Equal(t, "", v.String())
	})

	t.Run("*is.Maybe[testStruct] with value returns true", func(t *testing.T) {
		v := is.Value(testStruct{Name: "Ernesto"})
		require.Equal(t, "{Ernesto     }", v.String())
	})

	t.Run("empty *is.Maybe[testStruct] returns false", func(t *testing.T) {
		v := is.Nothing[string]()
		require.Equal(t, "", v.String())
	})
}

func Example() {
	type User struct {
		Name string
		AcceptedTermsAt *is.Maybe[time.Time]
	}

	u := User{
		Name: "Ernesto",
		AcceptedTermsAt: is.Value(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)),
	}

	if t, ok := u.AcceptedTermsAt.ValueOk(); ok {
		fmt.Println(t)
	}

	// Output:
	// 2009-11-10 23:00:00 +0000 UTC
}
