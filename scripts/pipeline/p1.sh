#!/bin/bash

# this script does the following:
# - creates a directory called tetris
# - uses sqirvy-query and gemini-1.5-flash to create a design for a web app
# - uses sqirvy-query and claude-3-5-sonnet-latest to generate code for the design
# - uses sqirvy-review and gpt-4o-mini to review the code
# - starts a web server to serve the generated code
export BINDIR=../../bin

design="create a design specification for a web project that is a \
    simple web app that implements a simple tetris game clone.       \
    do not generate any code, just describe  what is needed to create the project. \
    code should be html, css and javascript, in a single file named index.html \
    output will be markdown. your output will be to another LLM that will generate the code. "

rm -rf tetris && mkdir tetris 
echo $design | \
$BINDIR/sqirvy-query  -m gemini-1.5-flash         | tee tetris/plan.md    | \
$BINDIR/sqirvy-code   -m claude-3-5-sonnet-latest | tee tetris/index.html | \
$BINDIR/sqirvy-review -m gpt-4o-mini              >tetris/review.md   

python -m http.server 8080 --directory tetris

# test