# Learning Go: Chapter 1

## Staying Up to Date

You can install go releases as if they were a local module. If you install
this way, rather than globally, you can test the new version without losing
your previous installation. Imagine that you have go 1.15.2 installed, and you
want to test out 1.15.6. Here’s how you can do that safely.

```shell
$ go get golang.org/dl/go.1.15.6
$ go1.15.6 download
# … later
$ go1.15.6 build # or other commands
# … later, if 1.15.6 works
$ go1.15.6 env GOROOT
$ rm -rf $(go1.15.6 env GOROOT)
$ rm $(go env GOPATH)/bin/go1.15.6
# then install 1.15.6 globally
```
