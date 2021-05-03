# Introducing Go, Chapter 6: Functions

Doxsey begins with the black box metaphor for functions. A function is like a black box, he explains. You feed it input, and it returns output. He admits, however, that things are more complicated. A function can take *zero or more* inputs and return *zero or more* outputs. He doesn’t use the word, but you might call a code for its side effects rather than return values.

Doxsey gives the following additional tips. First, parameter names and the name of a variable in the calling function can be different. Second, you must pass a variable to a function, or the variable must exist in the functions scope already. Third, functions form a LIFO stack. Fourth, Doxsey shows named return values in Go, though he does not explain them clearly. Fifth, Doxsey specifies that Go allows for multiple return values, and he explains that they need parentheses in the function declaration and comma-separated items in the return statement. He also mentions that multiple values are especially useful for returning an error value or boolean to indicate success or failure.

## Variadic Functions

The last parameter (or only parameter) in a Go function can be variadic.

```go
func add(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

The `...` operator serves two functions. When placed before a parameter, it is the “pack operator,” and it bundles zero or more parameters of a given type into a slice for the function to operate on. Conversely, if you place it after an argument, then it is the “unpack operator,” and it hands zero or more items from a slice as separate arguments to a function.

For example, if we had the `add` function as above, we can use it with the “unpack operator” as below.

```go
func main() {
    grades := []int{80,100,95,}
    total := add(grades...)
    // more code goes here…
}
```

## Closure

Doxsey begins with the idea that you can assign an anonymous function to a variable.

```go
add := func(x, y int) int {
    return x + y
}
```

He then explains that when you create a function, it has access to other local variables in its scope.

```go
x := 0
increment := func() int {
    x++
    return x
}

fmt.Println(increment()) // 1
fmt.Println(increment()) // 2
```

You can also create a function that returns a function *and* that captures some local variables for use. For example, here is an even number generator.

```go
func makeEvenGenerator() func() int {
    i := uint(0)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}

func main() {
    nextEven := makeEvenGenerator()
    fmt.Println(nextEven()) // 0
    fmt.Println(nextEven()) // 2
    fmt.Println(nextEven()) // 4
}
```

## Recursion

Functions can call themselves. Here’s a factorial function in Go.

```go
func factorial(x uint) uint {
    if x == 0 {
        return 1
    }
    return x * factorial(x-1)
}
```

## `defer`, `panic`, and `recover`

The `defer` statement “schedules a function call to be run after the [enclosing] function” ends. You can use it as follows.

```go
func first() {
    fmt.Println("First!")
}

func second() {
    fmt.Println("Second!")
}

func main() {
    defer second()
    first()
}
```

We will see “First!” and then “Second!” in that order. Doxsey lists three advantages for `defer`.

1. It keeps, e.g., a `Close` function near an `Open` function. This makes programs easier to understand.
1. A function may have multiple `return` statements, but we only have to write the `Close` function once.
1. Even if a program panics during runtime, `defer` guarantees that their functions will happen before the panic.

### `panic` and `recover`

You can combine `recover` and `defer` to guard against a runtime `panic`.

```go
func main() {
    defer func() {
        p := recover() // a call to recover returns what was in the panic
        fmt.Println(p)
    }() // defer the anonymous function by adding parentheses
    panic("PANIC!")
}
```

## Pointers

Go passes arguments by value: it copies the value of the arguments into the function. That means that you cannot, in most cases, directly change the original in the function.

```go
func zero(x int) {
    x = 0
}

func main() {
    x := 5
    zero(x)
    fmt.Println(x) // x is still 5
}
```

However, you can write functions that modify their arguments using pointers.

```go
func zero(xPtr *int) {
    x = 0
}

func main() {
    x := 5
    zero(&x)
    fmt.Println(x) // x is 0
}
```

Pointers refer to the location in memory where a value is stored rather than the value itself. If you copy the memory location into a function, you *can* directly modify the value.

### The `*` and `&` operators

The `*` operator has two uses in Go. First, in a declaration, it identifies a pointer type. For example, `int` is an int type, but `*int` is a pointer to int type. Second, the asterisk dereferences pointer variables. In the zero function above, we write `*xPtr = 0` not `xPtr = 0`. If you try the second, Go won’t compile the program because `xPtr` is not an int type! It is a pointer to int. In order to get at the underlying value, you use the dereference operator `*`.

The `&` operator returns the address of a variable. `&x` returns a `*int`, that is a pointer to int type. If you change the program above to `zero(x)`, Go will again fail to compile the program because of a type error. When you declare zero, you promise to pass in a pointer to int, but `x` itself is an int, not a pointer to int.

### `new`

You can also get a new pointer via `new`.

```go
func one(xPtr *int) {
    *xPtr = 1
}

func main() {
    xPtr := new(int)   // x is 0 since pointers begin with their zero value
    one(xPtr)
    fmt.Println(*xPtr) // x is 1
}
```
