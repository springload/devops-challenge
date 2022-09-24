# 1

## RCA
`sh fail.sh` executes the script using the Bourne shell which then ignores the shebang at the top of the script. This means the `-e` flag is not set and the script won't exit on error.

## Solution
Removing the `-e` flag from the shebang and setting it explicitly in the script itself means it will be set for any shell running the script that respects this flag.

