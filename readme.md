# Bastard Go

A collection of terrible things in I felt like writing in Go

Most of the things in this repository use [`genny`][github_genny] to generate code, so to really play with
any of this horrible stuff, you should run

```bash
go get github.com/cheekybits/genny
go install github.com/cheekybits/genny

go get github.com/johnmurray/bastard-go
```

For more details on Genny and how the generation works and options to use with the command, head on over
to their GitHub page. For examples, see the `examples` folder.

### TOC

  * [__Templates__](#templates)
    * [__Either__](#either)
    * [__Maybe__](#maybe)
    * [__Try__](#try)
  * [__Contributing__](#contributing)
  * [__License__](#license)


## Templates

### Either

This is a right-bias `Either` type that uses [`genny`][github_genny] to generate templates which allow it to work
very vimilar to a version using generics. However, this is probably the clearest example (in this project) where
generics would greatly simplify the boilerplate of code-generation as it is a multi-type container that has the
ability to map to other `Either` types with different underlying "generic" types. Creating a many-to-many mapping
that makes generation tricky and cumbersome.

To use in your own projects:

```bash
cd $YOUR_PROJECT
mkdir either
```

```go
// Generate for as many types as you like. Simply including a comma-delimited list for
// `TypeA` and `TypeB` will cause genny to generate multiple templates. Note that providing
// multiple for both inputs will create every possible combination.
//
//go:generate genny -in=$GOPATH/src/github.com/johnmurray/bastard-go/either/either_base.go -out=either/either_base.go gen "TypeA=int,string,bool TypeB=int,string,bool"

// Another command will generate the "composition" functions which allow you to convert
// from one type to another // when using Map, FlatMap, etc. This is also necessary for
// generating same-type conversions. Such as:
//   eitherStringOrInt.MapRightToInt(func (i int) int { return i * 2 })
//
// Note that to avoid generating duplicate functions (due to so many input types and overlapping
// righ/left combinations) the `ToTypeA` and `ToTypeB` should be done one at a time, using the same
// value for each:
//
//go:generate genny -in=$GOPATH/src/github.com/johnmurray/bastard-go/either/either_compose_1.go -out=either/either_compose.go gen "FromTypeA=int,string,bool FromTypeB=int,string,bool ToTypeA=int ToTypeB=int"
//go:generate genny -in=$GOPATH/src/github.com/johnmurray/bastard-go/either/either_compose_2.go -out=either/either_compose.go gen "FromTypeA=int,string,bool FromTypeB=int,string,bool ToTypeA=string ToTypeB=string"
//go:generate genny -in=$GOPATH/src/github.com/johnmurray/bastard-go/either/either_compose_3.go -out=either/either_compose.go gen "FromTypeA=int,string,bool FromTypeB=int,string,bool ToTypeA=bool ToTypeB=bool"
```

### Maybe

This is a Scala/Haskell like `Some`/`Maybe` object that uses [`genny`][github_genny] to generate templates which
allow it to work very similar to a version using generics. To use in your own projects:

```bash
cd $YOUR_PROJECT
mkdir maybe
```

```go
// Generate for as many types as you like. Simply including a comma-delimited list for
// `Type` will cause genny to generate multiple templates.
//
//go:generate genny -in=$GOPATH/src/github.com/johnmurray/bastard-go/maybe/maybe_base.go -out=maybe/maybe_base.go gen "Type=int,string,bool"

// Another command will generate the "composition" functions which allow you to convert
// from one type to another // when using Map, FlatMap, etc. This is also necessary for
// generating same-type conversions. Such as:
//   IntMaybe.MapToInt(func (i int) int { return i * 2 })
//
// Note that by providing multiple input types, all permutations will be generated.
//
//go:generate genny -in=$GOPATH/src/github.com/johnmurray/bastard-go/maybe/maybe_compose.go -out=maybe/maybe_compose.go gen "FromType=int,string,bool ToType=int,string,bool"
```


### Try

This is a Scala-like `Try` object that uses [`genny`][github_genny] to generate templates which allow it to
work very similar to the Scala version with composition functions (`Map`, `FlatMap`, `Recover`, etc). To use
in your own projects:

```bash
cd $YOUR_PROJECT
mkdir try
```

```go
// Generate for as many types as you like. Simply including a comma-delimited list for
// `Type` will cause genny to generate multiple templates.
//
//go:generate genny -in=$GOPATH/src/github.com/johnmurray/bastard-go/try/try_base.go -out=try/try_base.go gen "Type=int,string,bool"

// Another command will generate the "composition" functions which allow you to convert
// from one type to another // when using Map, FlatMap, etc. This is also necessary for
// generating same-type conversions. Such as:
//   IntTry.MapToTry(func (i int) int { return i * 2 })
//
// Note that by providing multiple input types, all permutations will be generated.
//
//go:generate genny -in=$GOPATH/src/github.com/johnmurray/bastard-go/try/try_compose.go -out=try/try_compose.go gen "FromType=int,string,bool ToType=int,string,bool"
```

Of course you can use `genny` to generate it in a different package, so read more about `genny` if you're not familiar.
See my [blog post][johnmurray_io] for more details.


## Contributing

Have you'd like to implement that you think belongs in this repository? Feel free to open a PR to talk about
your idea, or to present some code you think would fit in. The more the merrier! :-D

## License

Project is licensed under MIT license. Please see `LICENSE` file for full details.


  [github_genny]: https://github.com/cheekybits/genny
  [johnmurray_io]: http://www.johnmurray.io/log/2017/11/27/Go-Try.html
