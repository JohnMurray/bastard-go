package main

import (
	"fmt"
)

//go:generate genny -in=$GOPATH/src/github.com/johnmurray/bastard-go/maybe/maybe_base.go.tmp -out=gen_maybe_base.go -pkg=main gen "Type=string,int"
//go:generate genny -in=$GOPATH/src/github.com/johnmurray/bastard-go/maybe/maybe_compose.go.tmp -out=gen_maybe_compose.go -pkg=main gen "FromType=string,int ToType=string,int"
func runMaybeExample() {
	fmt.Println("hello")

	a := NewStringSome("Hello")
	b := NewIntNone()
	d := 7

	c := NewIntSome(a.FlatMapToInt(func(s string) *IntMaybe {
		return b
	}).GetOrElse(d))

	if c.IsSome() {
		fmt.Printf("Got some: %v\n", c.Get())
	} else {
		fmt.Println("Got none")
	}
}
