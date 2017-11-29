# Examples

This folder contains examples of generating and using code from the bastard-go
project. See `main.go` for all the examples that are run. All generated files are
prefixed with `gen_` to make it clear. Similarly, all example files are prefixed
with `example_`.

To run all:
```bash
$ go build .

$ ./examples
```

To re-generate generated files, use `go generate` (note `genny` is required to be present).
Example:

```bash
$ go generate example_try.go
```