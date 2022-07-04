#!/bin/bash

a="test"
# this will print ok
if [ "$a" = "test" ]; then echo ok; fi
# this will print ok too
if [ "$a" == "test" ]; then echo ok; fi
# what's the difference?
