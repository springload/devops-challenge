#!/bin/bash

set -e #the flag should be explicitly set instead of being in the shebang line at the start
#this is because when the script is run with "sh fail.sh" it uses the Bourne Shell 
#which will ignore that first shebang line hence it will ignore the -e

echo "pre"
# due to the -e flag above this should stop the script
false
# shouldn't be able to see it
echo "post"
