# Advanced Testing with Go, Mitchell Hashimoto

Hashicorp uses Go as their primarily language, and their software faces
a number of challenges.  (NB: slides live [here][slides], and I've downloaded
them as well.)

+ Millions of people use their code, and a significant number of these people
  are enterprise users (with specific demands).
+ Many of their systems are distributed.
+ Many of their systems need to perform in extreme conditions.
+ Some of their systems have huge security implications.
+ Some of their systems must be correct, or they can cripple users.

Hashimoto divides his talk into two main areas: test methodology and writing
testable code.  (Dark slides cover methodology, and light slides cover writing
testable code.  The slides also move from basic to more complex and unusual.)

## Subtests (Test Methodology)

Subtests (new as of Go 1.8) have a number of advantages.

+ Subtests are closures. `t.Fatal` doesn't stop an entire run.
+ You can run a single test case from the commandline.
+ You can nest subtests infinitely.

## Table Driven Tests (Test Methodology)

These work well with subtests.  You pack a bunch of individual tests into
a slice of anonymous struct, and then run through the subtests with `range`.
Each individual struct is called within a `t.Run`, and the runs all have names
that can be clarifying and help readers of the test harness. Hashimoto
recommends giving subtests proper names rather than just naming them with data
or as indexes.

## Test Fixtures (Test Methodology)

+ `go test` always sets `pwd` as the package directory.  Thus, you can use
  relative paths for test fixtures.
+ Fixtures help test configuration, model data, binary data, and so on.

## Golden Files (Test Methodology)

+ A golden file has expected results, and you can test the result of your code
  to the golden file.
+ You can also provide an update mechanism that updates all golden files when
  something changes in the output.  (NB: you can easily add flags to testing
  via a test file.  TODO: look up how to do this.)
+ If the test fails, you should show diffs.

## Global State (Writing Testable Code)

+ Avoid global state as much as possible.
+ Wherever possible, use configuration instead.  Make the global state the
  default configuration, and allow tests to modify that default as needed.
+ Alternatively, make global state a variable.

## Test Helpers (Test Methodology)

+ Do not return an error from a test helper.  Instead, pass in `*testing.TB`
  and fail.
+ Make sure to use `t.Helper` (new as of Go 1.9).
+ You can use closures in a helper and return a function from the helper in
  ways that can be helpful.  (E.g., to make sure to close files or remove
  temporary items.)

## Repeat Yourself (Writing Testing Code)

In tests, you should repeat yourself (sometimes?).

+ Localized logic is more important than test lines of code.
+ When a test fails, you don't want to hunt down helpers in four other files.
  Put all the relevant logic right in the test itself even if this leads to
  repetition.
+ Limit helpers to repeat heavily used logic that does not often fail (e.g.,
  changing a directory) or something that fails all at once (e.g., creating
  a server or connecting to a database).
+ Note that helpers only help the person who (i) knows that they exist and
  (ii) knows what they do.
+ In the case of tests, they prefer a 200 line test to a 20 line test with
  abstracted helpers.

## Packages and Functions (Writing Testable Code)

Ideally, you want to split code into functions and packages when it makes
sense.  But it can be difficult to say exactly when it makes sense and when it
does not.

+ Only test exported functions (in general).
+ In general, don't test unexported functions or structs.
+ Hashimoto recommends against taking things too far.  He does sometimes test
  unexported functions and structs.
+ He likes internal packages as way to (smartly) "over-package."  You make it
  easier to refactor later.  Users cannot depend directly on things inside
  internal packages.  Also, he seems to say that if you use internal packages,
  you reduce the risk of having to update other things all at once.  Instead,
  you can replace the internal packages one at a time since each is imported
  separately.  (You cannot run into import cycles if you use internal
  packages wisely.  NB: I don't fully follow this.)

## Networking (Test Methodology)

+ If you are testing networking, make a real network connection.
+ Do not mock `net.Conn`.

## Configurability (Writing Testable Code)

+ Unconfigurable behavior makes it hard to write tests.
+ Over-parameterize structs initially to make them more testable.
+ If you use white-box tests, you can make these configurations unexported.
  That way, only tests can mess with them.
+ Sometimes he even adds a `Test` bool to a struct.  This allows you to turn
  certain things off during testing.

## Complex Structs (Test Methodology)

+ Mentions `reflect` and (more vaguely) 3rd party libraries.  He also mentions
  that `reflect` can fail in unhelpful ways when (e.g.) `reflect` says two
  things are not equal because one has an `int` and the other has an `int64`.
+ Sometimes, they use `testString` to stringify and compare strings rather
  than structs directly.  `testString` to be clear is something they write as
  needed.  The key idea seems to be that you can omit all sorts of unimportant
  details when you turn a struct into a string.  Equality libraries are
  (usually) all or nothing.  Note that `go-cmp` has ways of specifying that
  two things are not exactly equal, but they are close enough.

## Subprocessing (Test Methodology)

It is hard to test when there are subprocesses.  You usually have two choices.

1. You can actually do the subprocess.
1. You can mock the output or behavior of the subprocess.

They create a third option.  They execute a subprocess, but not necessarily
the actual (full) subprocess.

Also, when they do use mocks, they create executable mocks.  Thus, the code
still executes something but not necessarily the real thing.  They also like
to make the `*exec.Cmd` configurable.  (They took this technique from the
tests for the standard library in `os/exec`.)

## Interfaces (Writing Testable Code)

Like functions and packages, you can easily overdo interfaces.  That said,
there's again a lot of "You know when you know" involved here.

## Testing as a Public API (Writing Testing Code)

They export various testing files and APIs to allow people who use their code
to test things that use these packages.  They also provide things like
`slogtest` which confirm that other code upholds a contract.

## Custom Frameworks (Test Methodology)

In some cases, they write a big-ass custom framework to test complex things
that go way beyond unit tests.  But they still tie this material into `go
test`.  That is, they seem to still call it or run it starting with `go test`.

[slides]: https://speakerdeck.com/mitchellh/advanced-testing-with-go
