#!/bin/bash

# a test must pass
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

mkdir -p tmp
echo "-------------------------------"
echo "sqirvy no flags or args"
check_return_code go run .                                                            >tmp/no-flags-or-args.md
echo "-------------------------------"
echo "sqirvy -h"
check_return_code go run . -h                                                         >tmp/help.md
echo "-------------------------------"
echo "sqirvy https://sqirvy.xyz"
check_return_code echo $scrape  | go run . -f scrape https://sqirvy.xyz                >tmp/scrape.html
echo "-------------------------------"
echo "sqirvy -f code"
check_return_code echo $code |    go run . -f code                                      >tmp/code.html
echo "-------------------------------"
echo "sqirvy -f review"
check_return_code                 go run . -m gemini-2.0-flash-exp -f review main.go     >tmp/review.md
echo "-------------------------------"
echo "sqirvy -f query"
check_return_code echo $query |   go run . -m gpt-4-turbo          -f query main.go      >tmp/query1.md
echo "-------------------------------"
echo "sqirvy -f query (default if no -f)"
check_return_code echo $query |   go run . -m meta-llama/meta-llama-3.1-8b-instruct-turbo >tmp/query2.md
echo "-------------------------------"
