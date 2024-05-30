# “Mastering LPeg” Roberto Ierusalimschy

## Introduction

> LPeg is a pattern-matching library for Lua, based on Parsing Expression
> Grammars (1).

We use pattern matching to find extract specific items from texts.  Both the
texts and the items we want to extract can be very varied.  Many programs use
regexes for pattern matching, but LPeg is not based on regexes.

Here are some of the ways that LPeg patterns are different from regexes.

+ LPeg patterns are first-class objects.  This means that patterns can be
  treated like normal Lua values.
+ LPeg provides many functions to create and compose patterns; using
  metamethods, some of these functions are infix or prefix operators.
+ LPeg patterns are usually more verbose than patterns that use regexes.
+ LPeg allows you to build complex patterns piece by piece.  Thus, you can
  test each piece independently, give them clear names, and document them
  separately.  If you use a regex, you get one very noisy and dense item for
  the pattern.  You cannot name, reuse, or easily document a regex pattern.

## Basics

```lua
local lpeg = require "lpeg"

local p = lpeg.P("hello")
print(lpeg.match(p, "hello world"))     --> 6
print(lpeg.match(p, "hi world"))        --> nil
```

`lpeg.P` creates patterns that match a given string literally.  `lpeg.match`
is a function that takes a pattern and a string and tries to match the pattern
against the string.  If there is a match, `lpeg.match` returns the first
string index *after the match ends*; if there is no match, `lpeg.match`
returns `nil`.  The function `match` is also available as a method of
patterns: `p:match("hello world")`.

Note that `match` does not search for a pattern throughout a string.  Instead,
the function is equivalent to an anchored match in a regex library.  The match
begins (by default? always?) at the start of the string.  Thus, if a match is
successful, the return value of `match` is always one more than the number of
matched charcters.  There are other ways to search for a pattern in a string.

In addition, `lpeg.P` treats all characters literally.  LPeg provides other
functions for character classes, repetition, and other such abstract matching.

Here are some other basic pattern making functions in LPeg.

+ `lpeg.S` creates a pattern that matches any character in the set of
  characters given; again, the match is anchored.  For example,
  `lpeg.S("aeiou")` searches for any one of the five vowels.
+ `lpeg.R` creates a pattern that looks for one ore more ranges; each range is
  a two-character string.  For example, `lpeg.R("af", "AF", "09")` matches any
  hexadecimal digit.
+ `lpeg.locale` creates a set of patterns for the current locale.  It
  provides, for example, `space`, `alpha`, `alnum`, `digit`, `lower`, `upper`,
  `punct`, and others.

You can also call `lpeg.P` with numbers or booleans.  If you call `lpeg.P`
with a positive integer, then you get a pattern that matches that number of
characters no matter what the characters are.  In other words, such a pattern
only fails if the text does not contain enough characters.  If you call
`lpeg.P` with `true` it always succeeds, and if you call it with `false`, it
always fails.  Any use of `lpeg.P` with a boolean consumes no input.

The following patterns are equivalent.

+ `lpeg.P(true)`, `lpeg.P("")`, `lpeg.P(0)` always succeed and consume no
  input
+ `lpeg.P(false)`, `lpeg.S("")`, `lpeg.R("za")` always fail and consume no
  input

### A note about Unicode

LPeg operates will often do the right thing with Unicode text, but you still
have to be careful.  A single character for LPeg is a byte, not a Unicode code
point.  Literals, concatenations, repetitions, and predicates will do the
right thing since Lua itself uses UTF-8 for text encoding.  However, sets and
ranges in LPeg only accept ASCII characters.  Also, `lpeg.P(1)` matches
a single byte rather than a single Unicode character.

## Repetitions and Choices



LPeg makes it easy to make larger patterns by combining patterns.  The library
overloads existing operators for this purpose.

+ `p1 * p2` matches the first pattern followed by the second pattern
+ `p1 + p2` matches the first pattern or the second pattern (ordered choice)
+ `p1 - p2` matches the first pattern if the second pattern does not match
