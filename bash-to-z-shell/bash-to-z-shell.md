# *From Bash to Z Shell: Conquering the Command Line*, Kiddle, Peek, Stephenson

## `for` and `foreach` Loops

Bash-like shells use `for` to create loops; C-type shells (e.g., `tcsh`) use `foreach`.
I will focus on Bash-like shells

```bash
for f in *
do
    cp -i "{f}" "OLD-${f}"
done
```

Since you can use `;` as a statement separator in shell scripts, you can put
this all on one line.

```bash
for f in *; do; cp -i "{f}" "OLD-${f}"; done
```

That is useful during interactive sessions, but don't do it in scripts.
On the other hand, many programmers prefer to keep the initial `do` on the first
line.
(I see this as parallel to placing an initial opening brace on the same line.)

```bash
for f in *; do
    cp -i "{f}" "OLD-${f}"
done
```

You can redirect the results of a loop to a file or pass the results to a pipe as if it were a single command.

```bash
for f in *; do
    ls -l "{f}"
done > data.txt

for f in *; do
    ls -l "{f}"
done | grep -v 'txt$'
```

The environment variables that belong to a parent process are copied to every child process of the parent.
The child can do what it likes with those variables, and it can pass new versions as it sees fit.
However, a child process cannot change the environment variables of its parent.

Bash (and most Bourne and C shells) treat spaces and newlines as introducing new arguments.
Z shell, however, does not split expanded variables or command substitution at spaces and newlines.
Thus, if you have a multi-line address in `$addr`, you must quote it in Bourne and C shells, but you can use it without quotes in Z shell.
(Actually, C shells require more than just quoting.
They require `${addr:q}`.
Happily, I do not care too much about C shells.)
