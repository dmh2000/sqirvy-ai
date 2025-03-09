#!/bin/bash

# this script does the following:
# - creates a directory called tetris
# - uses sqirvy-query and gemini-1.5-flash to create a design for a web app
# - uses sqirvy-query and claude-3-5-sonnet-latest to generate code for the design
# - uses sqirvy-review and gpt-4o-mini to review the code
# - starts a web server to serve the generated code

design="create a design specification for a web project that is a \
    simple web app that implements a simple tetris game clone.       \
    the game should include a game board with a grid, a score display, and a reset button \
    Code should be html, css and javascript, in a single file named index.html. \
    Output will be markdown.  "

export BINDIR=../bin  
make -C ../cmd

rm -rf tetris && mkdir tetris 
echo $design | \
$BINDIR/sqirvy-cli plan -m gemini-1.5-flash            | tee tetris/plan.md    | \
$BINDIR/sqirvy-cli code -m claude-3-5-sonnet-latest    | tee tetris/index.html | \
$BINDIR/sqirvy-cli review -m gpt-4o-mini                     review  >tetris/review.md   

python -m http.server 8080 --directory tetris &

xdg-open http://localhost:8080
