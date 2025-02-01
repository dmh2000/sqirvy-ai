#!/bin/bash

# this script does the following:
# - creates a directory called tetris
# - uses sqirvy-query and gemini-1.5-flash to create a design for a web app
# - uses sqirvy-query and claude-3-5-sonnet-latest to generate code for the design
# - uses sqirvy-review and gpt-4o-mini to review the code
# - starts a web server to serve the generated code
export BINDIR=../sqirvy-llm/bin

design_prompt="describe the steps required to build a                       \
    simple web app that implements a simple tetris game clone.       \
    do not generate any code, just describe  what is needed to create the project. \
    code should be html, css and javascript, in a single file.       \
    output will be markdown. your output will be to another LLM that will generate the code. "

code_prompt="generate code for the specified design. output only the code, no comments or annotations. \
      output will be html. do not wrap the output  in triple backticks"

append() {
    input=$(</dev/stdin)
    echo $input
    echo " "
    echo "$1"
}

rm -rf tetris && mkdir tetris 

echo "$design_prompt"                                                     | \
$BINDIR/sqirvy-query  -m gemini-1.5-flash         | tee tetris/design.md  | \
append "$code_prompt"                                                     | \
$BINDIR/sqirvy-query  -m claude-3-5-sonnet-latest | tee tetris/index.html | \
$BINDIR/sqirvy-review -m gpt-4o-mini              > tetris/review.md

python -m http.server 8080 --directory tetris
