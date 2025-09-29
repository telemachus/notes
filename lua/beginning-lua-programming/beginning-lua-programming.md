# Beginning Lua Programming (Kurt Yung and Aaron Brown)

## Chapter 1: Getting Situated

I skipped this chapter.

## Chapter 2: First Steps

I skimmed this chapter.

### Valid Identifiers

A valid identifier (or valid name) in Lua can be any sequence of letters,
digits, and underscores with a few restrictions.

+ A valid identifier cannot start with a digit.
+ The following words are reserved in Lua 5.1.  You cannot use them as valid
  identifiers: `and`, `break`, `do`, `else`, `elseif`, `end`, `false`, `for`,
  `function`, `if`, `in`, `local`, `nil`, `not`, `or`, `repeat`, `return`,
  `then`, `true`, `until`, `while`.
+ By convention, you should not use identifiers that start with an underscore
  and contain only uppercase letters.  E.g., `_VERSION`.  Identifiers like
  this are used as internal global variables by Lua itself, and so we should
  avoid them.

### The Modulo Operator

The authors explain that "the modulo operator, `%`, is good for modeling
cyclical things like hours in a day, days in a week, or musical pitches in an
octave" (47).  They give the following example of modeling a clock face.

```lua
hour = 3
hour = hour + 2
hour = hour % 12
print(hour)         -- 5
hour = hour + 144
hour = hour % 12
print(hour)         -- 5
hour = hour - 149
hour = hour %12
print(hour)         -- 0
```

### Precedence and Associativity

Whenever two operators have the same precedence (e.g., `+` and `-`), Lua
determines the order of operations based on associativity.  Since both `+` and
`-` are left-associative, Lua will perform operations from left to right.

The only right-associative operators in Lua are `..` (string concatenation)
and `^` (exponentiation).  All other operators are left-associative.  However,
all unary operators are applied from right to left.  As Yung and Brown say,
"that's the only way that makes sense" (49).

Here is the order of precedence for operators in Lua.  The top of the list is
highest precedence.

+ `^`
+ `not`, `-` (unary minus), `#`
+ `*`, `/`, `%`
+ `+` `-` (subtraction)
+ `..`
+ `<`, `>`, `~=`, `<=`, `>=`, `==`
+ `and`
+ `or`

Because relational operators have relatively low precedence, the following work
without parentheses.

```lua
print(creditA - debitA == creditB - debitB)
print(credita - debitA > creditB - debitB)
```

Similarly, `and` and `or` have low precedence, so the following work without
parentheses.

```lua
print(creditA >= debitA and creditB >= debitB)
print(creditA >= debitA or creditB >= debitB)
```

Finally, because `or` has lower precedence than `and`, you can use these two
in the idiom `x and y or z`.  This works out to be `if x, then y; otherwise
z`.  However, this idiom only works if the middle value is true.  If the
middle value is false, then the third value is returned even when the first is
true.

You can use values other than `true` and `false` for all boolean operations.
As a result, you can use values other than `true` and `false` in the `x and
y or z` idiom.  `false` and `nil` are falsy in Lua.  All other values are
truthy.

### Expressions and Statements

The authors define a *statement* in Lua as "the smallest complete unit of Lua
code" (54).  Lua program is technically zero or more statements, though
usually one or more.  On the other hand, they define an *expression* as
something in Lua that has a value.  An expression can be a *literal
expression* (e.g., `5`, `'foo'`, or `true`) or a variable (e.g., `count`,
`name`, or `bool`).  Expressions can form larger expressions by using unary or
binary operators.  For example, `5` is an expression and so is `-5` or `5
+ 5`.

Just as individual expressions can combine into larger expressions, Lua has
compound statements.  For example, Lua provides the following compound
statements: `if`, `while`, `for, `repeat`, and `do`.

```lua
if expression then
    -- zero or more statements
elseif -- optional elseif clauses
    -- zero or more statements
else -- an optional else clasuse
    -- zero or more statements
end

while expression do
    -- zero or more statements
end

for variable = start, end, step do
    -- zero or more statements
end

repeat
    -- zero or more expressions
until expression
```

You use a `break` statement to exit from a `while`, `repeat`, or `for` loop
early.  You can only put a `break` statement at the end of a block.  However,
you can use `do...end` to create a block if you want to put `break` somewhere
that it cannot otherwise fit.

## Chapter 3: Extending Lua with Functions

Since functions can return values, functions are expressions that you can
(also) use as statements.  When you use a function as a statement, Lua
discards its return value.  This often makes sense to do, however, since you
often call a function for its side effects not to capture some value it
returns.

Lua functions can return multiple values.  The return values are adjusted to
fit how and where the function is called. For example, if a function returns
three values, and it is the first expression in a list of expressions, it only
returns one value, and the other two are thrown away.  In general, a Lua
function only returns multiple values if it is the last or only expression in
a list of values.  Otherwise, only the first return value is returned.

You can force a function call at the end of a value list to return only one
value by placing parentheses around it.

Here's a weird wrinkle: if a function returns no values, but it is placed
anywhere other than as the last (or only) expression in a list of values, then
it returns one value, namely `nil`.

Finally, when a function is used as a statement, it returns no values
regardless of how many values it returns in other cases.

The authors stress that the order of side effects is implementation defined
and not guaranteed.  For example, the Lua 5.1 interpreter assigns from right
to left, but this is not guaranteed.  Here's an example.

```lua
a, a = 1, 2
-- a = 1 since `a = 1` happens second.
```

There is no guarantee about the order of operations if you call a function
multiple times in an expression.  In Lua 5.1, these calls happen left to
right, but that may change, and you should not rely on it.

## Chapter 4: Working with Tables

Tables are Lua's only built-in data structure.  (You can use C to create
additional data structures for a particular application.)  Tables can be used
as arrays or dictionaries or both at the same time.  As a result, you can
create and fill tables in several ways.

### Tables as dictionaries

If you use a table as a dictionary, you can create it in the following way:

```lua
name_to_instrument = {
        ["John"] = "rhythm guitar",
        ["Paul"] = "bass guitar",
        ["George"] = "lead guitar",
        ["Ringo"] = "drums",
    }
```

You can access the values by indexing the dictionary. E.g.,
`name_to_instrument["John"]`.

This is the long way to create a dictionary-style table, and it will always
work. You place the keys inside square brackets and quote them, and you quote
the keys.  If the keys are valid identifiers, however, you can drop the square
brackets and the quotes.

```lua
name_to_instrument = {
        John = "rhythm guitar",
        Paul = "bass guitar",
        George = "lead guitar",
        Ringo = "drums",
}
```

In addition, if a key is a valid identifier, then you can get at the value
using a period instead of square brackets and quotes. E.g.,
`name_to_instrument.George`.

### Tables as arrays

You can also use a table as an array.  In this case, the keys are successive
integers.  By convention, Lua tables number their contents from 1 not 0.  You
don't have to follow this convention, but you should.  If you don't, you will
have inconsistencies or problems when you use the standard library or other
Lua modules.

You can create array-like tables in more than one way in Lua.

```lua
-- the long way.
foods = {
        [1] = "pizza",
        [2] = "mapo tofu",
        [3] = "momos",
}
-- the short way.
foods = { "pizza", "mapo tofu", "momos" }
```

If you use the shorter style, Lua will automatically index the items in the
list starting from 1.

### Mixed tables

Although I hate this, you can create tables that are simultaneously
dictionaries and arrays.  (More precisely, all tables are always dictionaries.
But some tables have portions that are indexed by integer keys as well as
portions that are indexed by non-integer keys.)

```lua
t = {
        a = "x",
        "one",
        b = "y",
        "two",
        c = "z",
        "three",
}
```

In this table, `t.a`, `t.b`, and `t.b` are non-integer keys, while `t[1]`,
`t[2]`, and `t[3]` are integer keys.  This is very hard to read or keep track
of, and I will avoid it whenever possible.

### Table creation and function calls

If you use a function call to populate a table, the return values may be
adjusted.  If the function call provides the value for an explicit key
(integer or otherwise), then it is adjusted to one value.  If the function
call provides the value of an implicit integer key, it will be adjusted to one
value unless it is the last thing in the table constructor.  In the case of an
implicit integer key that is last in the table constructor, all values are
returned and assigned in order.

This is confusing overall, but it provides one advantage: you can fill an
entire array-like table with a single function call.  Since it is the only
(and therefore last) item, all values will be returned and be assigned as
items in the array.

```lua
local f = function return 1, 2, 3, 4 end
local t = { f() }
-- t = { 1, 2, 3, 4 }
```

### Array length

You can use `#` to find the length of an array-like table.  However, this is
subject to some constraints and gotchas.  If the table has any gaps in its
sequence, the `#` operator is unpredictable.  Be sure not to build tables with
gaps unless you must and you know what you are doing.

### Looping over tables

You should use `ipairs` or `pairs` to loop over tables.  `ipairs` goes through
an array-like table in order from 1 to `#table`.  `ipairs` assigns two items
each time through the loop, an index and a value.  If you don't need the index
or value, you can assign them to `_`, which is the conventional dummy value in
Lua.  (This dummy value is not enforced in any way by Lua.  Lua is not Go.)
`pairs` loops over tables in a dictionary-like fashion.  The order is
arbitrary, and the keys can be anything.  However, `pairs` guarantees that
each key is visited exactly once.  During a `pairs` loop, you can remove a key
(set it to `nil`) and you can change the value of a key (by indexing it, not
by simple assignment).  However, you cannot add a new key to a table while
looping with `pairs`.

You can use the variables in a loop as upvalues in closures.  For example,
consider the following.

```lua
numbers = { "one", "two", "three" }
prepend_number = {}
for num, num_name in ipairs(numbers) do
    prepend_number[num] = function(s)
        return num_name .. ": " .. s
    end
end
prepend_number[2]("is company")
prepend_number[3]("is a crowd")
```

### Tables of functions

You can use tables to store functions, and this is how Lua organizes its
standard library.  For example, many functions concerning tables are stored in
a `table` table.

+ `table.sort` sorts a table in place.  By default, the function uses `<` to
  decide the order.  You can override the default sort if you provide
  a sorting function as a second argument.  The sorting function should take
  two arguments and return a true value if the first argument should go before
  the second.  This function only works for array-like tables or the
  array-like portion of mixed tables.
+ `table.insert` adds an item into an array-like table or the array portion
  of a mixed table.  You can also do `t[#t + 1] = x` instead of
  `table.insert(t, x)`  This function takes an optional third argument that
  determines where in the table to insert the item.  After insertion anywhere
  but the end, all other items are moved and given a new index.
+ `table.remove` removes an item from an array-like table or the array portion
  of a mixed table.  This function returns the item that was removed.  By
  default, this function removes the last item from the table.  But you can
  give an optional second argument that specifies what index of the table to
  remove.  After removal, the indexes of the table are adjusted so that there
  are no gaps.
+ `table.concat` takes an array of strings or numbers and concatenates them
  into a string joined by the second argument. E.g., `table.concat(t, ", ")`.
  You can pass optional third and fourth arguments: these determine the start
  and stop element in the table to concatenate.  If the second argument is
  nil, it defaults to the empty string.  If the third argument is nil, it
  defaults to 1.  If the fourth argument is nil, it defaults to the length of
  the table.  (If the third argument is greater than the fourth, then the
  return value will be an empty string.)
+ `table.maxn` returns the highest positive number used as a key for a table.
  This may seem to help the problem of gappy tables, but it has potential
  problems.  First, it can be slow.  Second, it will look at fractional keys
  as well as integer keys.  (Who is giving fractional keys to their tables?)
