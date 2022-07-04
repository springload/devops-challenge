## bash challenges

Bash is tricky. It has weird syntax and has some little bash-specific things that don't
work in other shells. We work across multiple systems and it's important to write portable
scripts, which will work even if bash is not present on the system.

Please, take a look at 2 scripts in this folder.

1. fail.sh - you might run the script as `./fail.sh` and it will only print `pre`. But if you run it as `sh fail.sh` it goes past the fail point. Why? How to make it work either way?
2. conditions.sh - this script has two conditions, try to run it as `./conditions.sh`. Will it run? What happens if you run it as `sh conditions.sh`? Why?
