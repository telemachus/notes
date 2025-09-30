# Introduction to Zshell

## Note

I forked this from Clément Nerma's [ZSH cheat sheet][zsh-cheat-sheet].


[zsh-cheat-sheet]: https://gist.github.com/ClementNerma/1dd94cb0f1884b9c20d1ba0037bdcde2

## Parameter Expansions

| Description                                           | Syntax                       |
| ------------------------------------------------------| ---------------------------- |
| Returns the length of name                            | `${#name}`                   |
| Returns item at name[index]                           | `${name[index]}`             |
| Returns items from name[from] to name[to]             | `${name[from,to]}`           |
| Negative indices work too                             | `${name[from,-1]}`           |
| Replace the first occurrence of pattern with repl     | `${name/pattern/repl}`       |
| Replace all occurrences of pattern with repl          | `${name//pattern/repl}`      |
| Returns name minus shortest match of pattern at start | `${name#pattern}`            |
| Returns name minus longest match of pattern at start  | `${name##*pattern}`          |
| Returns name minus shortest match of pattern at end   | `${name%pattern*}`           |
| Returns name minus longest match of pattern at end    | `${name%%pattern*}`          |
| Get directory name (like `dirname`)                   | `${name:h}`                  |
| Get filename (like `basename`)                        | `${name:t}`                  |
| Get extension                                         | `${name:e}`                  |
| Remove extension                                      | `${name:r}`                  |
| Convert to absolute path                              | `${name:A}`                  |
| Convert to lowercase                                  | `${name:l}`                  |
| Convert to uppercase                                  | `${name:u}`                  |
| Test whether name starts with pattern                 | `if [[ $name = pattern* ]]`  |
| Test whether name contains pattern                    | `if [[ $name = *pattern* ]]` |
| Test whether name ends with pattern                   | `if [[ $name = *pattern ]]`  |

To help understand `${name#pattern}` and company, here are some examples.

```zsh
name=/path/to/something/something.txt

print ${name#*/}               # path/to/something/something.txt
print ${name##*/}              # something.txt
print ${name%/something*}      # /path/to/something
print ${name%/something*}      # /path/to
```

These examples can use simple strings as well as `*`, `?`, and `[...]` for
globbing. `*` means zero or more of any character; `?` means any character
exactly once; `[...]` means any of the items in the brackets exactly once.

## Arrays

Note that in Zshell, array indices start at 1 rather than 0.

| Description                                                  | Syntax                                       |
| -----------                                                  | ------                                       |
| Declare an array                                             | `typeset -a name`                            |
| Declare and initialize an array                              | `typeset -a name=(val1 val2 val3 val4 val5)` |
| Get an array's length                                        | `${#name}`                                   |
| Get all values of an array                                   | `"${name[@]}"`                               |
| Copy an array into a new variable                            | `typeset -a other_name=("${name[@]}")`       |
| Clear array and assign new values                            | `name=(v1 v2 v3)`                            |
| Add new value to end of array                                | `name+=(v4)`                                 |
| Add new value at a specific index                            | `name[5]=v5`                                 |
| Access an array's element (or empty string if no such index) | `${name[index]}` or `$name[index]`           |
| Remove first element from an array (shift)                   | `shift name`                                 |
| Remove last element from an array (pop)                      | `shift -p name`                              |
| Iterate over an array's values                               | `for val in "${name[@]}"; do...done`         |
| Get array slice from specified index to end of array         | `${name:index}`                              |
| Get array slice of *length* starting at specified index      | `${name:index:length}`                       |
| Check whether an array is empty                              | `if (( ${#name} == 0 ))`                     |
| Check whether an array is not empty                          | `if (( ${#name} > 0 ))`                      |
| Remove element at index from an array                        | `name[index]=()`                             |
| Check whether a value is contained in an array               | `if (( $name[(Ie)value] ));`                 |

The last test needs some explaining: `$name[(Ie)value]` tests for 'value' in the
`$name` array. By using `(I)`, the test returns 0 if 'value' is not found; by
using `(e)`, the test uses plain string matching rather than pattern matching.
You can omit the `(e)` if you know that 'value' is already a plain string.

When you add an item to a specific index of an array, you can (deliberately or
accidentally) create a sparse array. Consider the following.

```
typeset -a name=(v1 v2)
name[100]=v100
```

Items at `name[3]` through `name[99]` are now the empty string.

## Associative Arrays (i.e., Maps, Hashes, or Dictionaries)

Associative arrays are the equivalent of hash maps or dictionaries in many other
programming languages: unlike arrays, they can use string keys, and these don't
necessarily have an order.

| Description                                            | Syntax                                                |
| -----------                                            | ------                                                |
| Create an associative array                            | `typeset -A name=()`                                  |
| Create an associative array with initial values        | `typeset -A name=( [key1]=value1 [key2]=value2 )`     |
| Clear an associative array and add new values          | `name=([newKey]=new_value [otherKey]=other_value)`    |
| Add a new key to the array                             | `name[key]=value`                                     |
| Access a value by key                                  | `${name[key]}` or `$name[key]`                        |
| Remove a key-value pair from an associative array      | `unset 'name[key]'`                                   |
| Get the number of elements in an associative array     | `${#name}`                                            |
| Iterate over an associative array by values            | `for value in "${name[@]}"; do...done`                |
| Iterate over an associative array by keys              | `for key in "${(k)name[@]}"; do...done`               |
| Iterate over an associative array by sorted keys       | `for key in "${(ko)name[@]}"; do...done`              |
| Iterate over an associative array by key-value pairs   | `for key value in ${(kv)name[@]}; do...done`          |
| Check whether a key is present in an associative array | `if [[ -v name[key] ]];` or `if (( ${+name[key]} ));` |
| Get all values from an associative array               | `"${name[@]}"`                                        |
| Get all keys from an associative array                 | `"${(k)name[@]}"`                                     |
| Check whether an associative array is empty            | `if (( ${#name} == 0 )`                               |

Note that Zshell's arrays and associative arrays are like Lua tables in several
ways.

+ They are both associative arrays under the hood. List-like arrays are simply
  associative arrays that are indexed by integers.
+ List-like arrays start their indexing at 1 rather than 0.
+ All arrays in Zshell support sparse indexing. (Empty items are filled with
  `""`, an empty string.)

However, unlike tables in Lua, an array in Zshell must be declared as an
associative array or not. If it is not, then it will be indexed by integers
starting at 1. Also, no array in Zshell can contain both string and integer
indexes.

## Arithmetic

Zsh provides several contexts for arithmetic operations. Variables inside
arithmetic contexts don't need the `$` prefix (though it's allowed).

| Description                                     | Syntax                           |
| -----------                                     | ------                           |
| Arithmetic expansion (returns result)           | `$(( expression ))`              |
| Arithmetic evaluation (for conditionals)        | `(( expression ))`               |
| Arithmetic assignment                           | `(( name = expression ))`        |
| Assign result to variable                       | `name=$(( expression ))`         |
| Arithmetic increment or decrement               | `(( name++ ))` or `(( ++name ))` |

Note that there is a subtle difference between `(( var = expression ))` and
`name=$(( expression ))`. In the first case, the value of `$?` will be `1`
(i.e., failure) when the value assigned is falsy. In the second case, the return
value of the assignment is always `0`—unless the assignment expression violates
shell syntax.

Examples

```zsh
(( foo=1+2 ))
printf 'foo=%s; $?=%s\n' $foo $?     # foo=3; $?=0

(( foo=0+0 ))
printf 'foo=%s; $?=%s\n' $foo $?     # foo=0; $?=1

foo=$(( 0+0 ))
printf 'foo=%s; $?=%s\n' $foo $?     # foo=0; $?=0
```

### Operators (in order of precedence)

+ Parentheses: `( )`
+ Unary: `+`, `-`, `!`, `~`; also `++` and `--` (pre/post increment/decrement)
+ Exponentiation: `**`
+ Multiplication, Division, and Modulus: `*`, `/`, and `%` (division is integer
  division; modulus follows C-style and takes the sign of the left-hand operand)
+ Addition and Subtraction: `+` and `-`
+ Bit shifts: `<<` and `>>`
+ Comparisons: `<`, `<=`, `>`, `>=`, `==`, and `!=`
+ Bitwise: `&`, `^`, and `|`
+ Logical: `&&` and `||`
+ Ternary: `condition ? true_val : false_val`
+ Assignment and assignment operations: `=`, `+=`, `-=`, `*=`, `/=`, and `%=`

**Examples:**
```zsh
# Note order of operations.
result=$(( 5 + 3 * 2 ))        # 11

# No $ needed inside (()).
a=10; b=3
print $(( a / b ))             # 3 (integer division)
print $(( a % b ))             # 1 (remainder)

(( count++ ))
(( ++count ))

if (( score >= 90 )); then
    echo "A grade"
fi

# Ternary operator
grade=$(( score >= 90 ? 1 : 0 ))
```

## Variables

| Description                                                    | Syntax                   |
| -----------                                                    | ------                   |
| Set a variable to a simple value                               | `name=something`         |
| Set a variable to the result of an arithmetic expression       | `name=$((expression))`   |
| Set a variable to the result of a command                      | `name=$(command)`        |
| Access a variable's value                                      | `$name` or `${name}`     |
| Use default value if variable is unset or empty                | `${name:-default}`       |
| Use default value and assign it if variable is unset or empty  | `${name:=default}`       |
| Use alternative value if variable is set and non-empty         | `${name:+alternative}`   |
| Error if variable is unset or empty                            | `${name:?error_message}` |
| Delete a variable                                              | `unset name`             |
| Get the list of all defined variables as an associative array  | `${(k)parameters[@]}`    |
| Get the value of a variable whose name is in another variable  | `${(P)name}`             |

The variable `$parameters` is a special associative array that the shell
provides for introspection. (I will discuss introspection later.)

The syntax `${(P)name}` allows for indirect reference to the value of
a variable. It is in fact, the dreaded "variable of a variable." Use it with
caution.

```zsh
user="alice"
alice_score=95
score_var="${user}_score"
actual_score=${(P)score_var}    # Gets value of alice_score, namely 95.
```

## Functions

### Defining Functions

Zshell accepts three ways to define functions.

```zsh
# Type 1: POSIX-compatible syntax
foo() {
    # ...
}

# Type 2: ZSH-specific syntax
function foo {
    # ...
}

# Type 3: ZSH-specific blend of types 1 and 2
function foo() {
    # ...
}
```

There is no difference in meaning or use for these three styles—provided that
you are in Zshell. However, only the first style is accepted by POSIX shells.

Recommendation: pick one of the first two types and use it consistently. They
are both fine, but consistency always helps. The third type requires more typing
than the second for no benefit. So, let's ignore it.

You can delete a function with `unset`. E.g., `unset -f foo`. Zshell keeps track
of all currently defined functions in the `$functions` associative array.

### Working with Parameters

| Action                                            | Syntax                            |
| ------                                            | ------                            |
| Get a parameter                                   | `$1` (second is `$2`, etc.)       |
| Get the name of the function itself               | `$0`                              |
| Expand all parameters as one word (joined by IFS) | `$*`                              |
| Expand all parameters as separate words           | `$@`                              |
| Get the number of parameters (not counting `$0`)  | `$#`                              |
| Remove the first parameter from `$@`              | `shift`                           |
| Remove the last parameter from `$@`               | `shift -p` or `set -- ${@[1,-2]}` |

### Local Variables and Scope

Within functions, you can define variables as local with the `local` builtin.
Local variables are undefined outside of the function (or block) where they are
defined. This is the variable's scope. If a local variable has the same name as
a variable in a wider scope, the local shadows the outer variable within the
function (or block).

Local variables are not exported to the child
processes by default. If you need a local variable in a child process, you can
use `local -x` or `local` and then `export`.

```zsh
local -x foo=bar
local fizz=buzz
export fizz
```

### Autoloading Functions

The `autoload` builtin registers a function for lazy loading. The name of the
function is defined immediately, but the definition is empty until the function
is first called.

There are two important flags for `autoload`.

+ `-U` — Ignore aliases when loading the function.
+ `-z` — Use zsh-style syntax (rather than ksh).

The command `autoload -Uz funcname` looks for a function definition in a file
named "funcname" in the shell's `$fpath`. The file should contain the body of
the function without the surrounding `function funcname { ... }` or `funcname()
{ ... }` wrapper.

## Aliases

| Description                                      | Syntax                                    |
| -----------                                      | ------                                    |
| Display the list of all defined aliases          | `alias`                                   |
| Get the list of all defined aliases, as an array | `${(k)aliases[@]}`                        |
| Define an alias                                  | `alias name="command arg1 arg2 arg3 ..."` |
| Remove an alias                                  | `unalias name`                            |
| Get the arguments, with escaped spaces           | `${@:q}`                                  |

**TODO**: continue from here.

## Conditionals

[A word on conditionals](#a-word-on-conditionals)

Syntaxes:

```zsh
# 1)
if expression
then
    # instructions
fi

# 2)
if expression; then
    # instructions
fi

# 3)
if expression; then ...; fi

# 4)
if expression; then
    # instructions
else
    # instructions
fi

# 5)
if expression; then
    # instructions
elif expression
    # instructions
else
    # instructions
fi
```

| Description                                                     | Syntax                       |
| --------------------------------------------------------------- | ---------------------------- |
| Check if a string is empty or not defined                       | `if [[ -z $VARNAME ]];`      |
| Check if a string is defined and not empty                      | `if [[ -n $VARNAME ]];`      |
| Check if a file exists                                          | `if [[ -f "filepath" ]];`    |
| Check if a directory exists                                     | `if [[ -d "dirpath" ]]; `    |
| Check if a symbolic link exists                                 | `if [[ -L "symlinkpath" ]];` |
| Check if a shell option is set                                  | `if [[ -o OPTION_NAME ]];`   |
| Check if two values are equal                                   | `if [[ $VAR1 = $VAR2 ]];`    |
| Check if two values are different                               | `if [[ $VAR1 != $VAR2 ]];`   |
| Check if a number is greater than another                       | `if (( $VAR1 > $VAR2 ));`    |
| Check if a number is smaller than another                       | `if (( $VAR1 < $VAR2 ));`    |
| Check if a command exits successfully (exit code `0`)           | `if command arg1 arg2 ...`   |
| Check if a command doesn't exit successfully (exit code != `0`) | `if ! command arg1 arg2 ...` |

Remember that the `$` prefix is optional in arithmetic expressions.
Thus, these are equivalent: `(( name1 < name2))` and `(( $name1 < $name2 ))`.

See the [ZSH manual][conditional-expressions] for a full list of conditional
expressions.

[conditional-expressions]: http://zsh.sourceforge.net/Doc/Release/Conditional-Expressions.html

## Loops

Syntaxes:

```zsh
# 1)
for itervarname in iterable
do
    # instructions
done

# 2)
for itervarname in iterable; do
    # instructions
done

# 3)
for itervaname in iterable; do ...; done
```

| Description                                                              | Syntax                     |
| ------------------------------------------------------------------------ | -------------------------- |
| Iterate over a range (inclusive)                                         | `for i in {from..to};`     |
| Iterate over a list of filesystem items                                  | `for i in globpattern;`    |
| Iterate over a list of filesystem items, fail silently if no match found | `for i in globpattern(N);` |

## Examples cheat sheet

Return a value from within a function:

```zsh
function add() {
    local sum=$(($1 + $2))
    echo $sum
}

function add_twice() {
    local sum=$(add $1 $2) # get the callee's STDOUT
    local sum_twice=$(add $sum $sum)
    echo $sum_twice
}

echo $(add 2 3) # 5
echo $(add_twice 2 3) # 10
```

## A word on conditionals

Conditionals use expressions such as in `[[ -z $VARNAME ]]`. These can also be
used in `while` loops, as well as be used outside of blocks:

```zsh
[[ -z $VARNAME ]] && echo "VARNAME is not defined or empty!"
[[ -f $FILEPATH ]] && echo "File exists!"
```

This works because conditional expressions (`[[ ... ]]` and `(( ... ))`) don't
actually return a value; they behave like commands and as such set the status
code to `0` if the condition is true, or `1` else.

If we want to display the message only if the condition is falsy:

```zsh
[[ -z $VARNAME ]] || echo "VARNAME is not empty!"
[[ -f $FILEPATH ]] || echo "File does not exist!"
```

## Introspection

| Item                               | Description                               |
|------                              |-------------                              |
| `${(k)parameters[@]}`              | All variable names                        |
| `$parameters[varname]`             | Type/attributes of specific variable      |
| `${(k)functions[@]}`               | All function names                        |
| `$functions[funcname]`             | Definition of specific function           |
| `${(k)aliases[@]}`                 | All alias names                           |
| `$aliases[aliasname]`              | Definition of specific alias              |
| `${(k)options[@]}`                 | All shell option names                    |
| `$options[optname]`                | Current state of specific option (on/off) |
| `$modules`                         | Loaded zsh modules                        |
| `$history`                         | Command history array                     |
| `$jobstates`                       | Background job information                |
| `$dirstack`                        | Directory stack contents                  |
| `typeset`                          | Shows all variables with types            |
| `typeset -f`                       | Shows all function definitions            |
| `typeset -p varname`               | Shows how variable was declared           |
| `which command` / `whence command` | Shows what command resolves to            |
| `${+varname}`                      | Test if variable exists (returns 1/0)     |
| `[[ -v varname ]]`                 | Alternative variable existence test       |
| `$ZSH_VERSION`                     | Zsh version string                        |
| `$OSTYPE`                          | Operating system type                     |
| `$MACHTYPE`                        | Machine/architecture type                 |
| `$SHLVL`                           | Current shell nesting level               |
