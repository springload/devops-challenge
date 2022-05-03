## git challenge

In this repository you can see a branch named `feature/git`.
If you look at the commits, it's a total mess.

1. There are too many commits with non-descriptive messages
2. Would you be able to tidy it up? For example, make a couple of commits
   with good commit messages out of those.
3. Specifically, it would be nice to have 2 commits instead of 4.

## My notes

Below would be my approach on this challenge.  
There is probably a more elegant way than mine out there to handle this situation.

### Steps
- Checkout the `feature/git` branch.
```
git checkout feature/git
```

- Using `git log` to list all the commits in the branch.
```

commit d6eff12db422dba1c973e875a3bce3e34bdf15da (HEAD -> feature/git, origin/feature/git)
Author: Eugene Dementyev <eugene@springload.co.nz>
Date:   Thu Apr 15 10:37:05 2021 +1200

    test 2

commit 0b102d05dc5c2af14773ddfb40700ff01ede560f
Author: Eugene Dementyev <eugene@springload.co.nz>
Date:   Thu Apr 15 10:36:36 2021 +1200

    test 1

commit ff4d72dcd869ac23083d6525631e6e274ba20f37
Author: Eugene Dementyev <eugene@springload.co.nz>
Date:   Thu Apr 15 10:35:20 2021 +1200

    update 2

commit c2ed588b04ce48f200a185b9a1d8096fa3fb215e
Author: Eugene Dementyev <eugene@springload.co.nz>
Date:   Thu Apr 15 10:35:08 2021 +1200

    update 1

commit 203f47e358d9491e8e8c567266e9ded32ee8476c
Author: Eugene Dementyev <eugene@springload.co.nz>
Date:   Thu Apr 15 10:34:25 2021 +1200

    Add git challenge

commit 96b7500a9c37cc6d77325a8df03790c4db2e2fdf
Author: Eugene Dementyev <eugene@springload.co.nz>
Date:   Thu Apr 15 10:01:20 2021 +1200

    Add Docker challenge

```

- Create two patch files based on the last 4 commits.
```
git diff 203f47e358d9491e8e8c567266e9ded32ee8476c..ff4d72dcd869ac23083d6525631e6e274ba20f37 > /tmp/patch1
git diff ff4d72dcd869ac23083d6525631e6e274ba20f37..d6eff12db422dba1c973e875a3bce3e34bdf15da > /tmp/patch2
```

- Reset the branch to remove the last 4 commits.  Also clean out the untracked changes
```
git reset HEAD~4

# Use 'rm' because the files were completely untracked...
rm mega_feature_*.txt

# Use 'git checkout' instead if the files were partially modified...
git checkout mega_feature_*.txt
```

- Reapply the changes using the patch files
```
git apply /tmp/patch1
git add mega_feature_1.txt 
git commit -m "Created mega_feature_1.txt"

git apply /tmp/patch2
git add mega_feature_2.txt 
git commit -m "Created mega_feature_2.txt"
```

- Forcefully, push the commits back into the remote `feature/git` branch.
```
git push -f origin feature/git
```

