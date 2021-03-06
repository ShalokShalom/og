# Changelog

## v0.6.1: Current version

## v0.6.0
  - Generics that are produced at compile time
    ```go
      genericFunction<T>(arg T): T -> arg

      main ->
        genericFunction<int>(a)
        genericFunction<string>(a)
    ```
  - Cleaner Makefile
  - Awesome cli tool to use multi-threaded compilation 
  - Multithreaded tests
  - Overall project compilation time down to 8s, tests down to 2s.
  - Beautiful progress, made verbose default
  - Removed option `-v, --verbose` to be the opposite: `-q, --quiet`
  - Add option `-w, --workers` to specify a number of thread for compilation
  - Changed `-V, --version` for little `-v, --version`
  - Auto build after compile, add option `-n, --no-build` to avoid that behaviour
  - Add option `-r, --run` to run after compile
  - Overview improvements
  - Allow comments starting with `#` 
  - First type-checker implementation
  - Better `AstWalker`
  - Colors everywhere !


## v0.5.0
  - Removed every non necessary terminal `return`
  - Makefile now compiles with the previous released version by default instead of the last compiled one. New rule `make new` that compiles itself with the last present `og`
  - Fix a bug on ParameterDecl that wanted a name for each type
  - Auto return for `if` statement
    ```go
      someFunc: int ->
        if true => 42
        else    => 0
    ```
  - Assignation for `if` statement
    ```go
      var a int = 
        if true => 42
        else    => 0
    ```
  - Overview section in the docs

## v0.4.2
  - Auto return for function with no return keyword (only for simple statement yet)
    ```go
      someFunc: int -> 2
    ```

## v0.4.1
  - Abstract AST walker and a derived AST Printer
  - `-a` CLI option to show the AST
  - Disambiguation of `for`, `if` and `else if` statements with a `;` before the block
  - Empty statement and empty onliner-function body with `;`
    ```go
      foobar -> ;

      if foo == bar => ;
    ```

## v0.4.0
  - New AST system for more flexibility
  - Switched to SLL prediction mode to greatly improve performances
  - Multiline `type`, `const` and `var` declaration
    ```go
      var (
        foo string
        bar = 2
      )
    ```
  - Disambiguation character `~` for `label` statements
    ```go
      ~myLabel: something()
    ```

## v0.3.0
  - Allow single line If statement
    ```go
      if b == 0 => 0
      else      => 1
    ```
  - Pretty error display (but trivial for some cases)
    ```
      path/file.og (8:6): Unexpected 'foo'
        bar foo
            ^
    ```
  - Allow empty Return statement
  - Small and dirty interpreter that allow only a few expressions

## v0.2.0
  - Variadic arguments
    ```go
      someFunc(a ...int) -> fmt.Println(a)
      main -> someFunc([]int{1, 2, 3}...)
    ```
  - Parenthesis allowed for import declaration
    ```go
      import (
        fmt
        strings
      )
    ```
  - For statement with 3 flavors
    ```go
      for _, i in a
      for i < 10
      for i := 0; i < 10; i++
    ```
  - Bitwise operators `&`, `|`, `^`, `&^`, `<<`, `>>`
  - Import renaming
    ```go
      import
        fmt
        strings
        "some/repo": otherName
        "another/repo": .
    ```

## v0.1.11
  - Function literal but with a mandatory desambiguation symbol `fn`
    ```go
      a := fn              -> fmt.Println(1)
      b := fn (a int)      -> fmt.Println(a)
      c := fn (a int) :int -> return a
      a := fn         :int -> return 1
    ```
  - Function Type with the same mandatory `fn`
    ```go
      a(arg fn: int) -> fmt.Println(arg())
      main          -> a(fn: int -> return 2)
    ```

## v0.1.10
  - Better `Makefile` that cares about whether a file has changed or not before applying rules
  - Proper Anonymous field in structs
    ```go
      struct Foo
        *SomeClass
        a int
    ```
  - Channel Send and Receive
    ```go
      c := make(chan int)
      c <- 1
      x := <-c
    ```
  - Select statement
    ```go
      select
        <-c1      => something()
        x := <-c1 => something()
        c1 <- 1   =>
          something()
          somethingElse()
        _         => defaultCase()
    ```
  - Pointer type for receiver in methods
    ```go
      // inline
      struct Foo
        nonPointer       : int -> return 1
        *pointerReceiver : int -> return 1

      // external
      Foo::nonPointer2       : int -> return 1
      Foo::*pointerReceiver2 : int -> return 1
    ```
  - External CHANGELOG
  - Function literal (Big performance issue caused by the grammar that keeps increasing in complexity, disabled until it's sorted out)
    ```go
      a := -> fmt.Println(1)
      a := (a int) -> fmt.Println(a)
    ```

## v0.1.9
  - `break`, `goto`, `continue`, `fallthrough`, `labels`

## v0.1.8
  - Proper `const` declaration
  - Proper `var` declaration
  - Map declaration
  - Chan type (`chan T`, `<-chan T` and `chan<- T`)
  - Fixed a bug with result type that were not being parsed. We can now use `interface` instread of `interface{}` everywhere

## v0.1.7
  - Add custom antlr4 directly into sources. No need to build it yourself

## v0.1.6
  - Release system for develop

## v0.1.5
  - Forced `Go` syntax highlight on `Og` files for `Github`
  - Rework translator to adapt to new method syntax and `@` alias
  - No need to specify the `{}` in `interface{}` for arguments types (not for return values yet)
    ```go
      fn(a interface): interface{} -> a
    ```

## v0.1.4
  - Class-like method declaration (nested into the struct)
    ```go
      struct Foo
        bar int
        f : int -> return @bar
    ```

## v0.1.3
  - Slice manipulation
      ```go
      someArr[1:x]
      ```
  - Interfaces
      ```go
      interface Foo
        Fn(a ArgType): ResType
      ```
  - Alias `@` => `this`
    ```go
      Foo::bar : SomeType -> return @someProp
    ```
## v0.1.2
  - Support for simple `struct ID {}` declaration. Still support `type ID struct {}`.
    ```go
      // Equivalent
      struct Foo {}
      type Foo struct {}
    ```
  - Alias `class` => `struct`.
    ```go
      // Equivalent
      struct Foo {}
      class Foo {}
    ```
  - Allow empty `class` declaration
    ```go
      struct Foo
    ```
  - Shorthand for `package main` => `!main`.
    ```go
      // Equivalent
      !main
      package main
    ```
  - Cli support with options and recursive directory walking
    ```bash
      og -o lib src
    ```
  - External method declaration
    ```go
      struct Foo
      Foo::bar -> doSomething()
    ```

## v0.1.0
  - Initial release
