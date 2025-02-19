#!/bin/bash

BINDIR=../../bin
TARGET=./sqirvy-cli
TESTDIR=./test

# rebuild the binaries
make

# a test must pass
check_return_code() {
    local cmd="$1"
    $cmd $2 $3 $4 $5 $6 $7 $8 $9
    local return_code=$?
    
    if [ $return_code -ne 0 ]; then
        echo "Command '$cmd' failed with exit code $return_code"
        exit 1
    fi
    
    return $return_code
}

# ok if a test fails
ignore_return_code() {
    local cmd="$1"
    $cmd $2 $3 $4 $5 $6 $7 $8 $9
    local return_code=$?
    
    return 0
}

scrape="scrape this url and create a single html file containing html,css and js that \
   creates a dummy webpage that has the same layout and styling as the original webpage. \
   do not include any explanations or other text in the output. remove any triple backticks from the output.  \
   the output should be ready to be served as a webpage"

plan="create a plan for scafolding a single page web app using the vit framework. \
   the app should be sleek and modern. \
   the app should have a responsive layout that works on desktop and mobile. \
   the app should have a dark theme. \
   the app should have a sidebar that can be toggled on and off. 
   the functionality of the app will be determined later.
   the app should be built using the latest version of vit. \
   the app should be built using the latest version of typescript.
   this is a plan only, do not generate any code."

code="create a simple webpage with a counter and buttons to increment and decrement the counter. \
   the counter should be stored in a cookie so that it persists across page reloads. \
   the counter should be initialized to 0 when the page is first loaded. \
   the counter should be incremented by 1 when the increment button is clicked. \
   the counter should be decremented by 1 when the decrement button is clicked. \
   the counter should never be less than 0. \
   the counter should be displayed in the center of the page. \
   the increment and decrement buttons should be displayed below the counter. \
   the increment and decrement buttons should be centered horizontally. \
   the increment and decrement buttons should be styled so that they are visually distinct. \
   use html, css and javascript in a single file"

query="what is the sum of 1 + 2 + 3"   

mkdir -p $TESTDIR
make -s debug

echo "-------------------------------"
echo "sqirvy no flags or args"
check_return_code                 $TARGET                                             >$TESTDIR/no-flags-or-args.md
echo "-------------------------------"
echo "sqirvy -h"
check_return_code                 $TARGET -h                                          >$TESTDIR/help.md
echo "-------------------------------"
echo "sqirvy  plan"
check_return_code echo $plan |    $TARGET plan   -m gemini-1.5-flash main.go          >$TESTDIR/plan.md
echo "-------------------------------"
echo "sqirvy  code"
check_return_code echo $code |    $TARGET code                                        >$TESTDIR/code.html
echo "-------------------------------"
echo "sqirvy review"
check_return_code                 $TARGET review -m gemini-1.5-flash main.go          >$TESTDIR/review.md
echo "-------------------------------"
echo "sqirvy query"
check_return_code echo $query |   $TARGET query -m gpt-4-turbo  main.go               >$TESTDIR/query1.md
echo "-------------------------------"
echo "sqirvy "
check_return_code echo $query |   $TARGET query -m llama3.3-70b                       >$TESTDIR/query2.md
echo "-------------------------------"
