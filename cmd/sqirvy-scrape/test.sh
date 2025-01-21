#!/bin/bash

check_return_code() {
    local cmd="$1"
    $cmd $2 $3 $4 $5 $6 $7 $8 $9
    local return_code=$?
    
    if [ $return_code -eq 0 ]; then
        echo "Command '$cmd' executed successfully"
    else
        echo "Command '$cmd' failed with exit code $return_code"
        exit 1
    fi
    
    return $return_code
}

ignore_return_code() {
    local cmd="$1"
    $cmd $2 $3 $4 $5 $6 $7 $8 $9
    local return_code=$?
    
    return 0
}

makecopy="scrape this url and create a single html file containing html,css and js that \
   creates a dummy webpage that has the same layout and styling as the original webpage. \
   do not include any explanations or other text in the output. remove any triple backticks from the output.  \
   the output should be ready to be served as a webpage"


# Build latest
export BINDIR=./bin
mkdir -p $BINDIR
make -s build

echo "-------------------------------"
echo "$BINDIR/sqirvy-scrape (should fail)"
ignore_return_code go run .
echo "-------------------------------"
echo "$BINDIR/sqirvy-scrape -h"
check_return_code go run . -h
echo "-------------------------------"
echo "$BINDIR/sqirvy-scrape https://sqirvy.xyz"
check_return_code echo "summarize the url" | go run . https://sqirvy.xyz
echo "-------------------------------"
echo "$BINDIR/sqirvy-scrape https://sqirvy.xyz https://test-alert.vercel.app/"
check_return_code echo "summarize the urls" | go run . https://sqirvy.xyz https://test-alert.vercel.app/
echo "-------------------------------"
echo "'copy' $BINDIR/sqirvy-scrape https://sqirvy.xyz "
check_return_code echo $makecopy | go run . https://sqirvy.xyz  xyz.html
echo "-------------------------------"

rm -rf bin