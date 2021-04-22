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
