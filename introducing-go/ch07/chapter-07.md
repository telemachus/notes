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

### Embedded Types

Doxsey explains that a “struct’s fields usually represent the has-a relationship (e.g., a `Circle` has a `radius`).” This is a great way to think of structs initially: a circle has a radius, a person has a first name, an employee has an id number. But this is not the entire story about structs.

Structs can represent the is-a relationship as well as the has-a relationship. To do this, you use embedded structs.

```go
type Person struct {
    Name string
}

type Android struct {
    Person
    Model string
    ID    int
}

func (p Person) Talk() {
    fmt.Println(p.name, "says hello!")
}

func main() {
    a := new(Android)
    a.Name = "Hal"
    fmt.Println(a.Name)
    fmt.Println(a.Person.Name) // this works too
    a.Talk() // Hal says hello!
}
```

Because an android is a person, and a person has a name, an android has a name. Because an andriod is a person, and a person can talk, an android can talk. You can reach the `Name` field directly or through the embedded Person struct. (In general, you only need to go the long way when both the outer and inner struct have fields with the same name. Also, I think you should avoid giving both structs fields with the same name.)

## Interfaces

Go uses interfaces to gather types with overlapping abilities. For example, multiple shapes will define a method to calculate their area. We can create the following interface for them.

```go
type Shape interface {
    area() float64
}
```

We give the interface a name (as with structs), and then inside we list the methods that the interface must implement. In this case, shapes only have to implement an area method, but there will often be more than one method. How is this useful? Doxsey gives two uses for interfaces.

First, we can use interfaces as arguments to functions. This is effectively a kind of duck typing for a static language. Any item that implements the area method can be an argument to such a function. For example, consider the following.

```go
func totalArea(shapes ...Shape) float64 {
    var area float64
    for _, s := range shapes {
        area += s.area()
    }
    return area
}
```

Second, we can use interfaces as fields in a struct. We can create a struct that stores several shapes of different types as follows.

type MultiShape struct {
    shapes []Shape
}

multiShape := MultiShape{
    shapes: []Shape{
        Circle{0,0,5},
        Rectangle{0,0,10,10},
    },
}
```

Weirdly, `MultiShape` itself can satisfy the shape interface if we give it an area method.

```go
func (m *MultiShape) area() float64 {
    var area float64
    for _, s := range m.shapes {
        area += s.area()
    }
    return area
}
```

Now a `MultiShape` can contain `Circles`, `Rectanges`, and also other `MultiShapes`.

Doxsey says that you shouldn’t focus on types and taxonomies up front. First, write the code you need to implement the behavior of your program. Create structs as needed and add methods to help use them. As the program grows, you will see patterns, and then you can gather them together using interfaces. “There’s no need to have [all your interfaces] figure out ahead of time.”

Go provides packages in order to combine interfaces, types, variables, and functions into a single item. Doxsey will discuss packages in the next chapter.
