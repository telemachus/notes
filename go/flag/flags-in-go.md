# Handling flags and command line arguments with Go

## Manually

You can handle arguments completely by hand using `os.Args`. `Args[0]` will always be the name of the currently running program, and `Args[1:]` will contain the rest of the arguments as strings.

There are obvious reasons not to do this completely by hand. First, it’s easy to forget about `Args[0]`. Second, you would have to parse apart option combinations like `--foo=bar` or `--foo bar` yourself. Third, you would have to do type conversions yourself. That is, an argument `1` will be the string and not an integer.

## Using `flag` from Go’s Standard Library

Instead of doing everything by hand, you should use `flag` from Go’s standard library. This package provides predefined types of variables, handles parsing standard combinations automatically, and does type conversions for you.

In a nutshell, you need to do the following:

+ `import "flag"`
+ Define the flags and types that you want to work with.
+ Call `flag.Parse()`
+ Handle any non-flag arguments left in `flag.Args()`

Here’s a first simple example.

```go
var nFlag = flag.Int("n", 0, "specify the count of n")
var sFlag = flag.String("name", "", "what is the name?")
var bFlag = flag.Bool("loop", false, "should we loop?")
flag.Parse()

if *bFlag {
    for i := 0; i < *nFlag; i++ {
        fmt.Printf("sFlag = %q\n", *sFlag)
    }
} else {
    fmt.Printf("bFlag = %t, sFlag = %q, and nFlag = %d\n", *bFlag, *sFlag, *nFlag)
}
```

The functions `flag.Int`, `flag.String`, etc. return a pointer to the variable of their respective type.

If you prefer, you can use `Var` functions and the `&` operator to assign values to variables rather than pointers.

```go
var nFlag int
flag.IntVar(&nFlag, "n", 0, "specify the count of n")
var sFlag string
flag.StringVar(&sFlag, "name", "", "what is the name?")
var bFlag bool
flag.BoolVar(&bFlag, "loop", false, "should we loop?")
flag.Parse()

if bFlag {
    for i := 0; i < nFlag; i++ {
        fmt.Printf("sFlag = %q\n", sFlag)
    }
} else {
    fmt.Printf("bFlag = %t, sFlag = %q, and nFlag = %d\n", bFlag, sFlag, nFlag)
}
```

In addition to `Int`, `String`, and `Bool`, Go also provides `Duration`, `Float64`, `Int64`, `Uint`, and `Uint64`.

## More Complex Uses of `flag`

You can use `NewFlagSet` to create subcommands or to test flag settings. I’ll discuss both of these below.
