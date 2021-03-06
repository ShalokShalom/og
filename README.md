Og-Lang (Optimistic Golang)
===
### *v0.6.1* ([CHANGELOG](https://github.com/champii/og/tree/master/CHANGELOG.md))

<table>
  <tr><td><b>"Golang On Steroids"</b></td>         <td>- <em>Socrates</em></td></tr>
  <tr><td><b>"The Code like it should be. 5/7"</b></td><td>- <em>Mahatma Gandhi</em></td></tr>
  <tr><td><b>"(..Recursive Facepalm..)"</b></td> <td>- <em>Google</em></td></tr>
</table>

# Index

1. [Intro](#intro)
1. [Features](#features)
1. [Install](#install)
1. [Quick Overview](#quick-overview)
1. [Usage](#usage)
    1. [Interpreter (ALPHA)](#interpreter-alpha)
    2. [Basic file compilation](#basic-file-compilation)
    3. [Debug](#debug)
1. [Build](#build)
1. [Todo](#todo)
1. [Long term utopia](#long-term-utopia)

# Intro

<h3> Disclamer: This software is in its early stage.<br/>
New features come fast, at the cost of breaking things often.<br />
Testers and Contributors are most welcome
</h3>


`Og` is to be pronounced `Oh-Jee` and stands for ~~`Orgasmic Granny`~~ `Optimistic Golang`

`Oglang` is an indentation based language mainly inspired from [Livescript](http://livescript.net) that compiles to a subset of `GoLang`.


### Bootstraped Language

`Oglang` is written in itself. It is said to be a 'Bootstraped' language. In fact, `Oglang` needs the previous release of itself to build itself.

See the [Src](https://github.com/champii/og/tree/master/src) folder for both `Oglang` and `Golang` sources.

Built with [Antlr4](https://github.com/antlr/antlr4) from their `Golang` grammar.

![Hello world](https://github.com/Champii/og/blob/master/docs/hello_preview.gif)

### Goal

The main goal is to simplify the syntax, to borrow some concepts from Livescript and other functional languages, to implement Generics and macro processing, as well as some syntaxic sugar to avoid all the boilerplate code Golang forces us into.

# Features
- Indentation based
- Transpile to `Golang`
- No brackets (almost)
- Clean and nice syntax
- Auto return the last statement in a block
- Returnable/Assignable statements (`if` only for the moment)
- Inline method declaration in `struct`
- One-liners
- Syntaxic sugars
- Bootstraped language
- Antlr4 parser
- Multithreaded compilation
- CLI tool to parse and debug your files
- Generics (ALPHA)
- Interpreter (ALPHA)

# Install

```bash
# You just have to `go get` the repo
go get -u github.com/champii/og

# If your `$PATH` includes `$GOPATH/bin` (and it should)
og --version # or `og -V`
```

# Quick Overview

### [Full overview here](https://github.com/champii/og/tree/master/docs/overview.md) with compiled comparison

This is an quick overview of how `Oglang` looks like actualy. 

See the [Exemples](https://github.com/champii/og/tree/master/tests/exemples) folder or the [Src](https://github.com/champii/og/tree/master/src) folder for more exemples.

```go
!main

import
  fmt
  strings
  "some/repo"

struct Foo
  bar int
  getBar: int    -> @bar
  *setBar(i int) -> @bar = i

Foo::inc    (foo int): int -> @bar + foo
Foo::*setInc(foo int)      -> @bar = @bar + foo

interface Bar
  Foo()
  Bar(i int): SomeType

otherFunc(a string): string -> a

autoIfReturn: int ->
  if true => 1
  else    => 0

genericFunction<T>(arg T): T -> arg

main ->
  test := Foo{}
  test.inc(42)

  genericFunction<int>(42)

  var a int = 
    if true => 1
    else    => 0

  someArr := []string
    "value1"
    "value2"

  for i, v in someArray
    fmt.Println(i, v)

  switch test.getBar()
    42 => doSomething()
    _  => doDefault()

  callback := fn (arg int): int -> arg + 1

  go someFunc()
  go -> doSomething()
```

# Usage

![Og](https://github.com/Champii/og/blob/master/docs/og_preview.gif)


```
NAME:
  Oglang - Golang on steroids

VERSION:
  v0.6.1

USAGE:
  og [options] [folders...|files...]

  If a Print argument (-p, -b, -d, -a) is given, NO COMPILATION is done.

  If run without files, it will compile and execute '.'

OPTIONS:
  -r, --run                      Run the binary
  -o directory, --out directory  Output directory. If input is recursive folder, the tree is recreated (default: "./")
  -w jobs, --workers jobs        Set the number of jobs (default: 8)
  -p, --print                    Print the file
  -d, --dirty                    Print the file before going through 'go fmt'
  -b, --blocks                   Print the file after it goes to preprocessor. Shows only block-based indentation
  -a, --ast                      Print the generated AST
  -i, --interpreter              Run a small interpreter (ALPHA)
  -q, --quiet                    Hide the progress output
  -n, --no-build                 Dont run 'go build'
  -h, --help                     Print help
  -v, --version                  Print version
```

### Interpreter (ALPHA)

Og embed a small interpreter that in fact compiles the given string into a `/tmp/main.go` skelton and run it. A better implementation will come.

```bash
./og -i
> 1+1
2
```

### Basic file compilation

By default `Og` compile every `.og` file in the current folder `.` and produce `.go` files that are along their `Og` source. It will then run `go build` on the folder and try to run the created binary
```bash
./og
```

With just a file name, the compiler will produce a `.go` file inside the same directory
```bash
./og file.og
```

You can give multiple files and folder that will be walked recursively
```bash
./og file.og folder/ anotherFile.og
```

The output flag `-o` will save the files into another folder. The folder hierarchy is recreated. 
```bash
./og -o lib src/file.og
```

### Debug

You can also print the file without affecting the fs with `-p`
```bash
./og -p src/file.og
```

The `-d` (`--dirty`) option shows you the bare generated file from the parser, without formating with `go fmt`. This is useful to check if the generated syntax is valid.

The `-b` (`--block`) option prints the output of the preprocessor who's in charge to create the blocks from indentation. No compilation is done.

The `-a` (`--ast`) option prints the generated AST from the parser

# Build

The current build time of the project is around 5s for all sources files with `./og` alone, and around 15s for full rebootstrap with `make re` (That bootstraps from old version then rebootstraps from itself, with `go build` and `go test` each time). 

Here is the procedure to regenerate the parser from the grammar if you want to make changes to it.

If you just want to (re)build the binary, you can call `make build` or just `go build` (needs a previously generated parser from grammar. See below)

You will need `Java`, the Antlr4 library is in `./parser/antlr4-4.7.1-SNAPSHOT-complete.jar`

```bash
# Get Og
go get -u github.com/champii/og
cd $GOPATH/src/github.com/champii/og

# This will regenerate the grammar,
# Compile the existing sources from the previous Og,
# And run the tests.
# Needs the last official `og` binary version at global scope.
make

# It cleans the `lib` folder,
# Then compiles og from the previous global version
# Then recomiles it from itself
# And run the tests
make re

# Simple exemple
og exemples/import.og
```

# TODO

## Golang basics
- [ ] Ternary expression
- [ ] Named return
- [ ] Fix type switch assignement `switch t := v.(type)` 

## Syntaxic Sugar
- [ ] Slices declaration without type `[1, 2, 3]`, `[]string` (need type inference)
- [ ] `*` Operator in slices to reference own lenght `arr = arr[*-1]`
- [ ] Suffix keyword `return foo if bar`, `foo += i for i in array`
- [ ] Returnable and assignable statements (for, switch, ...)
- [ ] Predicat recapture: `if a => that`
- [ ] External type declaration like Haskell: `myFunc :: string -> Foo -> Bar`
- [ ] Existance test (if foo? => bar) for non-nil value test
- [ ] `pub` visibility instead of capitalizing
- [ ] For with a range (for i in [0..10])
- [ ] Pattern matching
- [ ] Auto setup package name with folder name if not specified
- [ ] Error bubbling
- [ ] Function currying
- [ ] Function shorthand `(+ 10)`
- [ ] Generics
- [ ] Import pattern matching
- [ ] Remove that `fn` keywork that diminish lisibility
- [ ] Conditionnal expression like `res := getFirst() || getSecond()` that make `if` statements
- [ ] Assignation and return for `for`, `switch`
- [ ] `super` keyword

## Technical
- [ ] Perfs !! (More specific rules, reduce size and workload of Walkers, remove ambiguity in grammar)
- [ ] Do a single pass on AST instead of multiple walkers (for perfs)
- [ ] Fix bad perfs for nested struct instantiation 
- [ ] Simple type checker to catch errors before Golang formater/compiler does
- [ ] Don't recompile files that didn't change
- [ ] Language extensions ?
- [ ] Make tests truly executable
- [ ] VSCode extension
- [ ] Adapt Golang tooling like `gofmt` or `golint`
- [ ] Better error context (How to keep line number after preproc ?)

# Long term utopia

What we want `Og` to looks like in the futur

[Utopia](https://github.com/champii/og/tree/master/tests/exemples/utopia.og.exemple)
