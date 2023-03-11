# Learning Go: Chapter 3 Composite Types

## Arrays—Too Rigid to Use Directly

Go has arrays, but you will rarely use them. Arrays contain a given number of items, all of which must be the same type. You must specify the type when declaring the array.

```go
var x = [3]int // x[0], x[1], and x[2] all = 0
var x = [3]int{10, 20, 30}
var x = [12]int{1, 5:4, 6, 10: 100, 15} // A stupid way to create a sparse array
```

In the last case, you see that you can mix and match explicit and positional initialization. All unspecified indices get the zero value for whatever type the array stores. (Do not do this. Ever.)

If you use an array literal to initialize an array, you can replace the number with `...`: `var x = [...]{10, 20, 30}`.

You can compare arrays with `==` and `!=`. Go does not support true multidimensional arrays, though you can fake it. You cannot use a negative index to get at the last item in an array.

Go considers the size of an array to be part of its type. As a result, you cannot change the size of an array. You also cannot use a variable to specify the size of an array. You cannot convert arrays of different sizes to identical types. As a result, you cannot write functions that work on arrays of multiple sizes.

In a nutshell, do not use an array unless you know the exact size you need in advance. Some cryptographic functions use arrays because their checksums are defined precisely. But this is the exception to the rule.

## Slices

Where you would use arrays in another language, you probably want slices in Go. A slice is a dynamic list of items of the same type. 

Slices have both a length (which you can check with `len`) and a capacity (which you can check with `cap`). The capacity has to be at least as large as the length, but it can be larger. When the capacity reaches zero, Go will automatically increase the capacity of a slice if the program tries to add to that slice. More precisely, the Go runtime creates a new slice with a larger capacity and copies the item from the first slice to the second slice. In order to minimize how often this happens, Go doubles capacity up until a slice has a capacity of 1024. After that, Go increases the capacity by at least 25%. (I wonder if this is implementation dependent.)

You can create slices in several ways.

+ A slice literal: `var x = []int{1,2,3}` or (inside a function) `x := []int{1,2,3}`. 
+ A nil slice: `var x = []int` or (inside a function) `x := []int`. Both of these create a nil slice of integers. To be clear, this does not mean that this slice stores nil values instead of integers. The slice itself is initially a nil value.
+ A slice with `make`: `var y = make([]int, 0, 20)` or (inside a function) `x := make([]int, 0, 5)`. `make` takes three arguments to create a slice: a type, a length, and a capacity. You can use only two arguments, but in that case, the length and capacity take the value of the second argument. This is bad because it does something you probably don’t expect. Watch:

```go
x := make([]int, 10) // or var x = make([]int, 10)
x = append(x, 10)
fmt.Println(x) // [0 0 0 0 0 0 0 0 0 0 10]

// This is probably better in many cases
y := make([]int, 0, 10) // or var y = make([]int, 0, 10)
y = append(y, 10)
fmt.Println(y) // [0]
```

You can add items to an slice using `append`. You must assign the result of `append` back to the original slice. (If you don’t handle the return value of `append`, the compiler will throw an error. You *can* assign the return value of append to some other slice, but don’t do this!)

### Declaring Your Slice

How does Bodner recommend that we declare slices?

+ If the slice may not grow at all (“because your function might return nothing”), use `var` and create a `nil` slice. E.g., `var data []int`.
+ If you have starting values, or if a slice’s values won’t change, then use a slice literal. E.g., `data := []int{2,4,6,8}`.
+ If you have a reasonable sense (Bodner says “a good idea”) of how large the slice should be, but you don’t yet know the values you will store, use `make`. But you still have to decide what to do about the length value. He lists three possibilities:

1. If the slice will be a buffer, specify a non-zero length. (I’m not sure why, but okay.)
1. If you are certain about the size, specify the length and index into the slice to set values. (I.e., do not use `append` to add values.)
1. Otherwise, use `make` with a zero length and a specified capacity. You can then add items with `append`.

According to Bodner, the Go world is split between 2 and 3, but he prefers 3.

### Slicing Slices

You can return a group of values from a slice using a *slice expression*. Slice expressions take multiple forms, but let’s start with something simple. You give a starting and ending offset separated by a colon. E.g., `x[1:4]` returns items 1, 2, and 3 from the slice `x`. The starting offset is inclusive, and the ending offset is exclusive.

When you take a slice of a slice, you do not copy the data, you share memory between multiple variables. That means that changes to one element in a slice affect all slices that share that element.

However, you can use full slice expressions to limit how much memory is shared
between two slices.

```go
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2:2]
	z := x[2:4:4]
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
```

This protects changes *beyond that limit* from affecting other slices. But if you make changes to the earlier part of these slices, they will still affect each other.

You can convert an array into a slice using a slice expression.

```go
x := [4]int{1,2,3,4}
y := x[:2]
```

If you take a slice from an array, the array and slice will share memory in the same way that two slices sometimes share memory. Be careful about this.

### `copy`

If you want independent copies of a slice, you should use the `copy` function. This function takes two parameters: a destination slice and a source slice. The function copies values from the source to the destination. The function copies as many values as possible, but the maximum number is smallest length of the source and destination. (The capacity of source and destination does not matter, only their length.) You can use slice expressions with both destination and source. The `copy` function returns the number of elements copied. If you don’t need to know that number, you can ignore the return value of `copy`, and Go’s runtime will not complain.

You can use arrays as the source or destination with `copy` if you use a slice expression with the array.

### Strings and Runes and Bytes

Go represents strings as a sequence of bytes. Nothing requires the bytes to be in any specific encoding, but many Go libraries assume that strings are UTF-8 encoded.

You can treat a string like an array or slice in some ways. First, you can index into a string. Second, strings use zero-based indexes. You can also use slice expressions with strings.

Since strings are usually UTF-8 encoded sequences, however, you should be careful about dealing with them one byte at a time. A single UTF-8 code point can be one, two, three, or four bytes. Therefore, if you try to do things (e.g., reverse a string) one byte at a time, you may create invalid UTF-8 sequences. Go provides libraries to deal with different encodings and handle UTF-8 properly. Use those tools instead of raw bytes!

You can convert between strings, runes, and bytes. First, you can use `string` to convert a rune or byte into a string. You can also use `[]byte` or `[]rune` to convert a string into a slice of bytes or runes. Since Go usually handles data as a sequence of bytes, you will often convert strings to byte slices and byte slices to strings, but Bodner says it’s less common to convert either into a slice of runes.

## Maps

Arrays and slices are for sequences, but maps are for associations or lookup tables. You write a map type as `map[keyType]valueType`, and you can declare maps in several ways.

+ You can create a nil map with `var`: `var nilMap map[string]int`. The zero value for a map is `nil`, and the value of a map you create this way is `nil`. To be clear, that’s the value of *the entire map* not of values in the map. The nil map has length zero, and if you try to read from it, it returns the zero type for the map’s value type. (E.g., in this case for any key, the map returns `0`.) However, if you try to write to a nil map, Go panics.
+ You can also use `:=` and an empty map literal to create an empty (but not nil) map: `totalWins := map[string]int{}`. This map is also initially empty, but you can read from and write to this map safely.
+ You can also create a nonempty map using a literal: `totalWins := map[string]int { "Yankees": 5, "Red Sox": 0, }`. The trailing comma after the last value is required (the same as for struct literals).
+ You can use `make` with an initial number of how many key-value pairs you expect to put into the map. The map can grow beyond the initial size, and such a map still starts out with a length of zero. It looks like this: `ages := make(map[int][]string, 10)`. You can specify an initial size of 0 and increase it too: [as in this example](https://play.golang.org/p/IUemhABeRqU).

Other fun facts about maps:

+ Maps grow automatically as you add key-value pairs.
+ You can use `make` to create a map with an initial size.
+ The `len` function will tell you how many key-value pairs a map has.
+ The zero value for a map is `nil`.
+ You cannot compare maps with `==` or `!=`. You have to write your own function to compare maps.
+ The keys for a map can be any comparable type. Thus, you cannot use a slice or a map as the key for a map.

People often use the “comma ok” idiom with maps. Let’s check it out below:

```go
m := map[string]int{
    "hello": 5,
    "world": 0,
}
v, ok := m["hello"]
```

The second return value of a map lookup is a boolean. The return value will be true if the key is in already the map and false if the key is not yet in the map. You need to do this because even if the key is not in the map, the lookup will assign the nil value for that type. (Compare other languages that throw an error if you try to use the value from a key that's not in the map.)

You can delete from a map using Go’s built-in `delete` function: `delete(<map>, <key>)`. `delete` has no return value, and it does not throw an error if the key isn’t present or if the map itself is `nil`.

## Structs

Structs allow you to keep data together in one item in a more complex way than maps. You define a struct with a type in the following way:

```go
type person struct {
    name string
    age int
    address string
}
```

After that, we can define `person` variables: `var fred person`.  We can also assign a struct literal to a variable: `bob := person{}`. Both of these initialize all the fields in the struct to their zero values ("" and 0, in this case). For maps, there’s a difference between `var whatever map[whatever]whatever` and `whatever := map[whatever]whatever{}`. In the first case, you cannot add values to the map, but in the second you can. With structs, both types of definition give you a struct ready to use.[^1]

There are two different ways to assign a non-empty struct literal.

```go
peter := person{
    "Peter",
    52,
    "Address…", // Note the trailing comma; it’s required!
}

bob := person{
    name: "Bob",
    address: "Address…",
}
```

The first style requires you to list a value for all fields in the order they were declared in the type. The second style allows you to put items in whatever order you want, and you can leave items out. However, you must give the name of the field as well as the value of the initial assignment. If you do not list a value here, that field receives the nil value for its type. In this case, Bob’s age is initially 0.

You can read or change the value in a field using dot notation: `bob.age = 26`.

You can also use anonymous structs, which are structs without types. Here are two examples.

```go
var person struct {
    name string,
    age int,
    pet string,
}

person.name = "Bob"
person.age = 50
person.pet = "dog"

pet := struct {
    name string
    kind string
}{
    name: "Fido",
    kind: "dog",
}
```

Both `person` and `pet` are anonymous structs: there is no name for the type of struct they belong to.

Bodner mentions two cases where anonymous structs are useful. First, when you need to marshal external data (say JSON or a protocol buffer). Second, many people use anonymous structs to write table-based tests.

Some structs are comparable, and some are not. If a struct has all the same types, and those types are comparable, then the struct is comparable. However, there is one additional rule: the fields must have the same names, types, *and order*.

As a wrinkle, you can compare and assign back and forth between a struct and an anonymous struct if their fields have the same types in the same order and all the fields are themselves comparable. (I don’t think that this will come up often, but maybe I’m wrong?)

If two structs are not comparable directly using `==`, you can write your own function to compare them.

[^1]: Note that if a struct field has a `map` type, that field has to be initialized with `make` or a map literal, like any other map.
