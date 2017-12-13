package future

// Promise is the writable side to a Future whereas the Future is read-only.
// Similarly a Promise is the representation of a future-computed result.
type Promise struct {
	result  interface{}
	written bool
}

func (p *Promise) write(value interface{}) {
	if !p.written {
		p.result = value
		p.written = true
	}
}

func (p *Promise) get() interface{} {
	if !p.written {
		panic("you do not know what you are doing") // (╯°□°)╯︵ ┻━┻
	}
	return p.result
}

// Future is the representation of a future-computed result. Future is a read-only
// interface. To understand how values are written to a Future, look at Promise.
type Future struct {
	promise *Promise
}

// IsReady determines if a futures value has been computed
func (f *Future) IsReady() bool {
	return f.promise.written
}

// Get fetches the value of a future
func (f *Future) Get() interface{} {
	return f.promise.result
}
