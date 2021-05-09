# Introducing Go, Chapter 8: Packages

Go uses packages to organize code and allow easy code reuse. Doxsey mentions three advantages of packages.

1. They reduce the chance of overlapping names, and they help keep names short and succinct.
1. They organize code and make it easy to find things you want to use.
1. They speed up the compiler because you never have to recompile items from packages that have already been compiled (e.g., `fmt` or `strings`).

## The Core Packages

### Strings

Go provides a strings package with many useful functions. Doxsey mentions the following functions.

+ `strings.Contains(s, substr string) bool` tests whether a string contains a substring. E.g., `strings.Contain("foo", "f")` returns `true`.
+ `strings.Count(s, sep string) int` tells you how many times a string is found in another string. E.g., `strings.Count("test", "t")` returns 2.
+ `strings.HasPrefix(s, prefix string) bool` tells you whether a string starts with a specified prefix.
+ `strings.HasSuffix(s, suffix string) bool` tells you whether a string ends with a specified suffix.
+ `strings.Index(s, sep string) int` looks for a desired string in a target string. If the desired string is found, `Index` returns index position where the desired string begins. If the desired string is not found in the target string, then `Index` returns -1.
+ `strings.Join(ss []string, sep string) string` takes a slice of strings and returns a single string formed by placing `sep` between each member of the slice of strings.
`strings.Split(s, sep string) []string` is the reverse of `strings.Join`. It takes a string to split and a separator to split on. It returns a slice of split items. If the separator is not found, the slice will contain only one item: the original string.
+ `strings.Repeat(s string, count int) string` takes a string and returns a string formed by repeating `s` `count` times.
+ `strings.Replace(s, old, new string, n int) string` replaces `old` with `new` at most `n` times. If you want to replace every occurrence of `old` with `new`, no matter how many times, then use -1 for `n`.
+ `strings.ToLower(s string) string` and `strings.ToUpper(s string) string` do what you would expect them to do.

If you need to convert a string to a slice of bytes or a slice of bytes to a string, you use type coercion.

```go
byteSlice = []byte("test")
str = string([]byte{'t', 'e', 's', 't'})
```

### Input/Output

The `io` package provides two interfaces that you will see all over Go code: `Reader` and `Writer`. Types that implement the `Reader` interface support the `Read` method, and types that implement the `Writer` interface support the `Write` method. The `io.Copy` function uses these interfaces: `io.Copy(dst Writer, src Reader) (written int64, err error)`.

You can use a `bytes.Buffer` to read or write to a a slice of bytes or a string.

```go
var buf bytes.Buffer
buf.Write([]("test"))
// later, if you need the contents of the buffer as a slice of bytes
buf.Bytes()
```

If you only need to read from a string, use `strings.NewReader`. It's more efficient than a buffer, if you don’t need to support writing as well.

### Files and Folders

Doxsey shows you how to open a file and read it into a slice of bytes using `os.Open` and `file.Read`. He also shows you how to use `ioutil.ReadFile` (which is now `os.ReadFile`) to read an entire file without the extra step of opening it. (You also don’t need to worry about how large the file is, as you do with `file.Read`.)

You can get the contents of a directory using `os.Open` with a directory path instead of a file. You then use `Readdir` to list the contents. If you want to recursively walk through a folder, you can use `path/filepath.Walk`.

You can create a file using `os.Create` and then write to the file with `file.WriteString`.

### Errors

Go provides a built-in error type (`error` is the obvious name!). You can create your own errors using `errors.New`.

```go
err := errors.New("my very own custom error message")
```

### Containers and Sort

Go provides several types of collections in its `container` package. Doxsey gives an example of doubly linked lists from `container/list`. There’s also `container/heap` and `container/ring`.

See `ch08/simple-tail.go` for a naive version of `tail` in Go, using `container/ring`.

Doxsey also demonstrates how to write a custom sort function for a slice of whatever type you like. You need to create an interface for your type and implement three methods for your interface: `Len`, `Less`, and `Swap`. Once you do that, you can easily sort your slice. Here’s an example for me with students.

```go
type Student struct {
    Email string
    FirstName string
    LastName string
}

type ByLastName []Student

func (ss ByLastName) Len() int {
    return len(ss)
}

func (ss ByLastName) Less(i, j int) bool {
    return ss[i].LastName < ss[j].LastName
}

func (ss ByLastName) Swap(i, j int) {
    ss[i], ss[j] = ss[j], ss[i]
}
```

### Hashes and Cryptography

Go provides both cryptographic and non-cryptographic hash functions. The non-cryptographic ones are in `hash/xxx` and the cryptographic ones are in `crypto/xxx`. Doxsey gives a brief example of one of each.

## Servers

Doxsey shows how easily we can create servers and clients in Go.

### TCP

Go provides TCP clients and servers in its `net` package. You create a server using `net.Listen`, which takes two arguments: a network type (e.g., `tcp`) and a port to bind (e.g., `:9999`). `Listen` returns a `net.Listener` interface, which implements the following functions.

+ `Accept() (c Conn, err error)` waits for a client to connect and returns a `net.Conn`. `net.Conn` implements the `io.Reader` and `io.Writer` interfaces. We use the connection like a file: we read from it and write to it.
+ `Close() error` closes a listener.
+ `Addr() Addr` returns a listener’s network address

When we have a listener, we call `Accept` and wait for a connection. We then hand the connection off to a handler that receives and handles any incoming message and closes the connection.

I adapted Doxsey’s example to make a client that accepts messages from standard input and sends them to the server. If the message is “quit,” the client shuts down, and the program ends.

### HTTP

Go also provides `net/http` for HTTP servers. You use `http.HandleFunc` with two arguments: a string address and a function that accepts an `http.ResponseWriter` and a `*http.Request`. You can also use `http.FileServer` to serve static files. I don’t fully understand how to put these pieces together, but I can read a longer introduction on servers when I need it.

#### RPC

You can use `net/rpc` and `net/rpc/jsonrpc` to allow clients to call methods through a network connection. I see how this works, but I don’t have a sense of *why* I would want to “expose methods so they can be invoked over a network (rather than just in the program running them).”

### Parsing Command-Line Arguments

Go provides `flag` if you need to offer flags for your program.

## Creating Packages

You can create and easily import packages in Go. Name your package whatever you like via `package`, place it in your `$GOPATH`, and then import via `import "path/to/my/package."` You use the package as `package.Function` or `package.Struct` or `package.Constant`.

Go has simple rules for visibility. If you want to export something in your package, start its name with an uppercase letter. If you want to keep something in your package private, start its name with a lowercase letter.

You can give packages aliases when you import them. Use the structure `import alias "path/to/package"`. Then you can refer to the package as `alias`.

## Documentation

The `godoc` tool automatically has access to the documentation for any installed package. You can get documentation for a function by typing `go doc path/to/package Function`. (Note that the version `godoc path/to/package Function` no longer works.)

It is good practice to document every public function, structure, and constant in a package.
