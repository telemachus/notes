# Introducing Go, Chapter 5: Arrays, Slices, and Maps

Now Doxsey introduces some basic composite types.

## Arrays

An array is a fixed-length sequence of elements of a single type. For example, `var x [4]int`. You number the indexes of an array starting from zero.

Doxsey introduces the use of `range` with `for` loops here. Compare the following two loops.

```go
for i := 0; i < len(array); i++ {
    // do whatever
}

for _, value := range array {
    // do whatever
}
```

You can declare an array with initial values on one line or multiple lines. If you use multiple lines, you *must* include a trailing comma, but if you use one line, you *can* include a trailing comma.

```go
x := [5]float64{98, 93, 77, 82, 83,} // fine
x := [5]float64{98, 93, 77, 82, 83}  // fine
x := [5]float64{
    98,
    93,
    77,
    82,
    83 // not fine
}

x := [5]float64{
    98,
    93,
    77,
    82,
    83, // fine
}
```

## Slices

A slice is like an array, but the size is variable. To declare a slice, you say, for example, `var x []float64`. However, there are many ways to create a slice.

```go
var x []float64
x := make([]float64, 5, 10) // a slice with a length of 5 and a capacity of 10
x := make([]float64, 5) // a slice with a length and capacity of 5
y := [5]float64{1,2,3,4,5}
x := y[0:5]
```

### `append`

You can use `append` to add items to the end of a slice. You must assign the return value of `append`.

```go
slice1 := []int{1,2,3}
slice2 := append(slice1, 4, 5)
```

### `copy`

You use `copy` to copy items from one slice into another. If the length of the two slices are not the same, the smaller length is copied. You do not need to assign the return value of `copy`.

```go
slice1 := []int{1,2,3}
slice2 := make([]int, 2)
copy(slice2, slice1) // ignored return value is 2—the number of items copied
```

## Maps

A map is an unordered colleciton of key-value pairs. You can declare a map as follows: `var salaries map[string]int`. However, you cannot use a map declared that way. (Go panics if you try.) Instead, you need to use `make`.

```go
var y map[string]int
y["key"] = 10 // panics

x := make(map[string]int)
x["key"] = 10 // no panic
```

You can test whether a map has a specific key using the comma ok idiom: `name, ok := elements["whatever"]`. If `ok` is false, the map does not contain a key of that name. Here’s an idiomatic way to use the idiom:

```go
if name, ok := elements["Un"]; ok {
    fmt.Println(name, ok)
}
```
