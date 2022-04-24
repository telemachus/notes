# Beginning Lua Programming (Kurt Yung and Aaron Brown)

## Chapter 1: Getting Situated

I skipped this chapter.

## Chapter 2: First Steps

I skimmed this chapter.

### The Modulo Operator

The authors explain that "the modulo operator, `%`, is good for modeling
cyclical things like hours in a day, days in a week, or musical pitches in an
octave." (47)  They give the following example of modeling a clock face.

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

Because relational operators have relative low precedence, the following work
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
