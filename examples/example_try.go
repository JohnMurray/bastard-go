package main

import (
	"fmt"
	"strconv"
)

//go:generate genny -in=$GOPATH/src/github.com/johnmurray/bastard-go/try/try_base.go.tmp -out=gen_try_base.go -pkg=main gen "Type=string,int"
//go:generate genny -in=$GOPATH/src/github.com/johnmurray/bastard-go/try/try_compose.go.tmp -out=gen_try_compose.go -pkg=main gen "FromType=string,int ToType=string,int"
func runTryExample() {
	fmt.Println("hello")
	composedValue := NewStringSuccess("37").
		FlatMapWrapToInt(strconv.Atoi).
		MapToString(func(i int) string {
			return fmt.Sprintf("Parsed to: %d", i)
		})

	if composedValue.IsSuccess() {
		fmt.Println(composedValue.GetValue())
	}
}
