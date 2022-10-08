## git challenge

In this repository you can see a branch named `feature/git`.
If you look at the commits, it's a total mess.

1. There are too many commits with non-descriptive messages
2. Would you be able to tidy it up? For example, make a couple of commits
   with good commit messages out of those.
3. Specifically, it would be nice to have 2 commits instead of 4.

I have made a seperate PR for this challenge. I used "git rebase -i HEAD~4" and sqush 4 commits into 2 with new messages.