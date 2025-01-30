#!/bin/bash

# this script does the following:
# - creates a directory called tetris
# - uses sqirvy-query and gemini-1.5-flash to create a design for a web app
# - uses sqirvy-query and claude-3-5-sonnet-latest to generate code for the design
# - uses sqirvy-review and gpt-4o-mini to review the code
# - starts a web server to serve the generated code
export BINDIR=../../bin

design="describe the steps required to build a web project that is a \
    simple web app that implements a simple tetris game clone.       \
    do not generate any code, just describe  what is needed to create the project. \
    code should be html, css and javascript, in a single file. \
    output will be markdown. your output will be to another LLM that will generate the code. "

code="generate code for the specified design. output only the code, no comments or annotations. \
      output will be html. do not wrap the output  in triple backticks"


rm -rf tetris && mkdir tetris 
echo $design | $BINDIR/sqirvy-query  -m gemini-1.5-flash                        >tetris/plan.md     && \
echo $code   | $BINDIR/sqirvy-query  -m claude-3-5-sonnet-latest tetris/plan.md >tetris/index.html  && \
               $BINDIR/sqirvy-review -m gpt-4o-mini tetris/index.html           >tetris/review.md   && \
python -m http.server 8080 --directory tetris
