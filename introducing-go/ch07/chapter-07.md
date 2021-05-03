# Introducing Go, Chapter 7: Structs and Interfaces

## Structs

Structs are composite data structures that contain named fields. You can create a circle type as follows.

```go
type Circle struct {
    x float64
    y float64
    r float64
}

// This also works.
type Circle struct {
    x, y, r float64
}
```

Every field has a name and a type.

### Initialization

You can create token circles in several ways.

```go
var c Circle // all the fields in c are set to their zero value
c := new(Circle) // new returns a pointer to a struct
c := Circle{x:0, y: 0, r: 5}
c := Circle{
    x:0,
    y: 0,
    r: 5,
}
c := Circle{0, 0, 5} // only do this for very simple obvious struct fields
```

### Fields

You access fields using the `.` operator. E.g., `c.y` and `c.r`.

## Methods

You will often want functions to change fields in structs, but since Go is pass-by-value, you can’t do this unless you use pointers.

```go
type Employee struct {
    firstName string
    lastName string
    salary int
}

// this won’t work
func payRaise(e Employee, raise int) {
    e.salary += raise
}

// this will work
func payRaise(e *Employee, raise int) {
    e.salary += raise
}

func main() {
    var e Employee
    e.salary = 20
    fmt.Println(e.salary)
    payRaise(&e, 20)
    fmt.Println(e.salary)
}
```

As you would imagine, Go would need a lot of functions with structs (or pointers-to-structs) as their first parameter. Therefore, Go provides methods as a cleaner way to implement the same behavior.

```go
func (e *Employee) payRaise(raise int) {
    e.salary += raise
}

func main() {
    var e Employee
    e.salary = 20
    fmt.Println(e.salary)
    e.payRaise(20)
    fmt.Println(e.salary)
}
```

Notice that we don’t need the operator when we call `e.payRaise`. Go automatically passes a value or pointer, depending on the signature of the function.
