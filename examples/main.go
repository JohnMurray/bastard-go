package main

import (
	"fmt"
)

// The main functino just runs
func main() {

	examples := []struct {
		name string
		f    func()
	}{
		{
			name: "Maybe Example",
			f:    runMaybeExample,
		},
		{
			name: "Try Example",
			f:    runTryExample,
		},
	}

	for _, ex := range examples {
		fmt.Printf("Running %s\n", ex.name)
		ex.f()
		fmt.Println("==============================\n\n")
	}
}
