package main

import (
	"fmt"
)

//go:generate genny -in=$GOPATH/src/github.com/johnmurray/bastard-go/either/either_base.go.tmp -out=gen_either_base.go -pkg=main gen "TypeA=string,int TypeB=string,int"
//go:generate genny -in=$GOPATH/src/github.com/johnmurray/bastard-go/either/either_compose.go.tmp -out=gen_either_compose_1.go -pkg=main gen "FromTypeA=string,int FromTypeB=string,int ToTypeA=int ToTypeB=int"
//go:generate genny -in=$GOPATH/src/github.com/johnmurray/bastard-go/either/either_compose.go.tmp -out=gen_either_compose_2.go -pkg=main gen "FromTypeA=string,int FromTypeB=string,int ToTypeA=string ToTypeB=string"

func runEitherExample() {
	right := NewEitherIntOrStringRight("Hello")
	left := NewEitherIntOrStringLeft(5)

	hello := right.MapToString(func(s string) string {
		return (s + ", World!")
	})
	if hello.IsRight() {
		fmt.Println(hello.GetRightOrElse(""))
	}

	stillLeft := left.MapToString(func(s string) string {
		return s + ", World!"
	})
	if stillLeft.IsLeft() {
		fmt.Println("yep, still left")
	}
	stillLeftButNowAllStrings := stillLeft.MapLeftToString(func(i int) string {
		return "now  a string yo"
	})

	if stillLeftButNowAllStrings.IsLeft() {
		fmt.Println("all strings, but still a left")
	}

	left.FlatMapRightToString(func(s string) *EitherIntOrString {
		return right
	})
	if left.IsLeft() {
		fmt.Println("can't change my association, yo")
	}
}
