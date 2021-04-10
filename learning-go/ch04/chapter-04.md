# Chapter 4: Blocks, Shadows, and Control Structures

## Blocks

Go handles scoping via blocks. A block determines what can be seen and where. If you declare variables outside of any functions, then they are in the package block. They can be seen by the entire package. When you import something, it is visible within the file block. A package can span many files, so package block may be larger than file block. There is also a block for every function (function block), and control structures (e.g., if, for, and switch) define their own blocks.

You can see any identifier defined in an outer block within any inner block. But in the outer block, you cannot see identifiers defined in an inner block. Such identifiers have gone out of scope. You can also use the same identifier more than once, and in that case, you will end up shadowing the identifier.

### Shadowing Variables

f you create a new variable with the same name as a variable in a containing block, the second variable shadows the first. From the inner block, you cannot reach the shadowed variable, but it reappears (so to speak) when the smaller block ends. Here’s an example:

```go
func main() {
    x := 10
    if x > 5 {          // start of inner block
        fmt.Println(x)      // prints 10
        x := 5              // shadows original x; you probably want x = 5
        fmt.Println(x)      // prints 5
    }                   // end of inner block; inner x no longer exists
    fmt.Println(x)          // prints 10
```

## `if`

Go uses `if...[else if]...[else]`. Brackets are required, but you don’t (normally? ever?) need to add parentheses around the condition.

```go
n := rand.Intn(10)
if n == 0 {
    // ...whatever
} else if n > 5 {
    // ...whatever
} else {
    // ...whatever
}
```

Go allows you to create variables that only exist for as long as the `if...else` statement. After the following code runs, `n` is no longer visible.

```go
if n := rand.Intn(10); n == 0 {
    // ...whatever
} else if n > 5 {
    // ...whatever
} else {
    // ...whatever
}
```

## `for`

Go provides only one keyword for loops: `for`. However, you can use `for` in four different ways in Go.

+ A complete, C-style `for`
+ A condition-only `for`
+ An infinite loop `for`
+ `for...range`

### The Complete `for` statement

This is exactly what you think it will be: `for i := 0; i < 10; i++ { ... }`. You don’t need parentheses around the three parts of a `for` loop in Go.

### The Condition-Only `for` Statement

You can leave off the initialization and increment (first and third part) of a complete `for` loop, leaving only the condition. This makes a `for` that is equivalent to a `while` loop in other languages.

```go
i := 1
for i < 100 {
    fmt.Println(i)
    i *= 2
}
```

### The Infinite `for` Statement

You can also omit all three parts of the complete `for` loop. This creates an infinite loop. E.g. `for { ... }`. You can use `break` to get out of the loop. Go also provides a `continue` keyword. Bodner recommends using `continue` to prevent deeply nested conditions. Here’s his example:

```go
for i := 1; i <= 100; i++ {
    if i%3 == 0 && i%5 == 0 {
        fmt.Println("Fizzbuzz")
            continue
    }
    if i%3 == 0 {
        fmt.Println("Fizz")
            continue
    }
    if i%5 == 0 {
        fmt.Println("Buzz")
            continue
    }
    fmt.Println(i)
}
```

Because we use `continue`, we avoid having nested conditions and branches.

### The `for-range` Statement

You can use `for` together with `range` in order to iterate over built-in and user-defined compound types.

```go
evenVals := []int{2, 4, 6, 8, 10, 12}
for i, v := range evenVals {
    fmt.Println(i, v)
}
```

When you loop using `range`, you get two variables per loop. The first is the index of the item, and the second is the item itself. (Go users tend to default to `i` when the first value is an index of a slice or array and `key` when the first value is a map key.) If you don’t need the first of the two values, use `_` to tell Go’s runtime that the value doesn’t matter and won’t be used. If you don’t need the second value, you can simply leave it off. Go considers this an exception to its usual rule about return values. (**NB**: In the case of channels, `range` only returns a single value, apparently. Bodner will discuss this further in Chapter 10.)

When you iterate over strings using `range`, the iterator proceeds rune-by-rune rather than byte-by-byte. If your string has an invalid UTF-8 value, you get the Unicode replacement character (hex 0xfffd) instead of the character.

When you iterate over a compound type using `range`, you are iterating over a copy. Therefore, you cannot modify the original by changing the value inside the loop. However, you *can* modify the original compound type if you use the index to reach into the compound. Consider the example below.

```go
evenVals := []int{2, 4, 6, 8, 10, 12}
for i, v := range evenVals {
    v = v*2                             // This has no effect on evenVals.
    evenVals[i] = evenVals[i] * 2       // This changes evenVals.
}
```

You can use `break` and `continue` with `for-range` loops.

### Labeling Your `for` Statements

You can use a label in order to tell `break` or `continue` which loop to jump out of. This comes in handy when you have nested loops, and you want to make sure that, say, you only run some code if a specific condition isn’t met. Here’s Bodner’s example:

```go
outer:
    for _, outerVal := range outerValues {
        for _, innerVal := range outerVal {
            // Process innerVal.
            if invalidSituation(innerVal) {
                continue outer
            }
        }
        // Here we have code that runs only when all of the
        // innerVal values were successfully processed.
    }
```

Notice that the label (`outer:`) is outdented, relative to the `for` loop. This is normal.

### Choosing the Right `for` Statement

Bodner gives the following advice about picking the right `for` statement.

+ Use `for-range` in general when you are iterating over all the contents of a built-in compound type. (I think that this also applies to user-defined compound types based on the built-in types.)
+ However, you should use the complete `for` loop when you don’t want to iterate over the whole collection. (You may know that you want to start late or end early, for example.) However, you can’t use a complete `for` loop in this way to iterate over part of a string. In that case, you still need `for-range` to give you runes. Inside the loop, you have to write code to skip items or break out of the loop.
+ You will need the condition-only and the bare `for` loop a lot less often.
+ Use the condition-only `for` loop in the same cases where you would use a `while` loop in another language. A simple rule of thumb is that you use `for` if you know how many times the loop will run, and you use `while` if you know only the condition that will stop the loop.
+ Another consideration in Go, however, is that a condition-only `for` loop suits some scoping situations better than a complete `for` loop. If the variable already exists or should exist after the loop, then a condition-only `for` loop will be better.
+ Bodner offers two cases where an infinite `for` loop can be useful. First, if you want to simulate a `do...while` loop in Go, you should use `for...if { break }`. Second, the infinite `for` loop is good to implement the iterator pattern, which Bodner will show in a later chapter.

## `switch`

Go has a more useful `switch` statement than other C-related languages. Bodner begins with the following example.

```go
func main() {
    words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
    for _, word := range words {
        switch size := len(word); size {
        case 1, 2, 3, 4:
            fmt.Println(word, "is a short word!")
        case 5:
            wordLen := len(word)
            fmt.Println(word, "is exactly the right length:", wordLen)
        case 6, 7, 8, 9:
        default:
            fmt.Println(word, "is a long word!")
        }
    }
}
```

What should we notice here?

+ You don't need parentheses around the condition, as with `if` statements.
+ You can create a locally scoped variable for the `switch`, as with `if` statements.
+ The `switch` statement requires braces, but you don’t need braces for the `case` or `default` clauses.
+ You can have multiple lines inside one `case` or `default` clause, and you still don’t need braces.
+ A variable like `wordLen` is scoped to its block. Thus, `wordLen` is only visible inside of the `case 5:` clause.
+ By default, there is no fall-through. There is a `fallthrough` keyword, but Bodner recommends that you not use it.
+ If a `case` clause has no code in its block (as here for `case 6, 7, 8, 9:`, then nothing happens for that clause. (It does *not* fall through to the `default` behavior.
+ If you need multiple items to go in one clause, you can use commas, as here in the first and third clauses.
+ In this code, `switch` operates on an integer. But you can switch on anything that can be compared with `==`.

### Blank Switches

You can leave out the comparison value in a blank switch. That allows you to write more flexible `switch` statements. Here’s an example:

```go
words := []string{"hi", "salutations", "hello"}
for _, word := range words {
    switch wordLen := len(word); {
    case wordLen < 5:
        fmt.Println(word, "is a short word!")
    case wordLen > 10:
        fmt.Println(word, "is a long word!")
    default:
        fmt.Println(word, "is exactly the right length.")
    }
}
```

Bodner offers the following advice for choosing between `if` and `switch`: “Favor blank `switch` statements over `if/else` chains when you have multiple related cases. Using a `switch` makes the comparisons more visible and reinforces that they are a related set of concerns.” On the other hand, you should not have “unrelated comparisons on each case in a blank `switch`.” In a case where your comparisons are unrelated, you should use `if/else` statements. (And Bodner adds “or perhaps consider refactoring your code.”)

## `goto`-Yes, `goto`

Go provides `goto`. You probably should not use it, but sometimes `goto` with a label provides cleaner code than the alternative. I’m going to leave it at that for the moment.
