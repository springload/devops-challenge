#!/bin/bash

a="test"
# this will print ok
if [ "$a" = "test" ]; then echo ok; fi
# this will print ok too
if [ "$a" == "test" ]; then echo ok; fi
# what's the difference?

# In bash the two are equivalent, and in sh = is the only one that will work.
# '==' is a bash-specific alias for '=' and it performs a string comparison 
