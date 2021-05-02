# Introducing Go, Chapter 2: Types

Doxsey warms us up to Go’s use of types by talking about types and tokens in philosophy and sets in math. Philosophers and mathematicians use these terms to talk about properties that things have in common and in order to group (relevantly) similar things together. Programming languages use types for the same reasons.

Go is statically typed. All variables have a specific type that does not change. Many people find static typing annoying or difficult at first, but static typing can help in the long run.

## Numbers

Go has several different numeric types. Most importantly, we can distinguish integers and floating-point numbers.

### Integers

Integers are whole numbers: numbers with no decimal part. They can be positive or negative, and include the number 0. Since computers require space to store numbers, computers use many different sizes of numbers. Go provides several types of integers: uint8, uint16, uint32, uint64, int8, int16, int32, and int64. The series starting with “u” are “unsigned.” This means that there are only positive numbers and zero. The numbers correlate with how many bits of storage each type uses. The size of a type determines the range of numbers that you can store in a variable of that type. For example, int8 uses eight bits of memory, and you can store from -128 to 127 in a variable of type int8. Finally, there are two aliases for integer types: byte is the same as uint8, and rune is the same as int32.

In general, use int unless you have a special reason to use something else.

### Floating-Point Numbers

Floating-point numbers contain a decimal part. Again, Go provides multiple types of floating-point numbers of different sizes: float32 and float64. You should use float64 unless you have a special reason to use something else.

## Strings

A string is a sequence of characters stored as individual bytes. (Beware: a single character can be stored in one, two, three, or four bytes!) You create string literals using double quotation marks or backticks. (Backticks allow you to have strings that contain newlines. Also double-quoted strings allow special backslash escapes (e.g., `\n` and `\t`).

Here are some common string operations:

+ `len("Hello world")` returns the length of the string
+ `Hello, world"[1]` accesses the second character in a string as a byte (i.e., 101 in this case)
+ "Hello, " + "world" concatenates two strings together

## Booleans

Go provides two booleans (true and false) and three boolean operators (&&, ||, and !).
