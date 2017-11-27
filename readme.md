# Bastard Go

A collection of terrible things in I felt like writing in Go.


### Try

This is a Scala-like `Try` object that uses `[genny][github_genny]` to generate templates which allow it to
work very similar to the Scala version with composition functions (`Map`, `FlatMap`, `Recover`, etc). To use
in your own projects:

```bash
go get github.com/cheekybits/genny
go install github.com/cheekybits/genny

go get github.com/johnmurray/bastard-go

cd $YOUR_PROJECT
mkdir try
```

```go
// Generate for as many types as you like. Simply including a comma-delimited list for `Type` will cause genny
// to generate multiple templates.
//
//go:generate genny -in=$GOPATH/src/github.com/johnmurray/bastard-go/try/try_base.go -out=try/try_base.go gen "Type=int,string,bool"

// Another command will generate the "composition" functions which allow you to convert from one type to another
// when using Map, FlatMap, etc. This is also necessary for generating same-type conversions. Such as:
//   IntTry.Map(func (i int) int { return i * 2 })
//
// Note that by providing multiple input types, all permutations will be generated.
//
//go:generate genny -in=$GOPATH/src/github.com/johnmurray/bastard-go/try/try_compose.go -out=try/try_compose.go gen "FromType=int,string,bool ToType=int,string,bool"
```

Of course you can use `genny` to generate it in a different package, so read more about `genny` if you're not familiar.


  [github_genny]: https://github.com/cheekybits/genny
  [johnmurray_io]: http://www.johnmurray.io/log/2017/11/27/Go-Try.html
