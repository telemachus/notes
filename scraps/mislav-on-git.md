# [git.md](https://gist.github.com/mislav/5147811)

What I believe is the most confusing thing for git beginners is that the main
commands have multiple uses depending on the arguments or context, and some of
those uses aren't well-described by the command's name.

## Merge

- **merges** together different branches (different _histories_, to be more exact)
- **fast-forwards** if a merge isn't necessary; effectively repositioning the
  branch head to a descendant commit
- [unwanted outcome](http://mislav.uniqpath.com/2013/02/merge-vs-rebase/) of
  `git pull`

## Rebase

- **replays** a set of commits as if they were based off another parent
- **interactively** used for history cleanup: reorder & squash related commits
- should be used for `git pull` instead of merge

## Reset

- **resets** (unstages) changes added to the index; undoes `git add`
- **repositions** the branch head to another commit; optionally resetting index
  & working tree
  - **undo a commit**, either discarding changes or keeping them for the next commit
  - **discard** all working copy changes since the last commit

## Checkout

- **switch** branches
  - **create** a new branch and switch to it
- **checkout** a file or directory as it were in another commit
  - **discard** all working copy changes since the last commit
  - **resolve** merge conflicts
