// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package main

// MapToInt is a right-bias mapping function and an alias for MapRightToInt
func (e *EitherStringOrString) MapToInt(f func(string) int) *EitherStringOrInt {
	if e.isLeft {
		return &EitherStringOrInt{
			left:   e.left,
			isLeft: true,
		}
	}
	return &EitherStringOrInt{
		right:  f(e.right),
		isLeft: false,
	}
}

// FlatMapToInt is a right-bias mapping function and an alias for FlatMapRightToInt
func (e *EitherStringOrString) FlatMapToInt(f func(string) *EitherStringOrInt) *EitherStringOrInt {
	if e.isLeft {
		return &EitherStringOrInt{
			left:   e.left,
			isLeft: true,
		}
	}
	return f(e.right)
}

// MapLeftToInt maps the left'ness of the either to a new either of type EitherIntOrString
func (e *EitherStringOrString) MapLeftToInt(f func(string) int) *EitherIntOrString {
	if e.isLeft {
		return &EitherIntOrString{
			left:   f(e.left),
			isLeft: true,
		}
	}
	return &EitherIntOrString{
		right:  e.right,
		isLeft: false,
	}
}

// MapRightToInt maps the right'ness of the either to a new either of type EitherIntOrString
func (e *EitherStringOrString) MapRightToInt(f func(string) int) *EitherStringOrInt {
	if e.isLeft {
		return &EitherStringOrInt{
			left:   e.left,
			isLeft: true,
		}
	}
	return &EitherStringOrInt{
		right:  f(e.right),
		isLeft: false,
	}
}

// FlatMapLeftToInt maps the left'ness of the either to a new Either of type EitherIntOrString
func (e *EitherStringOrString) FlatMapLeftToInt(f func(string) *EitherIntOrString) *EitherIntOrString {
	if e.isLeft {
		return f(e.left)
	}
	return &EitherIntOrString{
		right:  e.right,
		isLeft: false,
	}
}

// FlatMapRightToInt maps the right'ness of the either to a new Either of type EitherStringOrInt
func (e *EitherStringOrString) FlatMapRightToInt(f func(string) *EitherStringOrInt) *EitherStringOrInt {
	if e.isLeft {
		return &EitherStringOrInt{
			left:   e.left,
			isLeft: true,
		}
	}
	return f(e.right)
}

// MapToInt is a right-bias mapping function and an alias for MapRightToInt
func (e *EitherStringOrInt) MapToInt(f func(int) int) *EitherStringOrInt {
	if e.isLeft {
		return &EitherStringOrInt{
			left:   e.left,
			isLeft: true,
		}
	}
	return &EitherStringOrInt{
		right:  f(e.right),
		isLeft: false,
	}
}

// FlatMapToInt is a right-bias mapping function and an alias for FlatMapRightToInt
func (e *EitherStringOrInt) FlatMapToInt(f func(int) *EitherStringOrInt) *EitherStringOrInt {
	if e.isLeft {
		return &EitherStringOrInt{
			left:   e.left,
			isLeft: true,
		}
	}
	return f(e.right)
}

// MapLeftToInt maps the left'ness of the either to a new either of type EitherIntOrInt
func (e *EitherStringOrInt) MapLeftToInt(f func(string) int) *EitherIntOrInt {
	if e.isLeft {
		return &EitherIntOrInt{
			left:   f(e.left),
			isLeft: true,
		}
	}
	return &EitherIntOrInt{
		right:  e.right,
		isLeft: false,
	}
}

// MapRightToInt maps the right'ness of the either to a new either of type EitherIntOrInt
func (e *EitherStringOrInt) MapRightToInt(f func(int) int) *EitherStringOrInt {
	if e.isLeft {
		return &EitherStringOrInt{
			left:   e.left,
			isLeft: true,
		}
	}
	return &EitherStringOrInt{
		right:  f(e.right),
		isLeft: false,
	}
}

// FlatMapLeftToInt maps the left'ness of the either to a new Either of type EitherIntOrInt
func (e *EitherStringOrInt) FlatMapLeftToInt(f func(string) *EitherIntOrInt) *EitherIntOrInt {
	if e.isLeft {
		return f(e.left)
	}
	return &EitherIntOrInt{
		right:  e.right,
		isLeft: false,
	}
}

// FlatMapRightToInt maps the right'ness of the either to a new Either of type EitherStringOrInt
func (e *EitherStringOrInt) FlatMapRightToInt(f func(int) *EitherStringOrInt) *EitherStringOrInt {
	if e.isLeft {
		return &EitherStringOrInt{
			left:   e.left,
			isLeft: true,
		}
	}
	return f(e.right)
}

// MapToInt is a right-bias mapping function and an alias for MapRightToInt
func (e *EitherIntOrString) MapToInt(f func(string) int) *EitherIntOrInt {
	if e.isLeft {
		return &EitherIntOrInt{
			left:   e.left,
			isLeft: true,
		}
	}
	return &EitherIntOrInt{
		right:  f(e.right),
		isLeft: false,
	}
}

// FlatMapToInt is a right-bias mapping function and an alias for FlatMapRightToInt
func (e *EitherIntOrString) FlatMapToInt(f func(string) *EitherIntOrInt) *EitherIntOrInt {
	if e.isLeft {
		return &EitherIntOrInt{
			left:   e.left,
			isLeft: true,
		}
	}
	return f(e.right)
}

// MapLeftToInt maps the left'ness of the either to a new either of type EitherIntOrString
func (e *EitherIntOrString) MapLeftToInt(f func(int) int) *EitherIntOrString {
	if e.isLeft {
		return &EitherIntOrString{
			left:   f(e.left),
			isLeft: true,
		}
	}
	return &EitherIntOrString{
		right:  e.right,
		isLeft: false,
	}
}

// MapRightToInt maps the right'ness of the either to a new either of type EitherIntOrString
func (e *EitherIntOrString) MapRightToInt(f func(string) int) *EitherIntOrInt {
	if e.isLeft {
		return &EitherIntOrInt{
			left:   e.left,
			isLeft: true,
		}
	}
	return &EitherIntOrInt{
		right:  f(e.right),
		isLeft: false,
	}
}

// FlatMapLeftToInt maps the left'ness of the either to a new Either of type EitherIntOrString
func (e *EitherIntOrString) FlatMapLeftToInt(f func(int) *EitherIntOrString) *EitherIntOrString {
	if e.isLeft {
		return f(e.left)
	}
	return &EitherIntOrString{
		right:  e.right,
		isLeft: false,
	}
}

// FlatMapRightToInt maps the right'ness of the either to a new Either of type EitherIntOrInt
func (e *EitherIntOrString) FlatMapRightToInt(f func(string) *EitherIntOrInt) *EitherIntOrInt {
	if e.isLeft {
		return &EitherIntOrInt{
			left:   e.left,
			isLeft: true,
		}
	}
	return f(e.right)
}

// MapToInt is a right-bias mapping function and an alias for MapRightToInt
func (e *EitherIntOrInt) MapToInt(f func(int) int) *EitherIntOrInt {
	if e.isLeft {
		return &EitherIntOrInt{
			left:   e.left,
			isLeft: true,
		}
	}
	return &EitherIntOrInt{
		right:  f(e.right),
		isLeft: false,
	}
}

// FlatMapToInt is a right-bias mapping function and an alias for FlatMapRightToInt
func (e *EitherIntOrInt) FlatMapToInt(f func(int) *EitherIntOrInt) *EitherIntOrInt {
	if e.isLeft {
		return &EitherIntOrInt{
			left:   e.left,
			isLeft: true,
		}
	}
	return f(e.right)
}

// MapLeftToInt maps the left'ness of the either to a new either of type EitherIntOrInt
func (e *EitherIntOrInt) MapLeftToInt(f func(int) int) *EitherIntOrInt {
	if e.isLeft {
		return &EitherIntOrInt{
			left:   f(e.left),
			isLeft: true,
		}
	}
	return &EitherIntOrInt{
		right:  e.right,
		isLeft: false,
	}
}

// MapRightToInt maps the right'ness of the either to a new either of type EitherIntOrInt
func (e *EitherIntOrInt) MapRightToInt(f func(int) int) *EitherIntOrInt {
	if e.isLeft {
		return &EitherIntOrInt{
			left:   e.left,
			isLeft: true,
		}
	}
	return &EitherIntOrInt{
		right:  f(e.right),
		isLeft: false,
	}
}

// FlatMapLeftToInt maps the left'ness of the either to a new Either of type EitherIntOrInt
func (e *EitherIntOrInt) FlatMapLeftToInt(f func(int) *EitherIntOrInt) *EitherIntOrInt {
	if e.isLeft {
		return f(e.left)
	}
	return &EitherIntOrInt{
		right:  e.right,
		isLeft: false,
	}
}

// FlatMapRightToInt maps the right'ness of the either to a new Either of type EitherIntOrInt
func (e *EitherIntOrInt) FlatMapRightToInt(f func(int) *EitherIntOrInt) *EitherIntOrInt {
	if e.isLeft {
		return &EitherIntOrInt{
			left:   e.left,
			isLeft: true,
		}
	}
	return f(e.right)
}
