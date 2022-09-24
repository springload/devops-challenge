# 1

## RCA
`sh fail.sh` executes the script using the Bourne shell which then ignores the shebang at the top of the script. This means the `-e` flag is not set and the script won't exit on error.

## Solution
Removing the `-e` flag from the shebang and setting it explicitly in the script itself means it will be set for any shell running the script that respects this flag.

# 2

## RCA
`sh conditions.sh` actually runs correctly for me locally, however some shells may have issues with the `==` operator, `=` is the standard POSIX comparison operator.

## Solution
Use `=` operator for better cross-platform compatibility.