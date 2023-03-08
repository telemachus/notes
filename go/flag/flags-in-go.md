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

You can use `NewFlagSet` to create subcommands or to test flag parsing. I’ll discuss both of these below.

### Testing Flag Parsing

Eli Bendersky explains [how to test flag parsing](https://eli.thegreenplace.net/2020/testing-flag-parsing-in-go-programs). Normally, we parse flags and then immediately pass the choices to our program. As he puts it, the normal sequence looks like this: flags come in -> parse flags -> do work. If we want to test the step where we parse the flags, we need to interrupt that normal sequence. We want it to look like this instead: flags come in -> parse flags -> store flags in a configuration variable -> do work. We can hook the testing to the configuration variable.

Here’s Bendersky’s code for flag parsing.

```go
type Config struct {
  verbose  bool
  greeting string
  level    int

  // args are the positional (non-flag) command-line arguments.
  args []string
}

// parseFlags parses the command-line arguments provided to the program.
// Typically os.Args[0] is provided as 'progname' and os.args[1:] as 'args'.
// Returns the Config in case parsing succeeded, or an error. In any case, the
// output of the flag.Parse is returned in output.
// A special case is usage requests with -h or -help: then the error
// flag.ErrHelp is returned and output will contain the usage message.
func parseFlags(progname string, args []string) (config *Config, output string, err error) {
  flags := flag.NewFlagSet(progname, flag.ContinueOnError)
  var buf bytes.Buffer
  flags.SetOutput(&buf)

  var conf Config
  flags.BoolVar(&conf.verbose, "verbose", false, "set verbosity")
  flags.StringVar(&conf.greeting, "greeting", "", "set greeting")
  flags.IntVar(&conf.level, "level", 0, "set level")

  err = flags.Parse(args)
  if err != nil {
    return nil, buf.String(), err
  }
  conf.args = flags.Args()
  return &conf, buf.String(), nil
}
```

And here is a simple program that uses `Config` and `parseFlags`.

```go
func doWork(config *Config) {
  fmt.Printf("config = %+v\n", *config)
}

func main() {
  conf, output, err := parseFlags(os.Args[0], os.Args[1:])
  if err == flag.ErrHelp {
    fmt.Println(output)
    // This is Bendersky’s code, but I would use os.Exit(0) since -help is not
    // an error.
    os.Exit(2)
  } else if err != nil {
    fmt.Println("got error:", err)
    fmt.Println("output:\n", output)
    os.Exit(1)
  }

  doWork(conf)
}
```

Finally, here is what testing such flag parsing looks like.

```go
func TestParseFlagsCorrect(t *testing.T) {
  var tests = []struct {
    args []string
    conf Config
  }{
    {[]string{"-verbose"},
      Config{verbose: true, greeting: "", level: 0, args: []string{}}},

    {[]string{"-level", "8", "-greeting", "joe", "-verbose", "foo"},
      Config{verbose: true, greeting: "joe", level: 8, args: []string{"foo"}}},

    // ... many more test entries here
  }

  for _, tt := range tests {
    t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
      conf, output, err := parseFlags("prog", tt.args)
      if err != nil {
        t.Errorf("err got %v, want nil", err)
      }
      if output != "" {
        t.Errorf("output got %q, want empty", output)
      }
      if !reflect.DeepEqual(*conf, tt.conf) {
        t.Errorf("conf got %+v, want %+v", *conf, tt.conf)
      }
    })
  }
}
```
