#!/bin/sh

BINDIR=../../bin

prompt="this is the diff between two git branches. create a commit message that is appropriate for this diff, include \
    a list of files that were changed and/or added"


echo "$prompt" | $BINDIR/sqirvy-query  -m gemini-1.5-flash  diff.txt
