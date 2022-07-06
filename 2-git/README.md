## git challenge

In this repository you can see a branch named `feature/git`.
If you look at the commits, it's a total mess.

1. There are too many commits with non-descriptive messages
2. Would you be able to tidy it up? For example, make a couple of commits
   with good commit messages out of those.
3. Specifically, it would be nice to have 2 commits instead of 4.

## Answers

1. We can do it all via `git rebase`. First at all we need to check commit history via `git log`, after that we need type `git rebase -i HEAD~4` where we specify nubmer of commits. After submit we will see a first window where we can `pick` first commit and `squash` second one. The same things we should do with third and fourth commits. On a second and third windows we need to write right commit message and push it with flag `--force`.
