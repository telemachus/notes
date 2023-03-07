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

[slog]: https://pkg.go.dev/golang.org/x/exp/slog
