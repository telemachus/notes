# Chapter 2: Types and values

Lua is dynamically typed: type definitions do not occur since values carry
their own type and can be changed through reassignment at any time.

There are eight basic types in Lua:

1. nil
2. boolean
3. number
4. string
5. userdata
6. function
7. thread
8. table

The function `type` returns a string type value for any item.

Variables do not have predefined types, nor is the type of a variable static.

## Nil

Nil is a type with a single value, `nil`. The point of this type (and value) is
to be different from any other value. You can use `nil` to check for lack of
a proper, valid, or useful value. All global variables have a value of `nil`
before they are first assigned something. You can also assign `nil` to any
variable as a way of destroying it within a longer running program for memory
reclamation.

## Booleans

There are exactly two boolean values: `true` and `false`. However, any value
can be used in a condition in Lua code. For the purpose of conditions, `nil`
and `false` are *falsey* values, and everything else is *truthy*. (NB: Both
0 and an empty string are true.)

## Numbers

Up until Lua 5.3, Lua had only real, floating point numbers. As of 5.3, Lua has
both integers and floats. For the most part, I am unlikely ever to care very
much about when Lua converts one into the other.

Numeric constants can be written in Lua with an optional decimal portion and an
optional decimal exponent. Lua can also handle hexadecimal constants.

## Strings

Lua strings are sequences of characters that can include embedded zeros as well
as Unicode characters. However, the standard library does not have support for
anything beyond UTF-8.

Strings in Lua are immutable values. When you perform substitutions, additions,
or any kind of change on a string, a new string value is created. But since
strings are automatically memory managed, the programmer does not need to worry
about allocation or deallocation of strings. Strings in Lua can be **very**
large without problems. You can get the length of a string by prefixing it with
the length operator, `#`.

### Literal strings

Strings can be marked with double or single quotes, with no difference in
meaning except that within each type, the other type need not be escaped. Lua
can interpolate various C-like escape sequences (e.g. `\n`) within literal
strings. For hard to type characters, you can also give Lua sequences like
`\ddd` or `\x\hh` in decimal or hexademical form.

### Long strings

For multi-line strings, Lua provides a shortcut: matching double square
braces. When the first character of such a string is a newline, this newline is
ignored. Otherwise, everything is taken literally. Long strings of this form do
not interpret escape sequences. E.g.

```
page = [[
<html>
    <head>
    </head>
</html>
]]
```

If needed, you can add any number of equal signs before the opening delimiter.
In this way, you can have long strings with double square braces inside them
without prematurely ending the string. (The same thing can be done for comments
as well.)

If you need to have a literal string of arbitrary data, written as decimal or
hexadecimal constants, long strings are no good. Instead, you should use
literal strings with the special escape sequence `\z`. This escape causes Lua
to ignore all space up until the next non-space character. E.g.

```
data = "\x00\x03\x01\x03\x04\x05\x06\x07\z
        \x09\x09\x0A\x0B\x0C\x0D\x0E\x0F"
```

### Coercions

Lua will automatically coerce strings into numbers or numbers into strings in
various situations. You should avoid allowing this or relying on it.

## Tables

Tables are Lua's *only* collection data type. This can be tricky sometimes, but
mostly it works out well. Strictly speaking, all tables are associative arrays
(i.e. hashes). A table can be indexed by numbers, strings or in fact any value
in the language except `nil`.

Tables to a lot in Lua. They provide all standard data collection types:
arrays, lists, hashes, sets, records, etc. They also provide the structure that
Lua uses to handle the importing of other code (i.e. modules).

Strictly speaking, tables are not values or variables. They are dynamically
allocated objects. To create a table, you use the constructor `{}`. Tables are
anonymous as well, so each variable that points at a table is a distinct
reference.

```
a = {}
a["x"] = 10
b = a
print(b["x"])    -- 10
b["x"] = 20
print(a["x"])    -- 20
a = nil          -- a has no reference, but b still refers to the table
b = nil          -- Now there is no reference to the table which can be garbage
                 -- collected now.
```

Like global variables, table fields evaluate to `nil` before they are
initialized. You can also assign `nil` to a field to destroy an item from
a table.

To represent records in a table, you can use the field name as an index. Lua
provides syntactic sugar for this in the form of dot notation.

```
a.x = 10 -- This is the same as a["x"] = 10
```

Lua doesn't care which form you use, but according to RI, they suggest
different intentions to human readers. The dot notation, in his view, suggests
clear use of the table as a record with fixed, predefined keys. The string form
suggests that any string might be a key, but that for some reason we're using
this string as a key. (I don't know that I buy this, but ok.)

Be careful not to confuse `a.x` and `a[x]`. The first indexes the table `a` by
the string value `"x"`, while the second indexes the table `a` by whatever
value is held in *the variable* `x`.

When using a table as an array, you can start your indexes wherever you like,
but you should start with 1. This is because Lua's standard library for tables
assumes a starting-point of 1.

For proper sequences (i.e. tables indexed by numbers without any holes in their
ordering), the length operator `#` can be used to get the length of the table.
However, if the table is sparse (i.e. has holes in its sequence), then `#` will
not work correctly. It will simply stop at the first `nil` value in the
sequence. This may or may not be what you want.

## Functions

Functions are full, first-class values in Lua. You can store them in variables,
pass them to other functions as arguments, and return functions as results from
functions. This makes Lua highly flexible, and it also means that Lua provides
a great deal of support for functional programming style.

Lua can call functions written in Lua as well as functions written in C. Often
you will want to use C for performance-demanding code or to get at things not
easily reached from Lua itself. Lua's standard libraries are all written in C.

## Userdata

The userdata type allows arbitrary C data to be stored in Lua variables.
Userdata has only two predefined operations in Lua: assignment and equality
testing. You use userdata to represent new types within Lua libraries or
applications. As an example, the I/O library that comes with Lua uses userdata
to model open files.

## Threads

Threads are important for coroutines. More on both of these later.
