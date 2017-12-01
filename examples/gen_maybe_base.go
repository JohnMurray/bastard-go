// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package main

// StringMaybe is a container for string that indicates a _possible_ value or nothing.
// This is an alternative structure to use over nil as it prevents unsafe access
// (sort-of) and prefers safe composition over direct dereferencing.
//
// See Map, FlatMap, and friends.
type StringMaybe struct {
	value string
	empty bool
}

// NewStringSome creates a new some-type for StringMaybe. Note that while you can store
// the zero-value into a some-type, it is not recommended.
func NewStringSome(value string) *StringMaybe {
	return &StringMaybe{
		value: value,
		empty: false,
	}
}

// NewStringNone creates a new none-type for StringMaybe. This type holds no value.
func NewStringNone() *StringMaybe {
	return &StringMaybe{
		empty: true,
	}
}

// IsSome returns bool is the StringMaybe is a some-type
func (m *StringMaybe) IsSome() bool {
	return !m.empty
}

// IsNone returns bool is the StringMaybe is a none-type
func (m *StringMaybe) IsNone() bool {
	return m.empty
}

// Get is an unsafe method to get the value of a some-type. If called on a none-type,
// it will panic.
func (m *StringMaybe) Get() string {
	if m.empty {
		panic("You do not know what you are doing")
	}
	return m.value
}

// GetOrElse is safe alternative to Get as it does not panic and will return a default
// value (passed into the function) if the current StringMaybe is a none-type.
func (m *StringMaybe) GetOrElse(elseValue string) string {
	if m.empty {
		return elseValue
	}
	return m.value
}

// With allows you to pass a function with a paramter of string as a way to execute over the some-type
// value only if it exists.
func (m *StringMaybe) With(f func(string)) {
	if !m.empty {
		f(m.value)
	}
}

// IntMaybe is a container for int that indicates a _possible_ value or nothing.
// This is an alternative structure to use over nil as it prevents unsafe access
// (sort-of) and prefers safe composition over direct dereferencing.
//
// See Map, FlatMap, and friends.
type IntMaybe struct {
	value int
	empty bool
}

// NewIntSome creates a new some-type for IntMaybe. Note that while you can store
// the zero-value into a some-type, it is not recommended.
func NewIntSome(value int) *IntMaybe {
	return &IntMaybe{
		value: value,
		empty: false,
	}
}

// NewIntNone creates a new none-type for IntMaybe. This type holds no value.
func NewIntNone() *IntMaybe {
	return &IntMaybe{
		empty: true,
	}
}

// IsSome returns bool is the IntMaybe is a some-type
func (m *IntMaybe) IsSome() bool {
	return !m.empty
}

// IsNone returns bool is the IntMaybe is a none-type
func (m *IntMaybe) IsNone() bool {
	return m.empty
}

// Get is an unsafe method to get the value of a some-type. If called on a none-type,
// it will panic.
func (m *IntMaybe) Get() int {
	if m.empty {
		panic("You do not know what you are doing")
	}
	return m.value
}

// GetOrElse is safe alternative to Get as it does not panic and will return a default
// value (passed into the function) if the current IntMaybe is a none-type.
func (m *IntMaybe) GetOrElse(elseValue int) int {
	if m.empty {
		return elseValue
	}
	return m.value
}

// With allows you to pass a function with a paramter of int as a way to execute over the some-type
// value only if it exists.
func (m *IntMaybe) With(f func(int)) {
	if !m.empty {
		f(m.value)
	}
}
