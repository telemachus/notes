# Introducing Go, Chapter 4: Control Structures

Go has few control structures by design. They language is intentionally simple in many ways, and this is one of them.

## The `for` Statement

Doxsey shows two ways to use `for` loops. In the first way he shows, you use only a conditional in the `for` statement itself. In the second way, you have an initialization, a conditional, and an change, and each of the three items is separated by a semicolon. 

```go
i := 1
for i <= 10 {
    fmt.Println(i)
    i++
}

for i := 1; i <= 10; i++ {
    fmt.Println(i)
}
```

## The `if` Statement

Doxsey also shows how to write conditional statements in Go.

```go
if i % 2 == 0 {
    // even
} else {
    // odd
}

if test {
    // something
} else if otherTest {
    // something else
}
```

## The `switch` Statement

Finally, we meet the `switch` statement. For the moment, Doxsey only shows us one of the ways to use a `switch` statement.

```go
switch i {
case 1:
    // something
case 2:
    // something else
default:
    // another thing
}
