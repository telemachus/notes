# Learning Go: Chapter 2 Primitive Types and Declarations

## The Zero Value

Go assigns a zero value to any variable that is declared but not assigned a value.
(Because Go does this, we don’t need to worry about all sorts of problems programmers in C or C++ have.)
You must learn the zero value for the basic types.

## Literals

You have a literal when you write out a number, character, or string.
Go programs use several types of literals.
Here are four important types of literals.

### Integer Literals

Integer literals are sequences of numbers.
They are usually base ten.
But you can also use `0b` for base two (binary), `0o` for base eight (octal), and `0x` for base sixteen (hexadecimal).
You can use upper- or lowercase letters in the prefixes.
You can use a leading 0 without a letter as a base eight number.
Don’t use this short form for base eight: it’s confusing.

You may add underscores to integer literals.
The underscores have no effect on the program, but they may make it easier for humans to read long numbers.
You can’t have multiple underscores in a row, and you can’t put an underscore first or last in a literal number.

### Floating Point Literals

Floating point literals are numbers that allow for fractional parts of the value.
Floating point literals can also use `e` to signify an exponent.
You can write floating point literals as hexadecimal with the `0x` prefix and the letter `p` for any exponent.
You may also use underscores to format floating point literals.

### Rune Literals

A rune literal stands for a character and it must be surrounded by single quote marks.
Go does not allow you to use single and double quotes as you see fit.
You must use single quote marks for runes, i.e., single characters.
You can write out a rune as a single Unicode character, an 8-bit octal, an 8-bit hexadecimal, a 16-bit hexadecimal, or a 32-bit Unicode number.
For example, [all of the following are the same](https://play.golang.org/p/VO7PaTGgBTk):

```go
package main

import (
	"fmt"
)

func main() {
	a := 'a'
	b := '\141'
	c := '\x61'
	d := '\u0061'
	e := '\U00000061'

	fmt.Println(a, b, c, d, e)
	fmt.Println(string(a), string(b), string(c), string(d), string(e))
}
```

In general, people use base 10 to represent number literals, and they avoid hexadecimal escapes for rune literals.
There are exceptions.
People use octal representations for POSIX permissions (e.g., `0o777` instead of `rwxrwxrwx`).
People also use hexadecimal and binary for bit filters or on networking and infrastructure applications.

The zero value for a rune is the empty char.

### String Literals

Go provides two types of string literals.
The most common are interpreted string literals.
Here, you have zero or more rune literals, in any of the legal forms for runes, inside of double quote marks.
You cannot use unescaped backslashes, unescaped newlines, or unescaped double quotes.

Less commonly, you can use a raw string literal.
In this case, you delimit the raw string with backquotes.
Thus, you can use any literal character inside a raw string literal except a backquote.
(Question to self: can you use a backquote in a raw string literal with an escape?
No: [you cannot](https://github.com/golang/go/issues/36042).)

The zero value of a string is the empty string.

## Literals (in General, Again)

Go considers literals (partially) untyped.
As a result, you can use an integer literal in a floating point expression or assign an integer literal to a floating point variable.
But you can’t assign a string to a number.
Bodner puts it this way: “[literals] can interact with any variable that’s compatible with the literal.”
The programmer needs to know what types are compatible with what literals.
Programmers also have to keep in mind size restrictions.
As Bodner says, you can’t assign the literal 1000 to a `byte`.

## Booleans

The type `bool` can have one of two values: `true` or `false`.
The zero value for a `bool` is `false`. 

## Numeric Types

Go has many numeric types in three different categories: integer types, floating point types, and complex types.

### Integer Types

Go has signed and unsigned integers in several sizes: `int8`, `int16`,
`int32`, and `int64`; `uint8`, `uint16`, `uint32`, and `uint64`.
The zero value for all integer types is zero.

Go also provides aliases for integer types.

`byte` is an alias for `uint8`.
You can uses variables of `uint8` and `byte` interchangeably.
You should, however, use `byte` if you mean to use the variable as a byte.

Go also provides `int` as an integer alias.
On a 32-bit CPU, `int` is an `int32`.
On a 64-bit CPU, `int` is an `int64`.
Because of platform independence, you have to explicitly cast an `int` to an `int32` or `int64`.
Integer literals default to the `int` type.

Go provides `uint`, which works the same way as `int`.

Go provides `rune` as an alias for `int32`.

### What Should You Use?

+ If you are working with a format or protocol that uses an integer of a specific size or range, use that integer type.
+ If you are writing a library function that should work with any integer type, write two functions, one for `int64` and one for `uint64`.
+ Otherwise, use `int`.

### Integer Operators

Go provides all the usual operators for integers: `+, -, *, /, %`.
The result of integer division is an integer.
(Integer division truncates towards zero.)
You can combine all the arithmetic operators with `=` to create assignment operators.
You can also compare integers with `==, !=, > >=, <, <=`.

### Floating Point Types

Go provides two floating point types, `float32` and `float64`.
Bodner says "unless you have to be compatible with an existing format, use `float64`.
Floating point literals default to `float64`.
Also `float64` has more space for precision, and numbers in `float64` are more accurate.

Bodner suggests that you don’t use floats.
If you do use floats, Bodner suggests that you not use `==` or `!=` on floats, even though Go allows you to compare floats for equality.
If you divide a floating point number by zero, Go returns `+Inf` or `-Inf` depending on the sign of the number.

Note to self: I should learn how to avoid floating point numbers in my grading program.
I can think of one easy solution: define all numbers as `int` and allow truncation towards zero.
But I’m not convinced that this is the best way to handle things.

### Complex Types

I will skip this section since I won’t have any use for complex numbers.

### Strings and Runes

Go provides strings as a built-in type, and Go strings support Unicode.
You can include any Unicode character in a string literal.
You can compare strings with `==` and `!=`, and you can use `>`, `>=`, `<`, and `<=` on them as well.
You can concatenate strings using the `+` operator.

Strings are immutable. When you reassign the value of a string variable, you get a new string.

A rune represents a single code point.
The rune type is, as we mentioned earlier, an alias for `int32`.
You should use `rune` if you mean a `rune`.

## Explicit Type Conversion

Go does not automatically convert types.
Go puts the onus on the programmer to explicitly convert types.
In addition, you can’t use non-boolean types as booleans in Go.
Even stronger, you cannot convert other values to a boolean implicitly *or explicitly*.
Instead, you have to use comparison operators or write a function to return an appropriate boolean value.

## `var` versus `:=`

According to Bodner, “each declaration style communicates something about how the variable is used.”

Here’s the longest way to declare a variable in Go: `var x int = 10`.
In a case like this, you can leave off the `int` since it is clear that `10` is an integer literal.

Another form declares the type and automatically assigns the zero value for that type: `var x int` is equivalent to `var x int = 0` or `var x = 0`.

You can also declare multiple variables of one type: `var x, y int = 10, 20` or `var x, y int`, which is equivalent to `var x, y int = 0`.
And you can declare multiple variables of different types: `var x, y = 10, "hello"`.

Finally, if you have multiple variables, you can declare them in a declaration list:

```go
var (
    x   int
    y       = 20
    z   int = 30
    d, e    = 40, "hello"
    f, g string
)
```

`var` declarations can appear anywhere in a source file, and order does not matter.
That is, you don’t need to worry about the order of declaration and the order of use.

Within functions, you can use a short declaration format to replace `var` declarations that use type inference.
The following are equivalent:

```go
var x = 10
x := 10
```

You can also use `:=` to declare multiple variables at once.
As long as one variable on the lefthand side of `:=` is new, you can also use `:=` to assign values to existing variables.
(This is especially useful for things like `fileHandle, err := someFunction()`.

Bodner ends this section with advice:

+ `:=` is most common within functions, but there are some cases where you
  should avoid it.
+ Use the form `var x int` if you want to initialize a variable to its zero value.
  “This makes it clear that the zero value is intended.”
+ Use the long form to assign an untyped constant or literal to a variable if the default type for the constant or literal is not the type you want for the variable.
  That is, prefer `var x byte = 20` instead of `x := byte(20)`.
+ Be careful when you use `:=` with a mix of new and existing variables.
  Bodner recommends declaring new variables explicitly with `var` and then using `=` to assign values to the new and old variables.
+ Avoid declaring multiple variables on the same line except when you’re receiving multiple values from a function or the `something, ok` idiom (for maps, which we’ll see later in the book).
+ Avoid declaring true variables outside of functions in the package block.
  Package-level variables should be constants or quasi-constants.

## Using `const`

You can declare constants at package level or within a function.
However, in Go, you can only use constants to assign names to literals that the compiler can work out at compile time.
This means the following are the only valid assignments for constants:

+ Numeric literals
+ `true` and `false`
+ strings
+ runes
+ Built-in functions `complex`, `real`, `imag`, `len`, and `cap`
+ Expressions that use the preceding values and operators (e.g., 2 + 2)
+ `iota`, which is special, and we’ll discuss it later

You cannot use `const` to make other values immutable.
E.g., you cannot make an array, slice, map, or the field in a struct constant.

You can create typed or untyped constants.
The following two show the difference:

```go
// untyped
const x = 10
// you can use x in the following ways
var y int = x
var z float64 = x
var d byte = x

// typed: you cannot use x as a float or byte, for example
const typedX int = 10
```

## Unused Variables

You cannot declare a variable in Go and then not use it.
If you try, you will get an error when you try to run or compile the code.
However, the compiler does not catch unused package-level variables or unused constants.
(In the case of constants, this is definitely harmless.
As Bodner says, “if a constant isn’t used, it is simply not included in the compiled binary.”)

## Naming Variables and Constants

The rules for naming variables in Go are permissive.
You can use letters, numbers, the underscore, and even any Unicode letter or digit.
However, you should avoid using look-alike Unicode code points.

The idiom in Go is to use camel case for multiple-word identifiers.
Within a function, Go programmers prefer short (often one letter) names.
However, you should use more descriptive names for variables in the package block. Bodner also mentions that Go enforces rules about case in terms of accessibility.
He will talk more about this in Chapter 9.
