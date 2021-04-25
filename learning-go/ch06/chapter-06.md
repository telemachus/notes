# Chapter 6: Pointers

## A Quick Pointer Primer

Pointers refer to variables that track the memory address of a value. Different types take up different amounts of memory, but if you have the address, you can always get at the value. The zero value for a pointer is `nil`.

```go
var x int32 = 10
var y bool = true
pointerX := &x
pointerY := &y
var pointerZ *string
```

Each of the pointers above refers to a memory address. `pointerZ` has the zero value for a pointer, namely `nil` because it doesn’t yet point to an actual value. (`nil` is also the zero value for maps, slices, functions, channels, and interfaces. This is not a coincidence: all of these involve pointers.)

The `&` is the address operator. It returns the address of the memory location where a value is stored.

The `*` is the indirection operator. It returns the value of the variable you apply it to. (When you apply `*` to a variable, you *dereference* the variable. If you try to dereference a `nil` pointer, Go will panic.

A pointer type is `*` plus the name of a type. You can have a pointer type for any type. E.g., `*int` or `*string`.

Literals and pointers don’t mix well. Primitive literal values (numbers, booleans, and strings) lack a memory address. Thus, you cannot assign a primitive literal using `&`. You also can’t assign the value of a literal to a `nil` pointer. The following don’t work.

```go
var pString *string
*pString = "Hello, world!"
// panic: runtime error: invalid memory address or nil pointer dereference

pString = &"Hello, world!"
// cannot take the address of "Hello, world!"

*pString := "Hello, world!"
// non-name *pString on left side of :=
```

You can use `new` to create a pointer variable. It returns a pointer to a zero value of the given type. You *can* assign to this.

```go
var pString = new(string)
*pString = "Hello, world!"
```

However, Bodner says that “[t]he `new` function is rarely used.” For structs, you can use `&` before a literal. For pointers to numbers, booleans, and strings, you cannot assign literals directly. (This is true for struct fields that are pointers to these types as well.) Bodner recommends either using a variable that holds the necessary value or using a function that accepts a boolean, number, or string and returns a pointer to that type.

```go
func stringPointer(s string) *string {
    &s
}
```

## Pointers Indicate Mutable Parameters

Since Go is pass-by-value, functions receive copies. As a result, you normally cannot change the value of arguments to functions. This is good, says Bodner, since mutability is bad. However, sometimes you need to change the value of an argument, and pointers make this possible. Bodner expands on this with the following details.

First, if you pass a `nil` pointer, you cannot make the value of the argument non-nil. There’s no address to hang a value on. As Bodner puts it, “You can only reassign the value if there was a value already assigned to the pointer.” The following update fails.

```go
func failedUpdate(g *int) {
    x := 10
    g = &x
}

func main() {
    var f *int // f is nil
    failedUpdate(f)
    fmt.Println(f) // prints nil
}
```

Second, if you want to change the value of a pointer parameter, you need to dereference the pointer and set a new value. (If you try that above, Go will panic since you cannot dereference `nil`.) The following update works.

```go
func  update(px *int) {
    *px = 20
}

func main() {
    x := 10
    update(&x)
    fmt.Println(x) // prints 20
}
```

## Pointers Are a Last Resort

Bodner argues that you should avoid using pointers in Go if you can. Why? Not because they are difficult to understand or use—Bodner claims that they are not difficult to understand or use. Because Bodner thinks that mutability is bad and makes programs harder to understand.

There is, however, an important exception: “you should use pointer parameters to modify a variable…when the function expects an interface.” An example is working with JSON.

```go
f := struct {
    Name string `json:"name"`
    Age int `json:"age"`
}
err := json.Unmarshal([]byte({"name": "Bob", "age": 30}`), &f`
```

`Unmarshal` expects an interface as its second parameter, which must be a pointer.

As a side point, Bodner thinks that new Go developers come to believe that pointers are common because JSON use is so common. Instead, he insists that the use of pointers should be an exception.

In general, Bodner argues that you should return value types rather than pointer types from function. However, “use a pointer type as a return type if there is state within the data type that needs to be modified.” Two examples are I/O and concurrency. Bodner will talk about both of these later.

## Pointer Passing Performance

Bodner discusses the differences in performance between pointers and data types as arguments and return values for functions. I don’t think that I need to worry about this yet. However, and briefly, passing or returning pointers takes constant time, but data types vary—and grow worse—as the size of the data increases. Weirdly, to pass a pointer *into* a function takes roughly one nanosecond, but to return a pointer *from* a function takes roughly thirty nanoseconds. In any case, I won’t think about this initially.

## The Zero Value Versus No Value

You can use a pointer to distinguish between a field with a zero value (e.g., 0 or "") and a field that has never had a value assigned to it (i.e., `nil`). Bodner recommends using the "comma ok" idiom instead of a `nil` pointer.

JSON is again an exception. When you work with JSON, you often need to distinguish between a zero value and no value at all. In this case, a pointer value makes sense for fields that can be null.

## The Difference Between Maps and Slices

When you pass a map to a function, you can make changes to the original variable because Go implements maps as pointers to functions. This may seem convenient, but Bodner recommends not using maps for input parameters or return values. They are opaque and mutable, and both are bad for clarity. Instead, he recommends that you use structs.

Bodner also explains in detail why slices behave differently from maps when you pass them to functions. You can change everything about maps, but slices are more complicated. As I mentioned earlier, you can change values in a slice, but you cannot change its length—or capacity. Go implements slices as structs with three fields: a pointer to a memory location, an int for length and and int for capacity. When you pass a slice to a function, the function receives a *copy* of those three values. Since you have the memory address, you can alter the items in the slice. But if you make any changes to the length or capacity of the slice in the function, you change *the copies* of the length and capacity. Thus, the calling code cannot see those changes.

Finally, Bodner says that you should avoid modifying slices (even their contents) that you pass to functions. If you have to modify the content, you should document the changes clearly.

## Slices as Buffers

If you want to read data from an external resource (e.g., a file or network connection), then slices make good buffers.

```go
file, err := os.Open(fileName)
if err != nil {
    return err
}
defer file.Close()
data := make([]byte, 100)
for {
    count, err := file.Read(data)
    if err != nil {
        return err
    }
    if count == 0 {
        return nil
    }
    process(data[:count])
}
```

## Reducing the Garbage Collector’s Workload

When you use pointers, you give Go’s garbage collector more work. Try to avoid pointers when you can.
