## bash challenges

Bash is tricky. It has weird syntax and has some little bash-specific things that don't
work in other shells. We work across multiple systems and it's important to write portable
scripts, which will work even if bash is not present on the system.

Please, take a look at 2 scripts in this folder.

1. fail.sh - you might run the script as `./fail.sh` and it will only print `pre`. But if you run it as `sh fail.sh` it goes past the fail point. Why? How to make it work either way?
2. conditions.sh - this script has two conditions, try to run it as `./conditions.sh`. Will it run? What happens if you run it as `sh conditions.sh`? Why?

### Answers

1. When the `-e` flag is appended to the shebang, it is instructing the script
   to exit on error. As there is a boolean `false` on line 5 of the script, the
   code below this line will never run. This a treated as a boolean literal (0).

   This is bash specific which is why it exits after printing `pre` for `./fail.sh`

   You could either remove the `-e` after the shebang on line 1, or comment out/
   remove `false` to print the whole output.

2. Yes looking at the script and after testing it out, it will run for both.
   Here the variable `a` is set the string `test` and then two conditions are
   passed. Both conditions check if `a` is equal to the string `"test"` however
   the second condition contains an operator specific to bash (`==`) so it is best
   not to use in other implementations of shell. The script will still work when
   run as `sh conditions.sh` as the shebang on line 1 is telling the OS which
   interpreter to use.
