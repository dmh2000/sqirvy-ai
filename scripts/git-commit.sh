#!/bin/bash

# output the diff to stdout and pipe it into sqirvy -f commit then optionally accept it
make -C ../cmd
BINDIR=../bin
message=$(mktemp)
echo $message
git diff | $BINDIR/sqirvy . -f commit -m gemini-2.0-flash-exp | tee $message
read -a val -p "Do you want to commit this? [y/n] " -n 2 -r
if [[ $val == "y" ]]; then
    git commit -m "$(cat $message)"
fi
rm $message
