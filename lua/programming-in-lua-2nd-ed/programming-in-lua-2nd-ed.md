# Notes on *Programming in Lua* (2nd ed.)

(I'm reading the second edition again because Neovim targets Lua 5.1 via
LuaJIT.)

## Chapter 1: Getting Started

### Chunks

A *chunk* is a sequence of commands or statements that Lua executes.  A chunk
can be a single line or an entire file, and the Lua interpreter can handle
chunks of "several megabytes" (4).

Lua does not require any separators between statements, but you may add
semicolons if you want.  The following are equivalent and valid.

```lua
a = 1
b = a*2

a = 1;
b = a*2;

a = 1; b = a*2;

a = 1 b = a*2
```

Note the last one in particular.  You don't need semicolons or newlines.
Ierusalimschy calls the last version "ugly, but valid" (4).

Chunks can be one statement, multiple statements, or a mix of statements and
function definitions.  (It turns out that function definitions are actually
assignments, but more on this later.)

You can test code by using the Lua interpreter in interactive mode.  If you
run `lua`, the interpreter starts up in REPL mode.  You can also start as `lua
-i filename` to have the interpreter first read a file and then start in REPL
mode.  Or from within the REPL, you can call `dofile('filename')` if you want
to load a file from within the REPL.

### Some Lexical Considerations

Variable names in Lua can be "any string of letters, digits, and underscores,
not beginning with a digit" (5).  However, you should avoid using identifiers
that begin with an underscore.  The Lua developers reserve these for special
cases of their own.  By convention, many Lua developers use `_` for dummy
variables.  You *can* use identifiers with UTF8 characters, but you shouldn't.

The following words are reserved.

```
and         break       do      else        elseif
end         false       for     function    if
in          local       nil     not         or
repeat      return      then    true        until
while
```

Comments come in two forms.  Single-line comments start with a double hyphen
(`--`).  Block comments begin with `--[[` and end with `--]]`.  (You can make
block comments more flexible by adding any number of equal signs between the
two brackets.  That is, `--[===[` begins a block comment that ends with
`--]===]`.  This can be useful if you have nested block comments, but you
should probably avoid that.)

### Global Variables

In Lua, variables are global by default, and you don't need to initialize
a variable before you use it.  If you run `print(a)` before you initialize
`a`, then Lua gives no error.  `a` is simply `nil`, and Lua acts as though you
asked it to `print(nil)`.  (This is valid and prints the string 'nil'.)  You
can delete a global variable by assigning `nil` to it, but you rarely need to
do this with global variables.  (Ierusalimschy says "if your variable is going
to have a short life, you should use a local variable" (6).)

### The Stand-Alone Interpreter

If the first line of a file begins with `#`, the Lua interpreter ignores the
line.  (This way, you can put a shebang line there and run the file as
a command.)

The stand-alone interpreter has several options.  `-e` allows you to feed code
to execute immediately.  `-i` says enter interactive mode when all other
commands are done.  `-l` loads a library.  You can also change the prompt in
the REPL using the global `_PROMPT`.  E.g., `lua -i -e "_PROMPT=' lua> '"`.
While in interactive mode, you can get the value of an expression by prefixing
`=` to the expression.  E.g., `= math.sin(3)`.  You can also specify a file in
the environment variable `LUA_INIT` and fill that fill with Lua code.  This
allows you to do all sorts of things when the stand-alone interpreter runs.
Finally, when you run a script via the `lua` command, the arguments to the
script are available in the global variable `args`.  As of Lua 5.1, you can
also get the script's arguments from using the vararg syntax (`...`), but more
on this later.

## Chapter 2: Types and Values

Lua is dynamically typed.  You don't need to declare a type, and one variable
can change types in the course of a program.  Lua has eight basic types: nil,
boolean, number, string, userdata, function, thread, and table.  You can use
the builtin function `type` to get the string name of something's type.  E.g.,
`type(print)` returns 'function', and `type(true)` returns 'boolean'.  By
definition, `type(type(x))` always returns 'string'.

### Nil

Nil as a type has a single value, `nil`.  Uninitialized global variables are
`nil`, and you can delete a variable by assigning `nil` to it.  Nil indicates
that there is no useful value in a variable.

### Booleans

Lua provides two booleans, `true` and `false`.  But conditions evaluate
`false` and `nil` as false and everything else as true.  Thus, both zero and
the empty string are true in conditional tests.

### Numbers

Lua 5.1 has only floating-point numbers.  Even what look like integers are
actually floats.  You can compile Lua 5.1 to use other kinds of numbers, but
by default Lua uses doubles to represent all numbers.

### Strings

You should be able to put any UTF8 character into a Lua string.  (However,
I don't think that the string library handles UTF8 characters for all its
operations.)

Lua strings are immutable.  You can create a new string to modify an old one,
but you cannot modify an existing string.

You can use single or double quotes to delimit strings.  You can also use
C-like escape sequences inside Lua strings, whether the string is in single or
double quotes.

+ `\a` bell
+ `\b` back space
+ `\f` form feed
+ `\n` newline
+ `\r` cariage return
+ `\t` (horizontal) tab
+ `\v` vertical tab
+ `\\` backslash
+ `\"` double quote
+ `\'` single quote

You can also specify a character as `\ddd` where 'ddd' is a series of (up to)
three decimal digits.  (Please don't do this.)

You can easily create multiline strings using `[[` and `]]`.  Inside the
brackets, newlines do not need to be escaped.  But note that if the first
character of such a string is a newline, Lua ignores it.  (This turns out to
be convenient most of the time.)

If you want to put code that has brackets into such a string, you can simply
add any number of equal signs to the opening and closing brackets.  The same
goes for comments.  You can use `[===[` paired with `]===]` and `--[=[` paired
with `--]=]`.  Again, you should save this for extreme circumstances.

Lua 5.1 performs automatic coercion of strings to numbers and numbers to
strings in some cases, but you should avoid this.  If you want to convert
a string to a number, do it explicitly with the built-in `tonumber` function.
(This function returns `nil` if the conversion fails.)  If you want to convert
a number to a string, use the built-in function `tostring`.  (You can also
concatenate a number with an empty string to turn it into a string.)

You can get the length of a string using the prefix operator `#`.

### Tables

Lua provides only one built-in collection type: tables.  They can be used as
lists or dictionaries, but inside they are *always* dictionaries.  They are
very flexible, but also can be tricky and confusing.

Tables have no fixed size.  They grow as needed.

Ierusalimschy describes tables as "neither values nor variables; they are
*objects*" (14, his emphasis).  They are dynamically allocated objects, and
programs manipulate "only references (or pointers)" to tables (14).  You
cannot *declare* a table; you create tables by means of a constructor
expression.  The simplest form of constructor expression is simply `a = {}`.  That
creates a table and binds its reference to `a`.  If you later assign nil to
`a`, then Lua can garbage collect that table, assuming that there are no other
references to the table.

The table itself, however, is anonymous.  There's not a fixed reference
between a variable and a table.

```lua
a = {}
a['x'] = 10
b = a           -- a and b refer to the same table
print(b['x'])   -- 10
b['x'] = 20
print(a['x'])   -- 20
a = nil         -- b still refers to the table
b = nil         -- no references left to the table
```

You can use numeric indexes to create list-like tables in Lua.

```lua
for i=1, 100 do
    a[i] = i*2
end
print(a[9])     -- 18
print(#a)       -- 100
```

However, tables are never actually lists.  You can add dictionary-like
(key/value) entries into any table, even one with numeric indices.  In
addition, you can break `#` by creating a sparse array.  (In such cases, you
can use `table.maxn`, but I'm not sure whether this is wise.)  Thus, you need
to maintain discipline when using tables in Lua.

Lua provides some syntatic sugar for when you want to use tables as
dictionaries.  Instead of `table['key']`, you can use `table.key`.  However,
this only works if the key limits itself to certain characters.  More details
on this later.

### Functions

Lua functions can be stored in variables, passed as arguments to other
functions, and returned as results from function.  You can redefine functions
easily in Lua, and you can also erase them (by assigning them to `nil`) for
security.  Since Lua functions support nested functions with lexical scoping,
you can also write Lua in a functional style.

The Lua interpreter can call functions written in Lua and functions written in
C.  Most of the standard library (all of it?) are written in C.  Programs can
also define their own functions in C as needed.

### Userdata and Threads

Userdata allows C data to be stored in Lua variables.  Lua itself has no
predefined operations for userdata except assignment and equality testing.
Applications or libraries create new types in C and share them with Lua via
userdata.  (For example, the standard I/O library uses userdata to represent
files in Lua.)

Threads will be discussed in Chapter 9 when Ierusalimschy covers coroutines.

## Chapter 3: Expressions

Expressions denote values.  In Lua, expressions include numeric constants
(e.g., 49 or 12.3), string literals (e.g., 'hello' or "world"), variables,
unary and binary operations (e.g., -1 or 1 + 1), and function calls.  Less
obviously, function definitions and table constructors are also expressions in
Lua.

### Arithmetic Operators

Lua provides the usual operators for arithmetic, plus `^` for exponentiation.
The modulo operator in Lua always gives the result the same sign as the second
argument.  It is equivalent to the following formula.

```
a % b =- a - floor(a/b)*b
```

(By comparison, in C and Go the result of modulo is negative if the first
operator is negative.  Python works like Lua.)

### Relational Operators

Lua has the following relational operators.  They return true or false.

+ <
+ >
+ <=
+ >=
+ ==
+ ~= (i.e., not equal)

You can apply the `==` and `~=` operators to any two values.  If the values
differ in type, then Lua immediately considers them not equal.  `nil` is equal
only to itself.  Lua compares tables, userdata, and functions by reference.
Consider the following code.

```lua
a = {}; a.x = 1; a.y = 0
b = {}; b.x = 1; b.y = 0
c = a
a == c -- true
a == b -- false
a ~= b -- true
```

You can only apply the order operators to two strings or two numbers.  Strings
are compared in alphabetical order following the locale setting.  Lua raises
an error if you try to compare a string with a number in an order comparison.
(This is somewhat surprising since Lua performs automatic coercion elsewhere,
but I suppose that it's for our own good.)

### Logical Operators

Lua uses words rather than symbols as logical operators: `and`, `or`, and
`not`.  These logical operators consider `false` and `nil` false.  Everything
else is true.

Both `and` and `or` return one of their arguments rather than a literal
boolean value.  `And` returns its first argument if the first argument is
false; otherwise, `and` returns its second argument regardless of the truth
value of that argument.  `Or` returns its first argument if the first argument
is true; otherwise, `or` returns its second argument regardless of its truth
value.

Both `and` and `or` evaluate their arguments lazily.  That is, they will only
evaluate their second argument if they must.  You can use this to protect
yourself.  For example, `x = require('x'); if x and x.function then...end`.

Here are two common Lua idioms.

```lua
x = x or v          -- equivalent to if not x then x = v end
(a and b) or c      -- equivalent to ternary: a ? b : c
a and b or c        -- same as (a and b) or c because of precedence
```

Note that these idioms have some limitations.  The first only works if
x itself has not been deliberate assigned false.  The second only works if
b is truthy.  Still, these can come in handy.  You can use the first to assign
x a default value when it is not set.  And you can use the second, e.g.,  to
get the maximum of two numbers: `max = (x > y) and x or y`.

The `not` operator always returns `true` or `false`.  If you want to coerce
something to an actual boolean value (insteady of truthy evaluation), you can
use `not`.

### Concatenation

Lua performs string concatenation using `..`.  If an operand is a string, Lua
coerces that operand to a string and then concatenates it.  The return value is
a new string, and the operands are not changed.  (Remember: strings in Lua are
immutable values.)

### Precedence

Lua assigns the following operator precedence.

```lua
^
not # - -- Unary -
* / %
..
< > <= >= ~= ==
and
or
```

The exponentiation operator and the concatenation operators are right
associative.  All other binary operators are left associative.

As Ierusalimachy says, "When in doubt, always use explicit parentheses.  It is
easier than looking it up in the manual, and you will probably have the same
doubt when you read the code again" (22).

### Table Constructors

Constructors are expressions that create and initialize tables.  Here are some
examples.

```lua
x = {}              -- Create an empty table, and assign its reference to x.
days = {            -- Create an empty table; add the days of the week to it.
    'Sunday',       -- 'Sunday' is days[1], 'Monday' is days[2], etc.
    'Monday',
    'Tuesday',
    'Wednesday',
    'Thursday',
    'Friday',
    'Saturday'
}

a = {x=10, y=20}    -- Equivalent to a = {}; a.x = 10; a.y = 20

polyline = {        -- We can mix dictionary-style and list-style tables.
    color='blue',   -- These are keys in a dictionary part of this table.
    thickness=2,
    npoints=4,
    {x=0, y=0},     -- These are subtables, stored in positive indices 1..n.
    {x=-10, y=1},
    {x=0, y=1}
}
```

You can't use the `a.x = y` format for all values of `x`.  In those cases, you
can always do `a['x'] = y` instead.  E.g., `a['+'] = add` (where, by the way,
`add` can be a variable that refers to a function).

You can always include a trailing comma in all table constructors.  Also, if
you prefer, you can use a semicolon instead of a comma to separate elements in
a constructor.  (Ierusalimschy uses a semicolon to mark the break between an
array part of a table and a dictionary part of a table.  This, however, is
just a personal preference.)

## Chapter 4: Statements

Ierusalimschy says that "Lua supports an almost conventional set of
statements, similar to those in C or Pascal" (27).  The normal statements are
assignment, control structures, and procedure calls.  The less common
statements include multiple assignments and local variable declarations.

### Assignment

Assignment is the normal way to change the value of a variable or table entry.
Lua also allows for multiple assignment, where a list of values is assigned to
a list of variables in one setep.  When Lua performs multiple assignment, it
evaluates all values on the right-hand side first, and then it performs the
assignments.  Thus, you can use multiple assignment to swap values.

Lua silently adjust the number of values to the number of variables.  If the
list of values is shorter than the list of variables, the extra variables are
assigned `nil`.  If the list of values is longer than the list of variables,
the extra values are silently discarded.  (Note, however, that the extra
values are *still* evaluated, right?)

One other warning: you cannot assign multiple variables using a single value
on the right-hand side. `a, b, c = 0` does *not* assign zero to a, b, and c.
Instead, it assigns zero to a, and b and c are assigned `nil`.

### Local Variables and Blocks

By default, variables in Lua are global, but you can create local variables
with the `local` statement.  Local variables have block scope.  They are
scoped to the block where they are declared.  A block is the body of a control
structure, the body of a function or a chunk (a file or string where
a variable is declared).  (Note: local variables created in an `if` block in Lua do
not exist in the corresponding `else` block.)  You can also create a block by
wrapping some Lua code in a `do...end` block.

Lua treats the delcaration of a local variable as a statement.  Therefore, you
can put a local-variable declaration anywhere you can put a statement.  Like
Go, Lua encourages declaring variables only where they are needed or used.  It
is not bad practice (in Lua) to declare a local variable in the middle of
a block.  Finally, `local foo = foo` is a common idiom in Lua.  Such a line
creates a local version of `foo`. This is good if you want to preserve the
original value of `foo` (say in `old_foo` earlier), and (supposedly?) access
to a local variable is faster than access to a global variable.

### Control Structures

Lua has a small set of control structures: `if`, `while`, `repeat`, and `for`.
You must use `end` to terminate `if`, `while`, and `for` structures; you use
`until` to terminate a `repeat` structure.  There are two very different
versions of the `for` structure, though they are both named "for".  See
details below.

You can use any value in the condition expression of a control stucture.  Lua
treats `false` and `nil` as false and everything else as `true`.

#### `if...then...[elseif...then...][else...]end`

An `if` statement tests its condition and executes its "then" or "else"
section accordingly.  The "else" section is optional.  You can use as many
`elseif` sections as you like.  Lua has no switch staement, so you should
expect to use and see chains of `if...elseif`.

#### `while...end`

Lua tests the condition and runs as long as the result is truthy.  The body of
a while condition will never be run if the test is initially false.

#### `repeat...until`

Lua repeats the code in the body until the condition is true.  The body will
run at least one time since the test is not done until after the body is run.

If you create a local variable inside a repeat block, you can still use that
variable in the condition.

```lua
local sqr = x/2
repeat
    sqr = (sqr + x/sqr)/2
    local error = math.abs(sqr^2 - 2)
until error < x/10000
```

#### Numeric `for`

A numeric `for` works as follows.

```lua
for var = exp1, exp2, exp3 do
    -- something
end
```

The loop will execute *something* from exp1 to exp2, using exp3 as a step
value.  If you omit exp3, then Lua assumes a step value of 1.

Some notes and tips about numeric for loops in Lua.

1. If you want a loop with (virtually) no upper limit, use `math.huge` as
   exp2. You can then use `break` to get out of the loop.
2. The three expressions are only called once, before the loop begins.  If you
   put a function as one of the expressions, it is not called every time the
   loop runs.
3. The control variable is automatically local to the `for` loop.  If you need
   to capture a value inside a `for` loop for later use, you need to declare
   it before the loop begins.
4. Do not change the value of the control variable within the loop.  (From
   what I can tell, if you do so, the result is undefined.)  If you need to
   stop the loop, use `break`.

#### Generic `for`

The generic `for` loop runs over all the values returned by an iterator
function.  For example, `pairs` returns the keys and values in a table.

```lua
for k, v in pairs(t) do
    -- something
end
```

There are several iterators in Lua's standard library, and you can create your
own.  (More on how to do this in Chapter 7.)

Just as with a numeric `for` loop, loop variables are automatically local to
the loop body of a generic `for` loop and you should not assign anything to
those variables.

#### `break` and `return`

You can use `break` and `return` to jump out of blocks.

The `break` statement finishes a loop.  It ends the innermost loop (`for`,
`repeat`, or `while`) that contains it.  You cannot use `break` outside of one
of those loop structures.

A `return` statement causes the end of a function.  You can optionally return
values with `return`, but you can also use it to end a function's execution
early.  (There is an implicit empty return at the end of every Lua function.
So you don't need to write `return` to exit a function normally.)

You can only use `break` or `return` as the last statment of a block.
However, if you want to include one early (say in a function that you are
debugging), you can wrap an early `return` in a `do...end` block.

```lua
function foo()
    return          -- Syntax error: not allowed
    do
        return      -- This one is fine.
    end
    -- <other code that is temporarily disabled>
end
```

## Chapter 5: Functions

Functions can be statements or expressions.  As statements, they perform some
action or series of actions.  (Some languages call this type of function
a *procedure* or *subroutine*.)  As expressions, they do some computation and
then return one or more values.

The zero or more arguments to a function go inside parentheses.  In general
the parentheses are required when you call a function, even if there are no
arguments.  However, if an argument takes one argument, and the argument is
a string literal or a table constructor, then parentheses are optional.

Lua also provides syntactic sugar for function calls in object-oriented code.
You can write `o:foo(x)` in place of `o.foo(o, x)`.

Lua programs can call functions defined in Lua, C, or any other language used
by a host application.

Here is how function definition looks in Lua.

```lua
function sumTable(t)
    local sum = 0
    for _, n in ipairs(t) do
        sum = sum + n
    end
    return sum
end
```

How does Lua handle parameters?  First, parameters are local to a function by
default.  (You cant access `t` *by that name* outside of the `sum_table`
function.)  Also, you can call a function with a different number o arguments
than the number of defined parameters.  Lua adjusts the number of arguments to
the number of parameters, just as it adjusts return values.  Extra arguments
are thrown away (but evaluated?), and missing parameters are assigned the
value nil.

You can use Lua's argument adjustment to handle defaults for functions.  For
example, in the following example `n` defaults to one if you pass no argument.

```lua
function incrementCount(n)
    n = n or 1
    count = count + n
end
```

Lua functions can return multiple values, and many functions in the standard
library do so.  For example, `string.find` returns the start and end of
a pattern in a string.

```lua
s, e = string.find("hello Lua lovers", "Lua") -- s = 7, and e = 9
```

However, Lua will silently adjust "the number of results from a function to
the circumstances of the call" (37).  The details here are more complicated
than you might think.  (See [this article][lua-multivals] for details.)  In
a nutshell, however, you get all the results from a function only if the
function call is the last or the only expression in a list of expressions.  If
the function appears anywhere else in a list of expressions, then the function
call yields one result.  (That result may be `nil` if the function returns
nothing.)  If a function call appears as an expression inside another function
call or in a table constructor, the same rules apply.  If the inner function
call is the last or only expression, then it returns all of its results.
Otherwise, it returns a single value, which can be `nil`.  Note, however, that
if the outer function expects a specific number of arguments, Lua adjusts the
number of results from the inner function accordingly.  Inside a function,
a return statement like `return f()` returns all the values that `f` returns.
However, of course, those returns may be adjusted depending on the context.
Finally, you can force any call to return exactly one result by wrapping it in
an extra set of parentheses.

Lua provides a function `unpack` to return all the elements of a table
starting at index 1.  This is useful when you have a table, but you want to
call a function that expects individual arguments.

[lua-multivals]: https://benaiah.me/posts/everything-you-didnt-want-to-know-about-lua-multivals

### Variable Number of Arguments

Lua allows functions that receives a variable number of arguments.  An obvious
example is `print`.  You can also define your own functions that get
a variable number of arguments.  The basic tool here is `...`, which stands
for a variable number of arguments.  (The name for `...` is *varargs*.)

```lua
function sum(...)
    local total = 0
    for _, n in ipairs({...}) do
        total = total + n
    end
    return total
end
```

The expression `{...}` creates a table with all the arguments passed to the
call.  You can also use multiple assignment with `...`.  E.g., `local a, b,
c = ...`.  Here's an example of where varargs come in handy.

```lua
function fwrite(fmt, ...)
    return io.write(string.format(fmt, ...))
end
```

This allows us to format a string and print it all at once.  Note that you can
place any number of fixed parameters before varargs.

You can use `select` to help handle cases where varargs contains valid nils.
Here's how `select` works.

+ `select(number, ...)` returns all arguments starting with index position
  number.
+ `select('#', ...)` returns the total number of arguments received in
  varargs.

Here are examples.

```lua
print(select(2, 1, 2, 3))           -- 2    3
print(select(1, 1, 2, 3))           -- 1    2   3
print(select(3, 1, 2, 3))           -- 3
print(select('#', 1, 2, nil, 4, 5)  -- 5
```

Lua does not directly support named arguments, but you can easily imitate them
using a table as the sole argument to a function.

```lua
-- This is not valid in Lua, but we want something like it.
-- reame(old='temp.lua', new='temp1.lua')

function rename(t)
    return os.rename(t.old, t.new)
end

-- We call the function as follows.  The two items in the table can appear in
-- either order.
rename({old='temp.lua', new='temp1.lua'})
```

## Chapter 6: More About Functions

Ierusalimschy likes to tell us that Lua functions "are first-class values with
proper lexical scoping" (45).  In this chapter, he explains in some detail
what that means, and he gives examples of more complex uses of functions in
Lua.

What is a *first-class value*?  A first-class value, like numbers or
strings, can be stored in variables (global and local), stored in tables,
passed as arguments to functions, and returned by functions.

What does *lexical-scoping* mean?  Lexical scoping means that functions can
access variables from their enclosing functions.  This means that you can use
Lua functions to make closures and that Lua can handle many functional
techniques.

Strictly speaking, Lua functions are anonymous.  The name of a function is
actually just a variable that refers to the function.  We can add other
references and change the reference of given variables.  Ierusalimschy gives
the following example.

```lua
a = {p = print}             -- a.p is now a new way to call print
a.p("Hello, world!")        -- Hello, world!
print = math.sin            -- print is now math.sin
a.p(print(1))               -- 0.841470
sin = a.p                   -- sin is now print
sin(10, 20)                 -- 10   20
```

Going further, all function declarations are actually syntactic sugar for
assignments.  The following two are equivalent.  All function definitions are
statements.

```lua
function foo(x) return 2*x end

foo = function(x) return 2*x end
```

Lua also supports anonymous functions.  For example, the `table.sort` function
takes two parameters: a table and a function to sort the table. You can use an
anonymous function if you like.  The following are equivalent.

```lua
function sortnums(a, b) return a > b end

table.sort(t, sortnums)

table.sort(t, function(a, b) return a > b end)
```

### Closures

You can define functions inside of other functions in Lua, and when you do so,
the inner function has access to all the variables in the enclosing functions.
(This is called *lexical scoping*.)  Consider the following example.

```lua
function sortbygrade(names, grades)
    table.sort(names, function(n1, n2)
        return grades[n1] > grades[n2]
    end)
end
```

The anonymous function can access the table `grades`, but that table is
neither a global nor local to the anonymous function.  In Lua, these are
called *non-local variables* or *upvalues*.

Now consider another example.

```lua
function newCounter()
    local i = 0
    return function()
                i = i+1
                return i
            end
end

c1 = newCounter()
print(c1())         -- 1
print(c1())         -- 2
```

By the time the anonymous funnction is called, `i` appears to have gone out of
scope.  However, the anonymous function is a closure.  As Ierusalimschy
explains, “a closure is a function plus all it needs to access non-local
variables correctly” (48).  Every time that you call `newCounter`, you get
a new value of `i`.  The anonymous function then keeps track of the value of
*its* `i` throughout its lifetime.

Because you can store variables easily in Lua, you can redefine functions
based on their current implementation.  For example, we can use `math.sin` to
create a new version of `sin`.

```lua
do
    local oldSin = math.sin
    local k = math.pi/180
    math.sin = function(x)
        return oldSin(x*k)
    end
end
```

Notice also that `oldSin` is now stored in a private variable. You can *only*
access it through the wrapper function.  You can use this same technique to
restrict access to certain functions.  For example, consider the following.

```lua
do
    local oldOpen = io.open
    local accessOK = function(filename, mode)
        -- Check access here somehow.
    end
    io.open = function(filename, mode)
        if accessOK(filename, mode) then
            return oldOpen(filename, mode)
        else
            return nil, "access denied"
        end
    end
end
```

### Non-Global Functions

In Lua, functions do not have to be global.  You can also store functions in
local variables and in table fields.  Lua libraries (i.e., modules) use tables
to store their functions.  Here's an example.

```lua
local M = {}

M.foo = function(x, y) return x + y end

return M
```

There are several ways to define functions for such libraries.

```.
local M = {}
M.foo = function...end

local M = {
    foo = function...end
    bar = function...end
}

local M = {}
function M.foo(x, y)...end
```

You can also declare local functions in several ways.

```lua
local f = function(...)...end

local function f(...)...end
```

However, if you want to create a recursive local function, you must declare
its name before you define the function.  Compare the two below.

```lua
-- This won't work.
local fact = function(n)
    if n == 0 or n == 1 then
        return 1
    else
        return n*fact(n-1)
    end
end

-- This will work.
local fact
fact = function(n)
    if n == 0 or n == 1 then
        return 1
    else
        return n*fact(n-1)
    end
end
```

You can safely use the syntactic sugar version even for recursive functions.

```lua
-- This will work.
local function fact(n)
    if n == 0 or n == 1 then
        return 1
    else
        return n*fact(n-1)
    end
end
```

However, if you have indirect recursive functions, you must declare the
variables in advance again.

```lua
local f, g
function g()
    -- other code
    f()
    -- other code
end

function f()
    -- other code
    g()
    -- other code
end
```

(The last example seems awful to me.  Let's avoid it either way!)

### Proper Tail Calls

Ierusalimschy explains that "Lua is *properly tail recursive*" (52, emphasis
in the original).  What does this mean?  It means that under certain
conditions, the Lua interpreter can avoid extra stack variables.  As a result,
Lua can avoid overflowing the stack.  Here's an example.

```lua
function foo(n)
    if n > 0 then return foo(n-1) end
end
```

You might worry that this program will cause stack overflow if given very
larger numbers.  However, in Lua the program is safe with any number because
of tail-call elimination.  (Tail-call elimination occurs when an
implementation can detect that a function never needs to return to its caller.
Thus, they can avoid adding a return to the caller to the stack.)

This leads to the key question: in Lua, which returns are safe for tail-call
elimination and which are not?  Briefly, the function call at the end of
a function must "have nothing else to do after the call" (52).  Here are some
examples.

```lua
return g(x)

return x[i].foo(x[j] + a*b, i + j)
```

Notice that the second example looks very complex, but that doesn't matter.
Lua evaluates `x[i].foo`, `x[j] + a*b`, and `i + j` *before* the function
call, so there is no need to return to the enclosing function after the call
to `x[i].foo`.

On the other hand, the following returns fail to produce tail-call
elimination.  See the comments for why in each case.

```lua
g(x) -- The enclosing function has to discard results from the call to g.
return g(x) + 1 -- The enclosing function has to perform the addition.
return x or g(x) -- The enclosing function must adjust to one result.
return (g(x)) -- The enclosing function must adjust to one result.
```

According to Ierusalimschy, a tail call "is a goto dressed as a call" (52).
Therefore, he argues that you should consider using tail calls in Lua to
program state machines.  A program can represent each state with a function,
and to change state, you return another function.  He gives an example of
a maze game.  The current room is the state, and as players move from room to
room, the program maintains state via function calls rather than a table.  (As
Ierusalimschy points out, you could also use tables to write such a program.)

Without tail-call elimination, you could not use functions to write such
a program because "each user move would create a new stack level.  After some
number of moves, there would be a stack overflow" (53).  I doubt I will need
to worry about this much, but it's still a cool idea.  Here's the program.

```lua
function room1()
    local move = io.read()
    if move == 'south' then
        return room3()
    elseif move == 'east' then
        return room2()
    else
        print('invalid move')
        return room1() -- Stay in the room you're in on an invalid move.
    end
end

function room2()
    local move = io.read()
    if move == 'south' then
        return room4()
    elseif move == 'west' then
        return room1()
    else
        print('invalid move')
        return room2() -- Stay in the room you're in on an invalid move.
    end
end

function room3()
    local move = io.read()
    if move == 'north' then
        return room1()
    elseif move == 'east' then
        return room4()
    else
        print('invalid move')
        return room3() -- Stay in the room you're in on an invalid move.
    end
end

function room4()
    print('congratulations!')
end
```

You begin the game with a call to `room1()`.

## Chapter 7: Iterators and the Generic `for`

### Iterators and Closures

Ierusalimschy describes an iterator as "any construction that allows you to
iterate over the elements of a collection" (55).  He adds that iterators in
Lua are generally functions.  The function returns the next element from
a collection each time that the program calls the function.

Iterators must maintain state between iterations.  Otherwise, they wouldn't
know where they were or when to end.  Closures can provide the necessary state
since they can create state in the enclosing environment.  Since the closure
will need non-local variables, "a closure construction typically involves two
functions: the closure itself and a *factory*, the function that creates the
closure" (55).  He gives the following example.

```lua
function values(t)
    local i = 0
    return function()
        i = i+1
        return t[i]
    end
end
```

This function is like `ipairs`, but it returns values only instead of both
index positions and values.  You can use this iterator manually in a `while`
loop or automatically in a generic `for` loop.

```lua
t = {10, 20, 30}
iter = values(t)
while true do
    local element = iter()
    if element == nil then
        break
    end
    print(element)
end

t = {100, 200, 300}
for element in values(t) do
    print(element)
end
```

The generic `for` loop is much easier to write and read since it does all the
low-level work for you.  As Ierusalimschy explains, "it keeps the iterator
function internally, so we do not need the `iter` variable; it calls the
iterator on each new iteration; and it stops the loop when the iterator
returns `nil` (56).

Ierusalimschy also shows a more complex example, a function `allwords` that
iterates over all the words in a given input file.

```lua
function allwords()
    local line = io.read()
    local pos = 1
    return function()
        while line do
            local s, e = string.fine(line, '%w+', pos)
            if s then
                pos = e+1
                return string.sub(line, s, e)
            else
                line = io.read()
                pos = 1
            end
        end
        return nil
    end
end
```

As Ierusalimschy points out, this code may be somewhat complex, but the
calling code is still as simple as can be.  As Ierusalimschy suggests, this is
not a bad trade-off since you use iterators far more often than you write
them.

```lua
for word in allwords() do
    print(word)
end
```

### The Semantics of the Generic `for`

Internally, however, the generic `for` is far more complex than it may appear
initially.  Here's the syntax of the generic `for`.

```lua
for <variable list> in <expression list> do
    -- body
end
```

The variable list is one or more variable names, separated by commas, and the
expression list is one or more expressions, also separated by commas.
Usually, the expression list is only one element, namely a call to an iterator
factory such as `pairs` or `allwords`.  The list of variables can have
a single element (as in `allwords`) or it can have more (as in `pairs` and
`ipairs`).  The first variable in the list of variables is called the *control
variable*.  It cannot be `nil` at any time during the loop because the loop
ends when the control variable becomes `nil`.  (Thus, every iterator should
move towards a state where the control variable becomes `nil`.)

When a generic `for` loop begins, the expressions after `in` are evaluated.
They should produce the three values kept by the `for` loop, namely the
iterator function, an invariant state, and an initial value for the control
variable.  If there are multiple expressions, only the last can yield more
than one value, and the total number of values is adjusted to three.
(Additional values are discarded and `nil` values are added as necessary.)
Simple iterators return only the iterator function; they leave the invariant
state and control variable `nil` initially.

After initialization, the iterator function is called with two arguments,
namely the invariant state and the control variable.  The variables in the
variable list are then assigned the return values of the iterator function.
The loop continues until the control variable becomes `nil`.  If the control
variable is `nil` on the first round, then the code in the body of the `for`
loop is never run.  Thus, initialization happens once (no matter what),
the function is called one or more times (no matter what), and the code in the
body of the `for` loop is called zero or more times.

Ierusalimschy describes the for loop in the following terms.

```lua
-- for var1, ..., varN in <expression list> do <block> end

do
    local _f, _s, _var = <expression list>
    while true do
        local var1, ..., varN = _f(_s, _var)
        _var = var1
        if _var == nil then
            break
        end
        <block>
    end
end
```

### Stateless Iterators

A stateless iterator does not require a closure because it does not maintain
any state itself.  Instead, stateless iterators use the `for` loop itself to
maintain state.  A generic `for` loop passes two arguments to each call to the
iterator function, the invariant state and a control variable.  A stateless
iterator gets the next element for the iteration from these two values.

As an example, Ierusalimschy shows a version of `ipairs` in Lua.  (The actual
implementation is in C, but that isn't the point here.)

```lua
local function iter(t, n)
    i = i + 1
    local v = a[i]
    if v then
        return i, v
    end
end

function ipairs(t)
    return iter, t, 0
end
```

When a generic `for` loop initializes `ipairs`, it receives three values: the
internal iteration function, the table, and zero.  At the first call to the
iteration function—assuming that the table is not empty—the body of the `for`
loop receives two values, one and the first item in the table.  This continues
until the table is out of elements.  At that point, the control variable is
`nil`, and the loop ends.

The built-in `pairs` function is similar, but it uses a built-in iterator
function `next`.  The call `next(t, k)` works as follows.  If `k` is `nil`,
then `next` returns a key-value pair.  If `k` is a key in the table, then
`next` returns another key-value pair.  This continues until there are no more
key-value pairs.  The pairs are returned in random order.  Thus, you can use
`next` directly instead of `pairs`.

```lua
for k, v in next, t do
    -- body
end
```

Ierusalimschy also shows how to write and use an iterator to go through the
items in a linked list.

```lua
local function getnext(list, node)
    return not node and list or node.next
end

function traverseList(list)
    return getnext, list, nil
end

list = nil
for line in io.lines() do
    list = {val = line, next = list}
end

for node in traverseList(list) do
    print(node.val
end
```

### Iterators with Complex State

If you have an iterator that "needs to keep more state than fits into a single
invariant state and a control variable," you have two choices (60).  You can
use a closure, which Ierusalimschy says is the "simplest solution," or you can
put all the state into a table and use that table as the invariant.  This
works because the table itself is always the same reference (thus, invariant),
but you can change the data in the table as needed to move the state towards
an appropriate ending point.  Ierusalimschy also notes that iterators that
pack all their state into a table "typically ignore the second argument
provided by the generic `for` (the iterator variable)" (60).  They can ignore
that variable because that information is carried in the table invariant.

As an example, Ierusalimschy rewrites the `allwords` function.

```lua
local iterator

function allwords()
    local state = {line = io.read(), pos = 1}
    return iterator, state
end

function iterator(state)
    while state.line do
        local s, e = string.find(state.line, '%w+', state.pos)
        if s then
            state.pos = e+1
            return string.sub(state.line, s, e)
        else
            state.line = io.read()
            state.pos = 1
        end
    end
end
```

Here is Ierusalimschy's final advice about how to write iterators.

1. If possible, write stateless iterators.  Keep all the state in the `for`
   variables (the invariant state and the control variable).
2. If stateless iterators are impossible, use a closure.  Closures are more
   efficient than tables.
3. Coroutines are more powerful than the two previous options, but they are
   also (a little?) more expensive.

### True Iterators

The iterators that we've seen so far don't actually iterate; the `for` loop
handles the mechanics of iteration.  (Ierusalimschy suggests that *generator*
is a better name than *iterator*, but *iterator* is already the standard name
in many languages.)

However, Lua can support Ruby-style iteration too.  For these iterators, "we
do not write a loop; instead, we simply call the iterator with an argument
that describes what the iterator must do at each iteration. More specifically,
the iterator receives as argument a function that it calls inside its loop"
(61).

As his example, Ierusalimschy rewrites `allwords` yet again.

```lua
function allwords(callbackFunction)
    for line in io.lines() do
        for word in string.gmatch(line, '%w+') do
            callbackFunction(word)
        end
    end
end
```

You can pass a named or an anonymous function as the callback function.

```lua
allwords(print)

local count = 0
allwords(function(w)
    if w = 'hello' then
        count = count + 1
    end
end)
print(count)
```

Of course, you can also do this using the generic `for` loop.

```lua
local count = 0
for w in allwords() do
    if w == 'hello' then
        count = count + 1
    end
end
print(count)
```

Older versions of Lua did not have the `for` statement, and therefore true
iterators were very popular.  Ierusalimschy gives the following comparison of
the two styles in Lua 5.1.

1. They have approximately the same overhead: one function call per iteration.
2. It is easier to write the iterator with true iterators, but coroutines can
   make generic `for` iteration easy again.
3. The generic `for` loop is more flexible.  (He asks you to imagine iterating
   over two files at the same time and comparing them word by word.  He
   clearly thinks that this is easier with the generic `for` loop, but I am
   not sure why.)
4. The generic `for` loop allows you to use `break` and `return` from inside
   the iterator body.  With a true iterator, `return` returns from the
   anonymous function rather than the function doing the iteration.

Ierusalimschy prefers generators (i.e., the generic `for` loop) overall (62).

## Chapter 8: Compilation, Execution, and Errors

Ierusalimschy explains that even though we call Lua an *interpreted language*,
Lua still precompiles source code before running it.  (As Ierusalimschy says,
many languages do the same.)

### Compilation

Lua's `dofile` is actually a helper function for the more primitive `loadfile`
function.  `loadfile` compiles a chunk of Lua code and returns the compiled
chunk as a function.  In addition, instead of raising errors (like `dofile`
does), `loadfile` returns error codes.  `dofile` is essentially the following,
if we imagine it written on the basis of the primitive `loadfile`.

```lua
function dofile(filename)
    local f = assert(loadfile(filename))
    return f()
end
```

By using `assert`, `dofile` raises an error if `loadfile` returns an error
code.

`loadfile` is more flexible than `dofile`.  `loadfile` returns `nil` and an
error if it fails.  In that case, you can handle the error however you like.
In addition, `dofile` runs the returned function once, but with `loadfile`,
you can call the result several times if wanted.  (This is better than
repeatedly running `dofile` since the code is compiled only once via
`loadfile`.)

Lua also provides a function `loadstring`.  It works like `loadfile`, but as
the name implies, `loadstring` reads code from a string rather than a file.
Ierusalimschy warns that `loadstring` is expensive and that it can make your
code unclear.  He recommends that users avoid it if possible.  However, he
also shows that if you call the result of `loadstring` immediately, you get
a kind of cheap and dirty `dostring`: `loadstring(s)()`.  (I can imagine using
this in a snippets library, maybe.)  But Ierusalimschy also warns that if you
use `loadstring` that way, you may get a return value of nil plus an error
message. In addition, `loadstring` always "compiles its strings in the global
environment" (65).  It does not support lexical scoping at all.

Ierusalimschy says that the most common use of `loadstring` is to run external
code, and he gives the following examples.

```lua
print "enter your expression:"
local l = io.read()
local func = assert(loadstring("return " .. l))
print("the value of your expression is " .. func())

print "enter function to be plotted (with variable 'x'):"
local l = io.read()
local f = assert(loadstring("return " .. l))
for i=1, 20 do
    x = i -- global 'x' (to be visible from the chunk)
    print(string.rep("*", f()))
end
```

Notice a few things about this code.  First, Ierusalimschy uses `assert` to
get better error messages if the call to `loadstring` fails.  (Otherwise,
`loadstring` returns nil, and the error would be simply "attempt to call a nil
value".  See page 64.)  Second, Ierusalimschy has to add `return ` to the
string before calling `loadstring`.  Why?  Because "`loadstring` expects
a chunk, that is, statements. If you want to evaluate an expression, you must
prefix it with `return`, so that you get a statement that returns the value of
the given expression" (65).

Even more primitive than `loadfile` or `loadstring` is `load`.  `load`
receives neither a file nor a string.  Instead, it receives a reader function
that it calls repeatedly until the function returns nil.  Why would anyone
want this?  Ierusalimschy suggests two uses.  First, maybe the chunk is not in
a file or string but created dynamically or read from another source in
pieces.  (E.g., streaming or buffered reads?)  Second, maybe the chunk is too
big to fit easily into memory.

Ierusalimschy shows a technique for using (sort-of) lexical scoping with
`load` functions.  Lua treats any independent chunk "as the body of an
anonymous function with a variable number of arguments.  That is, if the chunk
is `a = 1`, Lua treats that like this: `function (...) a = 1 end`.  Since
chunks themselves ("like any other function", (65)), can declare local
variables, you can do the following.

```lua
print "enter function to be plotted (with variable 'x'):"
local l = io.read()
local f = assert(loadstring("local x = ...; return " .. l))
for i=1, 20 do
    print(string.rep("*", f(i)))
end
```

When the chunk is loaded, `i` becomes the value of the vararg expression in
the anonymous function.

You cannot use chunks to (directly) define functions because functions are
assignments.  As Ierusalimschy says, "They are made at runtime not compile
time" (66).  He gives the following example.

```lua
-- Imagine this chunk.
function foo(x)
    print(x)
end

-- Run this command.
f = loadfile("foo.lua")

-- The code in foo is compiled, but not yet defined. To define it, you must
-- run the chunk once. Hence, the following.
print(foo) -- foo is still nil; this prints nil.
           -- If you tried, foo("ok") at this point, you would get an error.
f()        -- Now foo is defined as a function.
foo("ok")  -- This prints "ok".
```

Finally, Ierusalimschy ends with this note: "In a production-quality program
that needs to run external code, you should handle any errors when loading
a chunk.  Moreover, if the code cannot be trusted, you may want to run the new
chunk in a protected environment, to avoid unpleasant side effects when
running the code" (66).

### Errors

Lua raises an error whenever it runs into an unexpected condition.  For
example, if a program tries to call something that isn't a function or to
index something that isn't a table or to add two values that aren't numbers.
Programs can also explicitly raise an error using `error`.  The function
`error` takes a string that it uses as an error message.

```lua
print n "enter a number:"
n = io.read("*number")
if not n then error("invalid input") end
```

You can use `assert` as a way to avoid having to write `if not...then error`
too often in your code.  `assert` takes two parameters.  The first is
a function call that may return an error, and the second is a string.  If the
function call is false (if it returns `false` or `nil`), then an error is
raised with the string as its message.

```lua
print n "enter a number:"
n = assert(io.read("*number"), "invalid input")
```

If the first argument to `assert` returns anything other than `false` or
`nil`, that value is returned.  Otherwise, an error is raised with the string
value as the message.

However, since `assert` (like all other functions) evaluates all its arguments
*before* the call, you have to be careful with the message.  Ierusalimschy
gives this example.

```lua
n = io.read("*number")
n = assert(tonumber(n), "invalid input: " .. n .. " is not a number")
```

When you call `assert`, the concatenation happens, even if `n` *is* a number.
In those cases, Ierusalimschy suggests that it "may be wiser to use an
explicit" test (68).

When a function runs into something unexpected, Ierusalimschy says that it can
do one of two things.  The function can return an error code (often `nil` in
Lua), or the function can raise an error.  He offers the following "general
guideline" for choosing what to do (68).  If the exception is "easily
avoided," the function should raise an error.  If not, the function should
return an error code.

As an example of the second situation, Ierusalimschy discusses `io.open`.
There is no easy way for a program to check if a file can be opened before
calling `io.open`.  (To check, you essentially have to try to open the file!)
Therefore, if `io.open` fails, it returns nil and an error message string.
The calling code can then handle the problem in an appropriate way.  For
example, he shows the following.

```lua
local file, msg
repeat
    print "enter a file name:"
    local name = io.read()
    if not name then return end -- no input
    file, msg = io.open(name, "r")
    if not file then print(msg) end
until file
```

Sometimes, you don't want to handle the problem in any special way, but you
still want to catch the error, you can use `assert`.

```lua
file = assert(io.open('non-existent-file', 'r'))
-- stdin: 1 non-existent-file: No such file or directory
```

The error message from `io.open` ends up as the second argument (the error
message) to `assert`, which is clever, but a bit too magical for me.

As an example of the first situation, Ierusalimschy gives the following
example.

```lua
if not tonumber(x) then
    -- error-handling code
end
local res = math.sin(x)
```

We can easily check whether the argument to `math.sin` is a number before we
run the call.  Thus, `math.sin` raises an error if you pass something other
than a number.

### Error Handling and Exceptions

In many cases, you don't do any error handling in Lua.  When Lua runs embedded
in some other application, error handling is often up to the host application.

If you want to handle errors in Lua itself, you need to use the `pcall`
function.  `pcall` wraps another function, and it returns two (or more)
values.  The first value is `true` if the wrapped function ran without an
error and `false` otherwise.  If the function runs without error, after `true`
`pcall` returns the values returned by the function.  Otherwise, `pcall`
returns `false` and then an error message.  As Ierusalimschy says, the error
message does not have to be a string.  You can return a table if you prefer.
In fact, you can return any Lua value, but I can't see a reason to return
anything other than a string or a table.  The error message return value is
usually a string.

This leads to the following idiom for Lua code.

```lua
if pcall(function()
    -- protected code
end) then
    -- regular code
else
    -- error-handling code
end
```

What you pass to `pcall` can be (as here) an anonymous function or a regular
function call.  Whatever arguments you need to pass to the function, you add
as arguments to `pcall` after the function itself.

Ierusalimschy says "These mechanisms provide all we need to do exception
handling in Lua.  We *throw* an exception with error and *catch* it with
pcall.  The error message identifies the kind of error" (70).

### Error Messages and Tracebacks

When you call `error`, you can add a second argument to specify the level
where the error should be reported.  By default, the value is 1, meaning that
the error is reported from wherever `error` is called.  But let's say we have
a function that checks whether another function has proper arguments.  In that
case, we might use `error(message, 2)` to say that the real problem is an
extra level up in the calling stack.

If you want more detailed tracebacks, you need to use `xpcall` rather than
`pcall`.  `xpcall` receives two arguments, a function call call with `pcall`
and an error-handling function to use if `pcall` returns an error.  The
advantage of `xpcall` is that it preserves more of the stack than `pcall`
does.  The disadvantage of `xpcall` (in Lua 5.1, at least) is that you cannot
add additional arguments for the function you want to test.  If your function
needs arguments, you have to wrap the function in another level of
indirection.  Here's an example.

```lua
-- See Stackoverflow: https://stackoverflow.com/a/30125834.
local function f(a,b)
  return a + b
end

local function err(x)
  print ("err called", x)
  return "oh no!"
end

local function pcallfun()
    return f(1,2)
end

status, err, ret = xpcall(pcallfun, err)
```

Ierusalimschy says that two common error handlers are `debug.debug` and
`debug.traceback`.  The first drops you into a Lua prompt to debug, and the
second creates an extended error message.  (He will discuss the debugging
library more in a later chapter.)
