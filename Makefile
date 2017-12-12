all: clean compile test

clean:
	@echo "Removing generated test files"
	rm -f coll/ringbuffer_string.go

compile:
	go build

test: test_generate
	@go test ./coll

test_generate:
	go generate coll/ringbuffer_test.go