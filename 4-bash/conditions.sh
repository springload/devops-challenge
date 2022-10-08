#!/bin/bash
a="test"
# this will print ok
if [ "$a" = "test" ]; then echo ok; fi
# this will print ok too
if [ "$a" == "test" ]; then echo ok; fi
# what's the difference?

#Both run the same way on my machine(macos).Referring to the man pages for "test" it has the below entry
#For compatibility with some other implementations, the "=" primary can be substituted with "==" with the same meaning.
#however running 'sh conditions.sh' on my homelab (Ubuntu 22.04) I get the below result meaning "==" is not supported in sh, Where as BASH supports both.
#sh test.sh
#ok
#test.sh: 6: [: test: unexpected operator

#Side note it seems that zsh also does not support the use of "=="
#zsh conditions.sh
#ok
#conditions.sh:6: = not found