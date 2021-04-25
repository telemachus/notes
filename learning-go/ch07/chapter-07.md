# Chapter 7: Types, Methods, and Interfaces

Go is a statically typed language (i.e., the type of an item cannot change and every item has a type), and Go has both built-in types (e.g., int, map, string) and user-defined types. You can attach methods to types, and Go provides type abstraction “allowing you to write code that invokes methods without explicitly specifying the implementation.”

## Types in Go

You can use any primitive type or compound type literal to create a user-defined concrete type. Here are some examples.

```go
type Person struct {
    FirstName string
    LastName string
    Age int
}

type Score int
type Converter func(string)Score
type TeamScores map[string]Score
```

You can declare types at any block level, and in general you can only access the type from within its scope. Exported package block level types are an exception. (More on these in Chapter 9.)

## Methods

Go supports methods for user-defined types. You define these at the package block level.

```go
type Person struct {
    FirstName string
    LastName string
    Age int
}

func (p Person) String() string {
    return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.age)
}
```

Method definitions look mostly like function definitions. The obvious difference is that you specify a receiver between the keyword `func` and the the name of the method (here, `String`). The Go idiom is to use a short abbreviation of the type’s name—usually its first letter. (Don’t use `this` or `self`. It’s not idiomatic in Go.)

You call methods as you would expect.

```go
p := Person {
    FirstName: "Fred",
    LastName: "Jones",
    Age: 47,
}
output := p.String()
```

### Pointer Receivers and Value Receivers

You can define methods on values or on pointers. Bodner offers the following rules and advice. First the rules:

+ If a method modifies the receiver, it *must* use a pointer receiver.
+ If a method needs to allow `nil` instances, it *must* use a pointer receiver.
+ If a method doesn’t modify the receiver, it *can* use a value receiver.

Now the advice: when a type has *any* pointer receiver methods, common practice dictates that you should be consistent and use pointer receivers for *all* methods, even those that don’t modify the receiver. Here is an example of a pointer receiver and a value receiver.

```go
type Counter struct {
    total int
    lastUpdated time.Time
}

func (c *Counter) Increment() {
    c.total++
    c.lastUpdated = time.Now()
}

func (c Counter) String() string {
    return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}
```

You can use the code as follows.

```go
var c Counter
fmt.Println(c.String())
c.Increment()
fmt.Println(c.String())
```

In the call `c.Increment()`, you are calling `Increment` on *a value* not a pointer. However, behind the scenes, Go automatically converts the item to a pointer type. That is, you can type `c.Increment()` instead of the clumsier `(&c).Increment()`.

However, you can’t use this convenience if you pass a value type into a function. Here’s an example:

```go
func doUpdateWrong(c Counter) {
    c.Increment() // This won’t work
}

func doUpdateRight(c *Counter) {
    c.Increment()
}
```

In `doUpdateRight`, we pass a pointer instance, but in `doUpdateWrong`, we pass a value type. As a result, you are invoking the `Increment` method on a copy, and your update to the item is lost.

Go considers both pointer and value receiver methods to be in the *method set* of pointer instances. For a value instance, however, Go considers only the value receiver methods to be in its method set. Bodner warns that this will matter when he looks at interfaces.

Finally, Bodner explains that it is not idiomatic to write simple `getter` and `setter` methods for Go structs (unless they need to meet an interface). Instead, “Go encourages you to directly access a field.” If you need to update multiple fields or if the update isn’t a straightforward assignment of a new value, then you should create a method. (He points to `Increment` as an example.)

### Code Your Methods for `nil` Instances

What happens if you call a method on a `nil` instance? Well, that depends. If you call a method with a value receiver on a `nil`, Go will throw a panic, and you’re done. In the case of a pointer receiver, you can make it work if you write code to handle the `nil` instance.

```go
type IntTree struct {
    val int
    lef, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
    if it == nil {
        return &IntTree{val: val}
    }
    if val < it.val {
        it.left = it.left.Insert(val)
    } else if val > it.val {
        it.right = it.right.Insert(val)
    }
    return it
}

func (it *IntTree) Contains(val int) bool {
    switch {
    case it == nil:
        return false
    case val < it.val:
        return it.left.Contains(val)
    case val > it.val:
        return it.right.Contains(val)
    default:
        return true
    }
}

func main() {
    var it *IntTree
    it = it.Insert(5)
    it = it.Insert(3)
    it = it.Insert(10)
    it = it.Insert(2)
    fmt.Println(it.Contains(2)) // true
    fmt.Println(it.Contains(4)) // false
}
```

Bodner admires the way that Go “allows you to call a method on a `nil` receiver,” but he also thinks that “most of the time it’s not very useful.” In general, you don’t need or want to call methods on `nil` pointer instances. However, if your pointer receiver method won’t work for a `nil` receiver, he still warns you to check for `nil` and return an error when one is found. (Otherwise, the method will panic.)

### Methods Are Functions Too

Methods share many features with methods. In particular, you can use a method instead of a function as a variable or parameter of function type. For example:

```go
type Adder struct {
    start int
}

func (a Adder) AddTo(val int) int {
    return a.start+ val
}

myAdder := Adder{start: 10}

f1 := myAdder.AddTo // This is called a *method value*
fmt.Println(f1(10)) // prints 20

f2 := Adder.AddTo  // This is called a *method expression*
fmt.Println(f2(myAdder, 15)) // prints 25
```

For a method expression, the first parameter is the receiver for the method. The function signature in this case is `func(Adder, int) int`.

### Functions Versus Methods

When you should use a function and when should you use a method? Bodner says that it all depends on state and data.

> [You choose a function or a method depending on] whether or not your functions depend on other data. As we’ve covered several times, package-level state should be effectively immutable. Any time your logic depends on values that are configured at startup or changed while your program is running, those values should be stored in a struct and that logic should be implemented as a method. If your logic only depends on the input parameters, then it should be a function.

### Type Declarations Aren’t Inheritance

You can declare types based on built-in Go types and struct literals and also based on user-defined types. However, Go does not have inheritance. You can’t use the two types interchangeably: if a method expects an `int`, you can’t give it a type with an underlying type of `int`. They don’t share a type. In addition, two user-defined types do not share the same methods. You have to define methods for each type separately. Finally, if a user defines a type based on a built-in type, then the built-in operators for that type do work on the user-defined type. (E.g., `+` for a type defined with an underlying type of `int`. You can also assign literals and constants of the underlying type to the user-defined type.)

### Types Are Executable Documentation

Bodner argues that creating user-defined types helps to make code clearer. The user-defined type “provid[es] a name for a concept and describ[es] the kind of data that is expected. It’s clearer for someone reading your code when a method has a parameter of type `Percentage` than `int`, and it’s harder for it to be invoked with an invalid value.”

### `iota` Is for Enumerations—Sometimes

Go does not have enumerations, but they do have a similar feature. You can use constants together with `iota` to produce a kind of enumeration. Bodner recommends that first you define a type based on `int` for all the values in the `iota` block. Here’s what it looks like.

```go
type MailCategory int
const (
    Uncategorized MailCategory = iota
    Personal
    Spam
    Social
    Advertisements
)
```

The Go compiler assigns 0 to `Uncategorized` and then increments the value for each subsequent item.

Bodner recommends that you only use `iota` in specific, limited circumstances. First, you should only use `iota` if you want distinguish a set of values, but you don’t care what the actual value is. You shouldn’t use `iota` for values that already exist (e.g., HTTP responses). Second, you should generally only use `iota` for internal code. If there is a set of values that are defined somewhere else, create constants for those with those values. Don’t use an `iota` and then force yourself to translate back and forth between the two representations.

## Use Embedding for Composition

You can embed structures within structures. In this way, [Go provides something like inheritance via composition and promotion](https://play.golang.org/p/sP1kcVbrbe__q).

```go
type Employee struct {
    Name string
    ID   string
}

func (e Employee) Description() string {
    return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
    Employee
    Reports []Employee
}

m := Manager{
    Employee: Employee{
        Name: "Bob Bobson",
        ID: "12345",
    },
    Reports: []Employee{},
}
fmt.Println(m.ID)
fmt.Println(m.Description())
```

Within `Manager`, `Employee` is an embedded fields. (Notice that no name is assigned to that field.) Any fields or methods that an embedded field has are promoted to the containing struct, and you can invoke such fields or methods directly on it. You can embed any type within a struct not only another struct.

If the containing and embedded types both have fields or methods with the same name, then you have to be more explicit in order to reach the embedded field or method.

```go
type Inner struct {
    X int
}

type Outer Struct {
    Inner
    X int
}

o := Outer{
    Inner: Inner{
        X: 10,
    },
    X: 20,
}
fmt.Println(o.X)       // 20
fmt.Println(o.Inner.X) // 10
```
