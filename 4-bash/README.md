## bash challenges

Bash is tricky. It has weird syntax and has some little bash-specific things that don't
work in other shells. We work across multiple systems and it's important to write portable
scripts, which will work even if bash is not present on the system.

Please, take a look at 2 scripts in this folder.

### Questions 
1. fail.sh - you might run the script as `./fail.sh` and it will only print `pre`. But if you run it as `sh fail.sh` it goes past the fail point. Why? How to make it work either way?
2. conditions.sh - this script has two conditions, try to run it as `./conditions.sh`. Will it run? What happens if you run it as `sh conditions.sh`? Why?

### Answers
1. 
-e instructs bash to exit immediately if it receives a non zero status. e.g. in this case, `false` will return a non zero status. Or if you `ls` a directory that doesn't exist, that will also be a non zero status. 

Shebang is ignored if you explicitly specify shell. e.g. `sh fail.sh`.

To fix this, explicitly add `set -e` in the Bash Script instead of in the Shebang to make sure it works no matter how the script is executed. 

2.
Both single `=` and double `==` is used for String comparison. The difference between the two is compatability. Single `=` is supported by all shells, whereas double `==` is not supported by some of the older shells. So it is advised to use single `=`. 