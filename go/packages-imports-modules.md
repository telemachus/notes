# [Packages, Imports and Modules in Go][pim], Alex Edwards

## Packages

A package is one or more files that share the same package name.  All the
files for one package must be in the same directory.  In general, one
directory can contain only one package.  The exception (that proves the rule)
is testing code.  You can have `package foo` and `package foo_test` in the
same directory.  But you cannot have `package foo` and `package bar` in the
same directory.

Files that are in the same package can see types, constants, variables, and
functions from any other file in the same package.  Within a package, all of
this is shared without concern for public or private visibility.

## The `main` package

The `main` package is special.  It signals to Go's compiler and runtime that
the package produces an executable.  A `main` package must contain a `main`
function in one of its files.  By convention, the `main` function should live
in a file called `main.go`, but this is not a rule that Go enforces.

You cannot run or build a package other than `main`.  This is a rule that Go
enforces.

## Importing and using standard library packages

You can import and use things from the standard library if you have a full Go
installation.  (In fact, I am not sure that you can even have a functioning Go
installation without the standard library.  TODO: find out.)

To import packages from the standard library, you use the full path to the
package in the standard library.  This will often be a one-word path (e.g.,
`fmt`), but it can often be a longer path (e.g., `net/http/httptest` or
`net/http`).

However, after you import the package, you can use material in the package
with just the final element from the import path.  This saves typing.  In
other words, you import `slog` as `log/slog`, but you use it simply as `slog`.

## Unused and missing imports

You cannot import a package without using something from it.  If you try, Go
will not run or compile your code.  You will get an "imported and not used"
error.

In the same way, you will receive an error at compile-time if you try to use
a package that you have not imported.  In this case, the error will tell you
that something is "undefined."

You can use `goimports` in conjunction with your editor to automatically edit
imports so that they are correct.  However, in some cases, package names are
ambiguous (e.g., `html/template` or `text/template`?).  In these cases, make
sure that `goimports` adds the correct import path.

## Exported vs unexported

Within a given package some things are exported and some things are
unexported.  How can we tell the difference?  Very simply, if an identifier
begins with a capital letter, it is exported.  Otherwise, it is unexported.

Unexported things are private to the package, and exported things are public.
You should export as little as possible initially because of [Hyrum's
law][hyrum].  (Briefly, Hyrum's law says that anything a user can get at
becomes part of your API.  You may only promise X, Y, and Z, but if users can
get at A, B, and C, then A, B, and C are also part of your API.  If you remove
or significantly change A, B, and C later, you will break things for that
user.  See also [XKCD's explanation][hyrum-xkcd].)  Since main is not normally
imported, you shouldn't add exported things within a `main` package.

## Modules

The Go Wiki [defines a module][wiki-mod] as "a tree of Go source files with
a `go.mod` file in the tree's root directory."  A simple module may have only
one directory and one package; more complex modules may have many directories
and several packages.

The module is identified by its path inside of the `go.mod` file.  In general,
you should use a URL that you control for the module path of any module that
you wish to make public.  If you are just noodling around or creating an
example, use `example` or `test` for the name of the module.  Go [promises
never to use those in the standard library][promise].

## Using multiple packages in your code

First, let's lay out two norms for Go code.

+ One package means one directory.  As I discussed above, all the code for
  a single package belongs in one directory.  (This is a rule that Go's
  compiler and runtime enforce. There may be ways to get around this—by
  mucking with the compiler—but you should never do it even if you could.)
+ Other than `main` packages, you should generally name the package after the
  directory (or the directory after the package).  This isn't a rule that Go
  enforces in any way, but you should do it anyhow. For further advice on
  naming modules, take a look at [this post][mod-naming].

You should avoid creating overly complicated, multi-package modules too early.
Often, all you need is one package for one module.  You can separate code into
as many files as you like, and you may never need more than one package.

That said, multiple packages can be handy.  For example, if you have code that
contains both a library and an executable, you can separate the library from
the executable easily by putting them in distinct packages.  Also, if you have
code that you do not want users to be able to import, you can put that into an
`internal` sub-path.  Anything there is only available for the main module
itself.  Users cannot get at packages in `internal`.  (This helps to protect
you from Hyrum's law.)

A few notes about `main` packages.  First, they do not have to live in the
module's root.  Second, you can have one module with multiple `main` packages.
Edwards gives the example of something that ships with both a commandline
application and a web application.

## Importing and using third-party packages

Before you can use third-party packages, you have to download them.  You can
use `go get` to download packages.  `go get` will automatically install
dependencies of what you request.

To import third-party packages, you need to use the full module path for the
package.  Normally, this will be a URL where the item lives and which you give
to `go get` to download the package.

When third-party packages have version numbers, you will normally omit that
part of the module part when you use the package in your code.  To clarify,
you use the full module path when you run `go get` and when you import code.
But when you want to get at variables, constants, and functions in the
package, you omit the version portion of the path.  As an example, you would
run `go get github.com/go-chi/chi/v5`, and you would import that same path,
including `/v5`.  But when you use the package, the identifier is simply `chi`
rather than `chi/v5`.

## Organizing import statements

There is no exact rule for how to organize import statements, and the Go
community has not yet reached a consensus about how to do this.  Edwards
offers his favorite style.

```go
import (
	// Put standard library packages first.
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"sync"
	
	// Put packages from the current module second.
	"example.com/internal/logger"
	"example.com/internal/smtp"
	
	// Put third-party packages third.
	"github.com/go-playground/form/v4"
	"github.com/spf13/pflag"
	
	// Put aliased packages last.
	_ "github.com/mattn/go-sqlite3"
)
```

[pim]: https://www.alexedwards.net/blog/an-introduction-to-packages-imports-and-modules
[hyrum]: https://www.hyrumslaw.com
[hyrum-xkcd]: https://xkcd.com/1172
[wiki-mod]: https://github.com/zchee/golang-wiki/blob/master/Modules.md#modules
[promise]: https://go.googlesource.com/website/+/1618ca05733f3f986996d9f0ad0f73b28362e317%5E%21/#F0
[mod-naming]: https://go.dev/blog/package-names
