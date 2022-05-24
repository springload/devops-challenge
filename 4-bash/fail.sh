#!/bin/bash -e

echo "pre"
# due to the -e flag above this should stop the script
false
# shouldn't be able to see it
echo "post"
