# Effective C: Chapter 1 Getting Started With C

Seacord begins with a "Hello, world!" program, but with a few new features.
First, Seacord includes stdlib as well as stdio. (We will see why in a minute.)
Second, he uses `puts` instead of `printf`. Third, he exits the program with
`EXIT_SUCCESS` from stdlib instead of `0`. (I like `EXIT_SUCCESS` since the
variable explicitly signals success, but `0` does not.)

Seacord also offers a second version of the program that checks for an error
from `puts`:

```c
int
main(void)
{
    if (puts("Hello, world!") == EOF) {
        return EXIT_FAILURE;
    }
    return EXIT_SUCCESS;
}
```

Seacord also briefly discusses five sources of portability problems:
implementation-defined behavior, unspecified behavior, undefined behavior,
locale-specific behavior, and common extensions. Implementation-defined
behaviors are not a major source of trouble, and you can mark them clearly with
a `static_assert` declaration, which Seacord promises to cover in Chapter 11.
Locale-specific behavior are what they sound like: behavior specific to given
nations, cultures, or languages. Common extensions are also what they sound
like. I am not worried about either of these since I write C for myself only.
Unspecified behavior exists when the C standard offers two or more options for
something. Undefined behavior is when the C standard doesn’t specify at all what
should happen in a given case. I should avoid both unspecified and undefined
behavior entirely.
