# Notes on `slog`

[`slog`][slog] is a structured logger that will likely become part of the Go
standard library when Go 1.21 is released.  In general, `slog` produces log
records that include "a time, a level, a message, and a set of key-value
pairs, where the keys are strings and the values may be of any type."

If you set a default logger (using `slog.SetDefault(logger)`), then you get
access to package-level convenience methods such as `slog.Info`, `slog.Debug`,
`slog.Warn`, and `slog.Error`.  If you use these convenience methods, then you
don't need to pass the desired log level as an argument.  On the other hand,
you also always have access to `slog.Log`, where you can pass the log level as
an argument.  E.g., `slog.Log(ctx, slog.LevelError, "message", ...)`.
(I'll say more about the context argument later.)

Note that some of the logging output receives special treatment.  First, the
time.  The library outputs a key-value pair with time information by default
as the first item in the log record.  (In fact, if you want to suppress the
time output, you must do a non-trivial amount of work.  I'll show how later.)
Next, the logger includes a `level=LEVEL` key-value pair.  Again, the user does
not explicitly pass `"level", "LEVEL"` as explicit arguments.  Either, the
caller uses a named convenience method (e.g., `slog.Info`) or the caller
includes a `slog` constant that represents the level.  Callers include
a message argument, but they don't explicitly pass `"msg", "message"`.
Finally, if you use the `Error` convenience method, then the error argument is
passed alone rather than as a key value pair.  (I.e., you don't pass `"err",
"error"`.)

## Architecture

The package defines a `Logger` type that has various methods defined on it.
When you call a method on a `Logger`, it creates a `Record` and passes that
`Record` to a `Handler`.  Each `Logger` is associated with a specific type of
`Handler`.  There are three built-in handlers available in `slog`.  First, if
you do not define a handler, then the default handler creates a string from
a `Record` and passes it to the (older) standard library `log` package.  The
result is semi-structured log output.  E.g., `2022/11/08 15:28:26 INFO hello
count=3`.  In addition, `slog` provides a JSON handler and a fully structured
text handler.  The text handler outputs structured pairs in the form
`key=value`, and the JSON handler outputs a single JSON record of `key=value`
pairs.

You create a text or JSON handler with calls to `slog.NewTextHandler` or
`slog.NewJSONHandler`.  You can configure the handler using `HandlerOptions`.
`HandlerOptions` is a struct with three public fields.  `AddSource` is
a boolean, and defaults to false.  If you set `AddSource` to true, then each
log record will include information about the source code of log calls.  (Yes,
there can be a problem here with indirection, wrappers, and helper methods
that call a log.)  The `Level` field should be set with the minimum record
level to be logged.  If the `Level` field is `nil`, then it defaults to
`slog.LevelInfo`.  Finally, the `ReplaceAttr` field takes a function that
allows you to adjust (or remove) attributes in a `Record` before the handler
prints out the `Record`.  (I'll give examples of these later.)

Methods of loggers that handle records work on what `slog` calls an `Attr`.
An `Attr` is a struct with `Key` and `Value` fields.  The `Key` field is
a string, and the `Value` field has a `Value` type.  `Value`s like `any` can
represent any type in Go.  In addition, `slog` can represent many Go values
without allocation.

You can pass key-value pairs to logger methods as alternating items, or you
can use constructors for greater efficiency.  For example, the following yield
the same output, but the second is more efficient.

```go
slog.Info("hello", slog.Int("count", 3))
slog.Info("hello", "count", 3)
```

`slog` also provides ways to include attributes every time you log
(`slog.With`) and to group a number of attributes (`slog.Group` and
`slog.WithGroup`).

Levels rank the severity of log events, and they also control what events
are actually logged.  By default, new handlers are set to `Info` level.  This
means that anything of `Info` level or higher will be logged, but anything
below `Info` level (e.g., `Debug` level events) will not be logged.  The
levels themselves are integers.  `slog` defines several standard levels
(`Debug`, `Info`, `Warn`, and `Error`), but applications can define additional
levels as well.  Normally, a logger has a constant level throughout the
lifetime of a program, but you can also get a dynamic level using `LevelVar`.

`slog` also provides methods that accept contexts.  First, the `Logger.Log`
and `Logger.LogAttrs` methods take a context as their first parameters.  Next,
methods like `slog.Info` also have `slog.InfoCtx` alternatives.  These too
take a context as their first parameters.

## Customizing `slog`

There are several ways to customize `slog`.  In order of least to most effort,
you can use a `LogValuer` to tweak values, you can wrap output methods, and
you can write a custom handler.  I'll talk a little about each of these below.

### `LogValuer`

Imagine you want to redact certain values in your logs (e.g., something with
security considerations).  The easiest way to do this is to wrap the value in
a type that implements the `LogValuer` interface.  This interface requires
only a single method, `LogValue`.  The `LogValue` method takes no arguments
and returns a single `slog.Value`.  Here's a simple example from `slog`'s
documentation.

```go
// A token is a secret value that grants permissions.
type Token string

// LogValue implements slog.LogValuer.
// It avoids revealing the token.
func (Token) LogValue() slog.Value {
	return slog.StringValue("REDACTED_TOKEN")
}
```

If you pass a token to your logger `REDACTED_TOKEN` will appear rather than
the actual value of the token.

`LogValue` can return Values that also implement `LogValuer`.  `slog` provides
`Value.Resolve` to prevent infinite loops or runaway recursion.  (The
documentation suggests using `Value.Resolve` if you write a handler.

### Wrapping Output

You can easily wrap output methods, but if you do, you need to think about the
depth in the stack of the eventual call.  As the documentation explains, the
following code has a bug.

```go
func Infof(format string, args ...any) {
	slog.Default().Info(fmt.Sprintf(format, args...))
}
```

If this implementation is in `mylog.go`, and you call `Infof` from `main.go`,
`slog` reports the source file (and line) from `mylog.go` rather than
`main.go`.

To fix this problem, you need to construct a new record and manually adjust
the position in the stack where the source code lives.  Here's an example that
fixes this bug.

```go
// Infof is an example of a user-defined logging function that wraps slog.
// The log record contains the source position of the caller of Infof.
func Infof(format string, args ...any) {
	l := slog.Default()
	if !l.Enabled(context.Background(), slog.LevelInfo) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelInfo, fmt.Sprintf(format, args...), pcs[0])
	_ = l.Handler().Handle(context.Background(), r)
}
```

## Writing a Handler

The `Handler` interface requires four methods.

```go
type Handler interface {
	Enabled(context.Context, Level) bool
	Handle(context.Context, Record) bool
	WithAttrs(attrs []Attr) Handler
	WithGroup(name string) Handler
}
```

Some important notes about handlers.

+ Any handler method can be called concurrently with itself or another handler
  method.  Handlers should protect against data races.
+ The context arguments in Enabled and Handle may be `nil`.  Handlers must not
  do anything with that argument that may cause a panic.
+ The `Enabled` method should be called early.  If the Level of a record is
  lower than the Level argument, then the record is ignored and no logging
  occurs.  By calling `Enabled` early, we prevent any waste on a record that
  will never be used to log anything.
+ `Enabled` can use values from the context argument to make a decision about
  whether or not the logger is currently enabled.
+ `Handle` will only be called if `Enabled` returns true.
+ Again, the context argument passed to `Handle` can provide values that may
  change how the record appears or is treated.
+ However, even if the context is canceled, the record should be processed.
+ If a record's Time is the zero time, then time should not appear in output.
+ If a record's PC is zero, then source file and line should not appear in
  output.
+ If an Attr's key is an empty string, and the value is not a group, the Attr
  should be ignored by the `Handle` method.
+ If a group has an empty key, the group's Attrs should appear in the record
  as non-group Attrs.  (I think that this is what "inline the group's Attrs"
  means.  I should double check this in the source.)
+ If a group has no Attrs, then ignore it (even if it has a non-empty key).
+ I am unsure what callers are supposed to do with the error from `Handle`.
+ `WithAttrs` returns a new Handler that has all the attributes of the
  receiver, plus the arguments to `WithAttrs` itself.
+ The new Handler returned by `WithAttrs` owns the slice of Attrs that it
  receives.  As such, it may change that slice.
+ `WithGroup` returns a new Handler; the group name given as an argument is
  added to the end of the list of groups that the receiver already has.

[slog]: https://pkg.go.dev/golang.org/x/exp/slog
