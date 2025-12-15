# *The C Programming Language*, Kernighan and Ritchie

## Chapter 5: Pointers and Arrays

K&R start the chapter with a succinct definition.
"A pointer is a variable that contains the address of a variable" (93).
They explain that C code makes heavy use of pointers and that pointers and arrays are "closely related" (93).
They acknowledge that pointers, like `goto`, can lead to complex code, but they insist that pointers can be used well.

### Pointers and Addresses

Imagine that memory in a computer is a row of cells, each of which has an address.
Different data types require a different number of cells to store their values.
E.g., a short integer requires fewer cells than a double or a long.
A pointer stores the address of another item as its value or data.

The unary operator `&` returns the address of an object.
Thus, you can create a pointer as follows: `p = &c;`.
The variable `p` is now a pointer to the address of the value of `c`.
Thus, `p` points to `c`.

The unary operator `*`, the *indirection* or *dereferncing* operator returns the value of a pointer.
Thus, `*p` returns the value of `c`.
The indirection operator is also used in declarations.
E.g., `int x = 1; int *ip; ip = &x;`.
This can be confusing since you use `*` both to declare pointers and to get the value of a pointer.
These are very distinct uses.
NB: some people write pointer declarations with the indirection operator at the end of the type rather than the start of the variable name.
E.g., `int* ip;`.

Every pointer has a specific type.
That is, you cannot store a double in something declared as `int *ip;`.
However, you can cast items to the generic type *pointer to void* `(void *)`.

Once you have a pointer, you can use it anywhere that a variable or constant of that type can be used.
E.g., `int x = *ip;` instead of `int x = 0;`.

There are some fiddly rules concerning the precedence of `&` and `*`.
Most importantly, be careful when using the indirection operator together with incrementing or decrementing.
E.g., you need to write `(*ip)++` or `++*ip` to increment the value in `ip`.
If you write, `*ip++`, you will increment `ip` itself rather than its value.
Why?
Because both `++` and `*` associate right to left.

### Pointers and Function Arguments

Since C passes arguments to functions by value rather than by reference, you need pointers to change variables through functions.
For example, `swap(int x, int y)` cannot actually swap the values of the arguments.
Instead, you need to use `swap(int *x, int *y)` and to call the function as `swap(&a, &b)`.
Internally, the function would look like this.
```c
void swap(int *px, int *py)
{
	int temp = *px;
	*px = *py;
	*py = temp;
}
```

### Pointers and Arrays

Pointers and arrays have significant (and initially surprising) overlap.
In particular, "Any operation that can be achieved by array subscripting can also be done with pointers" (97).
Imagine the following: `int a[10]; int *pa = &a[0]`.
The name of an array is "a synonym for the location of" its first element.
So you can also write that as `int a[10]; int *pa = a;`.
At this point, `pa` points to the first item in `a`.
By definition `pa+1` points to the second item in `a`.
More generally, you can use `*(pa+n)` to get the value of `a[n]`.
In addition `&a[n]` and `a+n` are also equivalent.
The expression `a+n` is the address of the nth element beyond `a`.

However, there is one crucial difference between an array name and a pointer.
Since a pointer is a variable, you can assign to it, and you can increment or decrement.
You cannot assign to an array name or increment or decrement them.

The programmer does not need to worry about the type or size of the variables in an array.
That is, `pa+1` always points to the second item in the array regardless of what the array stores.

However, also note that C does not enforce array boundaries.
If the array has only two members, and you do `pa+2`, C will merrily point to the next memory location after the end of the array.

### Address Arithmetic

"C is consistent and regular in its approach to address arithmetic" (100).
When `p` is a pointer to an element of an array, `p++` points to the next
element, and `p+=i` points `i` elements after what `p` points to.

They demonstrate this consistency with a simplistic storage allocator.
(The allocator is simplistic because you have to call `afree` in LIFO order.)

```c
#define ALLOCSIZE 10000

static char allocbuf[ALLOCSIZE]; // We will use this array for storage.
static char *allocp = allocbuf;  // The next free position is initially 0.

char *alloc(int n)
{
    if (allocbuf + ALLOCSIZE - allocp >= n) { // There is room for n more.
        allocp += n;
        return allocp - n;
    } else { // There is not enough storage.
        return 0;
    }
}

void afree(char *p)
{
    if (p >= allocbuf && p < allocbuf + ALLOCSIZE) {
        allocp = p;
    }
}
```

You can only do a limited range of arithmetic on pointers.

+ You can do addition and subtraction of integers on pointers.
+ You subtract two pointers that point to the same array.
+ You can use comparison and equality operators on pointers to the same array: `==`, `!=`, `<=`, `<`, `>=`, `>`.
+ You can compare any pointer for equality (or inequality) with zero.

However, arithmetic or comparison with pointers of different arrays, is undefined.
One small exception: the first element after the end of an array is available for pointer arithmetic.
This makes possible looping and various idioms.

Here are the valid pointer operations.

+ Assignment of pointers of the same type
+ Adding or subtracting a pointer and an integer
+ Subtracting or comparing two pointers to the same array
+ Assigning a pointer zero and comparing a pointer with zero
+ Assigning a pointer of any type to a pointer to void
+ Assigning a pointer of one type to a pointer of another type with a cast

### Character Pointers and Functions

Think about passing a string constant to a function.
E.g., `printf("Hello, World.\n");`.
First, the string constant is an array of characters.
Second, the string constant is terminated internally by the null character, `\n`.
(This is how string functions can find the end of strings without knowing their length.)
When a function receives a string constant, it gets a pointer to the first element of the string.
You can write a function that handles a string constant in array style or pointer style.
For example, here are two versions of `strcopy(s, t)`.

```c
// Array style
void strcpy(char *s, char *t)
{
    int i = 0;
    while((s[i] = t[i]) != '\0') {
        i++;
    }
}

// Pointer style
void strcpy(char *s, char *t)
{
    // We can leave the comparison to zero implicit.
    while(*s++ = *t++) {
        ;
    }
}
```

K&R mention another idiomatic pair of pointer operations.

```c
*p++ = val; // Push val onto a stack.
val = *--p; // Pop top of stack and assign its value to val.
```

### Pointer Arrays; Pointers to Pointers

You can have an array of pointers just like you can have an array of other types.
This can be especially useful if you want to work on a group of strings.
As an example, K&R imagine a program that sorts the lines of a file.
Each line of input is added to an array of pointers to char.
Then the lines are sorted.
Then every line of the array is printed.
Here's what the printing looks like.

```c
void writelines(char *lineptr[], int nlines)
{
    while (nlines-- > 0) {
        printf("%s\n", *lineptr++);
    }
}
```

### Multi-dimensional Arrays

In addition to arrays of pointers, C also has what K&R call "rectangular multi-dimensional arrays" (110).
They mention that arrays of pointers are more common that multi-dimensional arrays.
As an example, they show the following array of arrays, which can be used to track days of the month for regular years and leap years.

```c
static char daytab[2][13] = {
    {0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
    {0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
}
```
