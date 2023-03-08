# The Python Apprentice, Robert Smallshire and Austin Bingham

## Overview

Python is a strongly but dynamically typed language.  Since Python is strongly
typed, everything in the language has a specific type.  Since Python is
dynamically typed, the type of an object can change over the course of
a program, and the interpreter does not check types like, e.g., the Go
compiler does.  Python uses duck-typing, and whether or not an object does
what it needs to do is only worked out at runtime.

### Scalar data types: integers, floats, `None`, `bool`

+ `int`: signed, (virtually) unlimited precision integers; these numbers are
  limited only by machine memory; you can express them in binary with an `0b`
  prefix, octal with an `0o` prefix, or hexadecimal with an `0x` prefix; you
  can use `int()` to convert other numeric types to integers; the `int()`
  constructor always rounds towards zero; you can also use `int()` to convert
  strings to integers, and you can specify a base when you convert strings to
  integers; e.g., `int('11', 2)` returns `3`
+ `float`: IEEE 754 floating-point numbers; Python floating point numbers have
  53 bits of precision; they provide between 15 and 16 significant digits; any
  number with a decimal point is a float; you can also use scientific notation
  for large numbers or small numbers; e.g., `3e8` or `1.616e-35`; you can use
  `float` to convert another numeric value or a string to a float; if you
  calculate a value with a mixture of integers and floating point numbers, the
  result is a `float`
+ `None`: special, singular null value; `None` often represents the absence of
  a value; you can assign the value `None` to a variable; you test for `None`
  using `is`; e.g., `a is None`; do not test for `None` with `==`
+ `bool`: Boolean values, `True` or `False`; you can covert other types to
  Boolean values using `bool()`; `bool()` is especially useful for simplifying
  some assignments and return statements; happily, `pylint` will tell you when
  you missed an opportunity to use `bool()`; according to `bool()`, every
  number except 0 and 0.0 is true, every string except "" is true, and every
  list except `[]` is true

## Relational operators

In Python, relational operators test and compare equivalence not identity.
That is, they don't compare or test whether two items are identical, but
whether you can use one in place of another.  They will discuss "object
equivalence" more later.

## Control flow: if-statements and while-loops

Python uses `if <boolean>:` to introduce a branch.  You can add an `else:`
after any `if`.  You can also have a multi-branch conditional using `if` ...
`elif` ... [`elif` ...] `else`.

You can use `while <boolean>:` to introduce a loop.  You can use `break`
inside of the block introduced by `while` to exit a loop early.  If there are
several loops, `break` exits the innermost one.

## Errata

Page 5: “for the most part Python is as flexible and adaptable as *[m]any* modern
programming language”
