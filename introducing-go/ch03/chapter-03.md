# Introducing Go, Chapter 3: Variables

A variable is a name for a storage location that may contain a value. Go considers variables to be of specific types. You can create variables in several ways in Go.

+ `var name type = value`
+ `var name type` and then later `name = value`
+ `var name = value` (Go’s compiler infers the type from the value.)
+ `x := value` (Go’s compiler infers the type from the value; you can only use this short form within a function.)

You can also define multiple variables at once:

```go
var (
    a = 5
    b = 10
    c = 15
)
```

## Scope

Go is scoped by blocks, and blocks map (more or less) onto curly braces in Go.

## Constants

Go provides constants as well as variables. You declare a constant using the key word `const` instead of `var`. E.g., `const name string = value`.

You can also declare multiple constants:

```go
const (
    a = 5
    b = 10
    c = 15
)
```
