# Introducing Go, Chapter 1: Getting Started

## How to Read a Go Program

Apparently, top to bottom and left to right. More seriously, Doxsey makes several observations.

+ Every Go program must start with a package declaration. Go uses packages to organize and reuse code. Go distinguishes between executable and library packages. (Executable packages must begin with `package main`, I believe.)
+ The `import` keyword brings in code from other packages for us to use.
+ Go supports two types of comments. Single-line comments begin with `//`, and multi-line comments begin with `/*` and end with `*/`.
+ You define functions in Go as `func name(<parameters>) return { ... }`.
+ `main()` is a special function in Go. Go’s runtime knows to call the `main` function first when you execute a program.
+ You can use `go doc <package> <function name>` to get information about functions from other packages on the command line.
