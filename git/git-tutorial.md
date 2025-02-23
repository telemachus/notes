# [gittutorial][gittutorial]

## Getting Help

The command `man git-x` is equivalent to `git help x`.  E.g., `man git-log` and
`git help log`.

[gittutorial]: https://git-scm.com/docs/gittutorial

## `git add`

You use `git add` for two things.  Or, rather, you use the command for one
thing, but it can seem like two things to new users.  When you run `git add`,
`git` "takes a snapshot of the given files and stages that content in the
index, ready for the next commit."  You use `git add` both for new and changed
files.  New users may find it odd to call `git add` on a file that isn't new,
but to `git`, you are adding the new content to the index.

## `git log` Tips

+ `git log -p` shows complete diffs
+ `git log --stat --summary` shows more than `git log`, but less than `git log
  -p`

## `git pull` and `git fetch`

`git fetch` draws in changes from remote, while `git pull` both draws in
changes from a remote and merges those changes.  `git fetch` is safer since
you can decide how to merge more explicitly afterwards.  `git` provides
a special alias, `FETCH_HEAD`, that you can use to compare fetched material
and a repo's current state.  E.g., `git log -p HEAD..FETCH_HEAD`.

## Two Range Notations in `git`

Two dots in a range means "show everything reachable from the right-hand side
but exclude everything reachable from the left-hand side.  Three dots in
a range mans "show everything that is reachable from either side, but exclude
everything that is reachable from both of them."

So, for example, `git log -p HEAD..FETCH_HEAD` shows everything that
`FETCH_HEAD` has done that `HEAD` hasn't see.  On the other hand, `git log -p
HEAD...FETCH_HEAD` shows everything that both `HEAD` and `FETCH_HEAD` have
done since the two forked.

When you use a range in `git`, the right-hand side does not need to be an
ancestor of the left-hand side.  Imagine that `main` and `stable` diverged in
the past from a common commit.  In that case, `git log stable..main` shows
commits made on the main branch but not the stable branch.  And `git log
main stable` shows commits made on the stable branch but not the main branch.

## `git show`

You can use `git show` with a SHA1 or part of a SHA1 to see details about
a commit. You can also use built-in aliases.  For example, `git show HEAD`.
You can also use branch names.

`git` also provides a way to count.  `git show HEAD^` shows the parent of
`HEAD`, and `git show HEAD~4` shows the fourth commit before `HEAD`.  In many
contexts, `^` and `~` mean the same thing.  But if a commit has two or more
parents, their meaning is different.  In that case, `item^n` means the `n-th`
parent of item, but `item~n` means the `n-th` generation ancestor of item
*following only first parents*.  This can get confusing quickly.  For more
information, see the ["Specifying Revisions" section of the `git rev-parse`
documentation][sr].

## `git grep`

Yuo can use `git grep` to search for strings in a particular version of
a repository or in the entire history of the repository.

[sr]: https://git-scm.com/docs/git-rev-parse#_specifying_revisions

## Specifying Files at Specific Commits

In general, if a command takes a filename, you can add an optional
specification for a particular version of the file.  E.g., `git diff
v2.4:Makefile HEAD:Makefile` or `git show v2.4:Makefile`.
