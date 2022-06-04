## git challenge

In this repository you can see a branch named `feature/git`.
If you look at the commits, it's a total mess.

1. There are too many commits with non-descriptive messages
2. Would you be able to tidy it up? For example, make a couple of commits
   with good commit messages out of those.
3. Specifically, it would be nice to have 2 commits instead of 4.

### Answers

- Check out the log with `git log -p`
- Create new commits on my branch to tidy up as an example.
- Note: another alternative could be to check out `feature/git` branch
  and run `git revert --no-commit` against all four SHA's.
  I chose not to do this as it would make the history quite ugly, and
  instead opted to make the changes and example commits against my own
  branch.
