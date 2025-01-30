#!/bin/bash

export BINDIR=../../bin

design="describe the steps required to build a web project that uses a node.js server and a \
simple web app that creates a simple tetris game clone. do not generate any code, just describe \
what is needed to create the project. the project should be able to run locally. \
code should be html, css and javascript, in separate files. \
include instructions for adding an http server for the project. \
output will be a markdown. your output will be to another LLM that will generate the code."

code="generate code for the specified design. there will be multiple code files generated. \
output will be a single file containing the code for all file. wrap each separate content in \
 triple backticks,  where language is the file type of each separate code file. \
 output will be markdown. do not stop until all the code is generated."

review="treat each section of content that is wrapped in triple backticks as a code file. review each file for \
style, correctness, security and best practices. output will be a markdown file with the review"
 

rm -rf tetris && mkdir tetris && \
echo $design  | $BINDIR/sqirvy-query -m gemini-1.5-flash                         >tetris/plan.md && \
echo $code    | $BINDIR/sqirvy-query -m claude-3-5-sonnet-latest tetris/plan.md  >tetris/code.md && \
echo $review  | $BINDIR/sqirvy-query -m gpt-4o-mini              tetris/code.md  >tetris/review.md 

