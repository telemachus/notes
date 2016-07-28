# Chapter 1: Getting started

## Chapter 1

### Chunks

Lua executes code as *chunks*. A chunk can be an entire file or a single line.
A sequence of commands or statements makes up a chunk. Chunks as large as
several megabytes are not uncommon, and there's no inherent limit on the size
of a chunk.

In between statements, Lua does not require a separator, but you can use
a semicolon for clarity. (RI prefers to use semicolons where there are two or
more statements on the same line, but not otherwise.)

When in interactive mode, the interpreter usually reads each line as a complete
chunk, unless it can tell that the line is not yet complete.

You can manually run a chunk from a file with the `dofile` function, which
takes a filename as a parameter. E.g.

```
dofile("important_library.lua")
```

`dofile` immediately executes a file.

### Some lexical conventions

Identifiers in Lua can be any string of letters, digits, and underscores, but
cannot begin with a digit. Avoid identifiers that begin with an underscore and
contain all capital letters afterwards. Lua reserves many of these for special
use. (E.g. `_VERSION`.)

Lua also reserves a small number of words for its own use:

    and     break       do          else    elseif
    end     false       goto        for     function
    if      in          local       nil     not
    or      repeat      return      then    true
    until   while

Lua is case-sensitive, so `and` is not the same as `AND`.

Lua provides two kinds of comments:

+ Single-line comments begin with `--` and run to the end of the line.
+ Block comments begin with `--[[` and run until `--]]`.

### Global variables

You do not need to declare variables before use in Lua. And it is not an error
to access a non-initialized variable. Any non-intialized variable simply has
the value `nil`. You can also assign any variable the value `nil` in order to
tell Lua that it is no longer in use. (This can sometimes be useful for manual
memory management, though I doubt I would normally need to do it.)

### The stand-along interpreter

The stand-alone interpreter (the file `lua.c` and the executable `lua`)
provides a way to run Lua code directly. It has a few simple options.

+ `-e` allows the direct evaluation of code on the invoking command-line
+ `-l` loads a library
+ `-i` enters interactive mode after any other arguments are handled

The interpreter looks for environmental variables named `LUA_INIT_5_2` or
simply `LUA_INIT`. Lines starting with `@` are assumed to point towards files,
which Lua will interpreter to load. All other lines are assumed to be Lua code,
which the interpreter will run as it starts up.

Lua provides a predefined variable that holds program arguments `arg`. The
script name itself is `arg[0]`, and all other arguments are `arg[1]...arg[2]`
etc. Options and their arguments are held in negative indices. E.g.

```
lua -e "sin=math.sin" script a b
-- arg[-3] = "lua"
-- arg[-2] = "-e"
-- arg[-1] = "sin=math.sin"
-- arg[0] = "script"
-- arg[1] = "a"
-- arg[2] = "b"
```

In general, scripts use only the positive indices. In addition to `arg`,
a script can retrieve its arguments as a vararg `...` instead of as a fixed
table.
