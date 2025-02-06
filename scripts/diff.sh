#!/bin/bash

# this script does the following:
# - creates a directory called diffs
# - uses sqirvy-code and claude-3-5-sonnet-latest to generate code as specified

prompt="create a simple web app that implements a simple tetris game clone.       \
    The game will include a game board with a grid, a score display, and a reset button. \
    Use html, css and javascript, in a single file named index.html. \
    The game should look modern and stylish. " 
    

export BINDIR=../bin  
make -C ../cmd

rm -rf diffs && mkdir diffs 
echo $prompt | $BINDIR/sqirvy -m claude-3-5-sonnet-latest -f code  >diffs/index.html

