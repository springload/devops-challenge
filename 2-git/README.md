## git challenge

### Question 
In this repository you can see a branch named `feature/git`.
If you look at the commits, it's a total mess.

1. There are too many commits with non-descriptive messages
2. Would you be able to tidy it up? For example, make a couple of commits
   with good commit messages out of those.
3. Specifically, it would be nice to have 2 commits instead of 4.

### Answer 
To fix the commit history, I did the following:
1. `git rebase -i HEAD~4`
2. squash the 2 commits (update 2 and test 2)
3. rename the 2 commits appropriately
4. `git log` to check the commits are correct
5. `git push --force` to push rebased branch to origin branch