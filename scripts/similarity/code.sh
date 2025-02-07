#!/bin/bash

# this script does the following:
# - creates a directory called diffs
# - uses sqirvy-code and claude-3-5-sonnet-latest to generate code as specified

prompt="create a simple web app that implements a simple tetris game clone.       \
    The game will include a game board with a grid, a score display, and a reset button. \
    Use html, css and javascript, in a single file. \
    The game should look modern and stylish. " 

export PATH=$PATH:../../bin
make -C ../../cmd

rm -rf code
mkdir  code

count=5
for ((i=0; i<count; i++))
do
    fname="code/$i.html"
    echo $fname
    echo $prompt | sqirvy -m gemini-2.0-flash-exp -f code  > $fname
done


