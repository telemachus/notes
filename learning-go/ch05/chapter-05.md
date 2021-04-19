# Chapter 5: Functions

## Declaring and Calling Functions

You declare functions in Go using the keyword `func`. You must specify the types of parameters and return values. If a function returns a value, you *must* use the keyword `return` in the body of the function. If the function does not return a value, you can use the keyword `return` to leave the function early.

```go
func div(numerator int, denominator int) int {
    if denominator == 0 {
        return 0
    }
    return numerator / denominator
}

// If there are multiple parameters or return values of the same type, you can
// give the type once only.
func div(numerator, denominator int) { ... }
```

Calling a function is straightforward: `result := div(5, 2)`.

### Simulating Named and Optional Parameters

Go does not support named or optional parameters, but you can simulate both by using a struct for the parameters of a function. However, Bodner feels that if you limit functions to a small number of parameters, you won’t miss named or optional parameters.

### Variadic Input Parameters and Slices

If you need to pass a varied number of parameters to a function in Go, you need a variadic parameter. You do that using `...`. The variadic parameter must be the last (or only) parameter, and you put `...` before the type. Here’s an example:

```go
func addTo(base int, vals ...int) []int {
    out := make([]int, 0, len(vals))
    for _, v := range vals {
        out = append(out, base+v)
    }
    return out
}
```

### Multiple Return Values

You can return multiple values easily in Go.

```go
func divAndRemainder(numerator, denominator int) (int, int, error) {
    if denominator == 0 {
        return 0, 0, errors.New("cannot divide by zero")
    }
    return numerator / denominator, numerator % denominator, nil
}
```

If a function returns multiple values, then in its definition, you put parentheses around the list of return types. However, you must not use parentheses when you return. (If you try, Go will throw a compile-time error.) Also, a function that returns multiple values must return all of them. By convention, any error returned should be the last (or only) value returned.

### Multiple Return Values Are Multiple Values

Bodner makes sure that Python programmers don’t get confused. If your function returns multiple values, then it returns *multiple* values not a single tuple that can be (optionally and automatically) destructured by assignment.

### Ignoring Returned Values

In Go, you must assign each value returned from a function. You cannot ignore some. However, if you assign a variable and never use it, you have a different problem since Go also does not allow unused variables. The answer is the dummy variable `_`. Place `_` wherever you don’t care about a return value. E.g., `result, _, err := divAndRemainder(5, 2)`.

Weirdly, you can ignore *all* of the return values for a function.

### Named Return Values

If you want, you can name the return values as well as the parameters for a function.

```go
func divAndRemainder(numerator, denominator int) (result int, remainder int, err error) {
    if denominator == 0 {
        err = errors.New("cannot divide by zero")
        return result, remainder, err
    }
    result, remainder = numerator/denominator, numerator%denominator
    return result, remainder, err
}
```

Note that named return values are initialized with the zero value for their type in the body of the function. Thus, if you don’t create an error for `err`, you are simply returning `nil` when you return `err`.

However, Bodner recommends that you not use this feature of Go. Why not? First, he worries about shadowing other variables. Second, you can accidentally ignore them in the body of the function and end up with weird surprises.

When you use `defer`, named return values can be useful, but Bodner will explain that later in the chapter.

### Blank Returns—Never Use These!

Go supports blank (aka, naked) returns. At the end of a function with named return values, you can simply write `return`.

```go
func divAndRemainder(numerator, denominator int) (result int, remainder int, err error) {
    if denominator == 0 {
        err = errors.New("cannot divide by zero")
        return
    }
    result, remainder = numerator/denominator, numerator%denominator
    return
}
```

## Functions Are Values

The type of a function is made from the keyword `func`, the types of the functions parameters, and the types of any return values. This combination is called the function’s *signature*. All functions with the same number and type of parameters and return values have the same signature.

Since functions are values, you can store functions as values in a map. Bodner gives [an example to make a calculator](https://oreil.ly/L59VY). It’s kind of cool, but so far I can’t see a real use for this.

### Function Type Declarations

You can use `type` to give a name to function types in the same way you use it to name types of struct.

```go
type opFuncType func(int, int) int

var opMap = map[string]opFuncType {
}
```

Bodner mentions two good uses for naming function types. First, you gain documentation through the name. Second, function types help get us to interfaces. (He’ll discuss interfaces later.)

### Anonymous Functions

You can assign anonymous functions to variables or use them inline. Bodner says that anonymous functions are useful for `defer` statements and goroutines. Here’s a less useful example for illustration.

```go
func main() {
    for i := 0; i < 5; i++ {
        func(j int) {
            fmt.Println("printing", j, "from inside an anonymous function")
        }(i)
    }
}
```

## Closures

If you declare a function inside of a function, the inner function is a closure, which Bodner explains as “a computer science word that means that functions declared inside of functions are able to access and modify variables declared in the outer function.” What advantages does this have?

+ If you declare a function inside a function, you hide the inner function. This means that you have fewer package-level declarations, which can make naming easier.
+ More importantly, you can pass functions to other functions or return them from other functions. This means that you can carry variables created within a function and use them outside of the function where they are created—without using global state.

### Passing Functions as Parameters

Bodner uses `sort.Slice` to explain passing functions as parameters.

```go
type Person struct {
        FirstName string
        LastName  string
        Age       int
}

people := []Person{
        {"Pat", "Patterson", 37},
        {"Tracy", "Bobbert", 23},
        {"Fred", "Fredson", 18},
}
fmt.Println("Unsorted:", people)

// sort by last name
sort.Slice(people, func(i int, j int) bool {
        return people[i].LastName < people[j].LastName
})
fmt.Println("Sorted by last name:", people)

// sort by age
sort.Slice(people, func(i int, j int) bool {
        return people[i].Age < people[j].Age
})
fmt.Println("Sorted by age:", people)
```
