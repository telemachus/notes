# Effective C: Chapter 2 Objects, Functions, and Types

Seacord begins with this: “every type in C is either an *object* type or
a *function* type.” He defines *object* as “storage in which you can represent
values” and, more precisely, a “region of data storage in the execution
environment, the contents of which can represent values…when referenced, an
object can be interpreted as having a particular type.” Variables are examples
of objects that have a declared type.

A type tells us what kind of object its value represents. Types are essential
because the same collection of bits can represent different values when
interpreted as different types of object. Seacord gives the example of the bit
pattern `0X3f800000`. If we interpret that as a float, it represents the number
1. If we interpret that as an integer, it represents the number 1,065,353,216.

Functions are not objects, but they have types. The type a function is
determined by the types of its parameters and what type it returns.

Finally, pointers are addresses—locations in memory where objects or functions
are stored. The type of a pointer is derived from its function or object. The
function or object type is called a *referenced type*. If T is a function or
object type, we would call a pointer to that a *pointer to T*.

Seacord explains that C is call-by-value (aka, pass-by-value). This means that
functions receive copies of values not the arguments themselves. As a result,
you need to use pointers in order to swap two values in a C function. As Seacord
puts it, you use pointer to “simulate” call-by-reference (aka,
pass-by-reference).

In C, you need `*` and `&` to work with pointers.  You use `*` both to declare
pointer variables (e.g., `int *pa`) and to dereference (or get the value of)
a pointer (e.g., `int t = *pa`). You use `&` to take the address of a variable
(e.g., `swap(&a, &b)`. The `&` symbol is a unary operator that “generates
a pointer to its operand.”

## Scope

C has four scopes: file, block, function prototype, and function. If the
declaration is outside of a block or parameter list, then an object or function
has file scope. If a declaration is in a block or list of parameters, it has
block scope. If a declaration is in the list of parameter declarations in
a function prototype, then it has function prototype scope. (Function prototype
scope ends as soon as the declaration ends.) Finally, function scope is for
labels within a function.

Scopes can nest. You can have, for example, a block within a block. And
obviously every block is in some file’s scope. Inner scopes can access items
from outer scope, but outer scopes can’t access items from inner scopes.

You can use the same identifier name in multiple scopes. If you do, then the
inner variables hide the outer ones; that means that you cannot reach the outer
ones temporarily. This can be fine or bad, depending on what you’re doing and
how confusing it is for you and others.

## Storage Duration

Objects have a storage duration that determines their lifetime. There are four
storage durations: automatic, static, thread, and allocated. Seacord
distinguishes scope from lifetime in the following way:

+ Scope and lifetime are entirely different.
+ Scope applies to identifiers; lifetime applies to objects.
+ The scope of an identifier is the area of code where the object denoted by the
  identifier can be reached by its name.
+ The lifetime of an object is the time period during which the object exists.

Automatic duration begins when the block an object is declared in begins and
ends when the block ends. If the block is recursive, then a new object is
created each time, each with its own storage.

Static storage duration is for the entire lifetime of the program. Thus you can
use a `static` declaration instead of a global variable in some cases. For
example:

```c
void
increment(void)
{
	static unsigned int counter = 0;
	counter++;
	printf("%d ", counter);
}

int
main(void)
{
	for (int i = 0; i < 5; i++) {
		increment();
	}
	puts("");

	return EXIT_SUCCESS;
}
```

Note that you cannot initialize a static object with a variable; you must use
a constant value. (Constant values include literal constants, `enum` members,
and “the results of operators such as `alignof` or `sizeof`.”)

Thread storage is for concurrent program, and Seacord does not deal with
concurrent programming in this book. Allocated storage is for dynamically
allocated memory, and he covers it in Chapter 6.

## Object Types

### Boolean Types

You can declare objects as `_Bool` and store the values 0 and 1 in them. (0 is
false, and 1 is true.) Even better, you can include `stdbool.h` and then use the
integer constants `true` and `false` instead using the type `bool`. (Seacord
recommends using `bool` from `stdbool.h`.)

### Character Types

C defines `char`, `signed char` and `unsigned char`. If you are dealing with
character data, you should use `char`. (You can use `signed char` and `unsigned
char` for small integer values.) For characters beyond ASCII, you can use
`wchar_t` for wide characters. (I.e., characters that potentially take up more
than one byte of storage.)

### Numerical Types

#### Integer Types

C provides `signed char`, `short int`, `int`, `long int`, and `long long int`.
You can omit the keyword `int` when declaring any of the compounds of `int`.
C also provides an unsigned integer type for each of these: `unsigned char`,
`unsigned short int`, `unsigned int`, `unsigned long int`, and `unsigned long
long int`. You can only use the unsigned versions to represent values from zero
or higher. You can find out the minimum and maximum values using `limits.h`.
Finally, `stdint.h` and `inttypes.h` provide more specific integer types (e.g.,
`uint32_t`).

#### `enum` Types

An enumeration (`enum`) lets you define a type that assigns names (aka,
enumerators) to integer values with an enumerable set of constant values. There
are several ways to declare enumerations.

```c
enum day { sun, mon, tue, wed, thu, fri, sat };
enum cardinal_points { north = 0, east = 90, south = 180, west = 270 };
enum months { jan = 1, feb, mar, apr, may, jun, jul, aug, sep, oct, nov, dec };
```

In the first case, the first value will automatically get a value of 0, and each
value is incremented by 1 (i.e., 0, 1, 2, …). You can also assign values to each
enumerator explicitly, as in the second case. Finally, you can specify only the
first value, as in the third case. In such a declaration, all other enumerators
will be incremented by 1.

#### Floating-point Types

C provides three floating-point types: `float`, `double`, and `long double`.

#### `void` Types

Seacord says that `void` is “rather strange.” By itself, the keyword `void`
means “cannot hold any value.” But the derived type `void *` means that the
pointer can reference any object. He will talk more about derived types later in
the chapter.

### Function Types

Function types are derived from their return type and the number and types of
its parameters. The return type of a function cannot be an array type.

When you declare a function, you have several options in terms of the
parameters. You can leave them blank altogether, you can specify types only, or
you can specify types and names.

```
int f(void);
int *fip();
void g(int i, int j);
void h(int, int);
```

Seacord recommends avoiding the second and fourth examples. The second in C++
means a function that accepts no arguments, but in C it declares a function that
accepts any number of arguments of any type. `g` and `h` both declare types, but
Seacord prefers `g` because giving identifiers makes code more
“self-documenting.” (I’m not sure how, frankly.) All the examples above are
“function prototypes.” They give the name, return type, and (possibly) parameter
lists of the function, but they don’t define its body. A function’s definition
gives the actual body or implementation of the function.

```c
int
max(int a, int b)
{
    return a > b ? a : b;
}
```

### Derived Types

#### Pointer Types

A pointer derives its type from the function or object that it points to (aka,
the referenced type). “A pointer provides a reference to an entity of the
referenced type.” You use `&` to take the address of something and `*` as an
indirection operator to get from address to value or to declare a pointer.

```c
int i = 17;
int *ip = &i;
int j = *ip; // j = 17
```

You can also use the `&` operator on the result of the `*` operator: `ip = &*ip;`.

The `*` operator “converts a pointer to a type into a value of that type. It
denotes *indirection* and operates only pointers.” 

#### Arrays

An array is a contiguously allocated sequence of object of the same type. The
type of an array is their element type plus the number of elements in the array.
You use `[]` to index into an array and get one element from it.
