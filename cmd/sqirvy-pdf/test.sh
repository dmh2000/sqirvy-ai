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

# Build latest
export BINDIR=./bin
mkdir -p $BINDIR
make -s build

# Run the tests
echo "sqiry-pdf -h"
check_return_code $BINDIR/sqirvy-pdf -h
echo "-------------------------------"
echo "$BINDIR/sqirvy-pdf < timetravel.pdf"
check_return_code $BINDIR/sqirvy-pdf <timetravel.pdf
echo "-------------------------------"
echo "$BINDIR/sqirvy-pdf timetravel.pdf"
check_return_code $BINDIR/sqirvy-pdf timetravel.pdf
echo "-------------------------------"
echo "$BINDIR/sqirvy-pdf README.pdf < timetravel.pdf"
check_return_code $BINDIR/sqirvy-pdf README.pdf < timetravel.pdf
echo "-------------------------------"
echo "$BINDIR/sqirvy-pdf  timetravel.pdf README.pdf"
check_return_code $BINDIR/sqirvy-pdf  timetravel.pdf README.pdf
# echo "-------------------------------"
echo "$BINDIR/sqirvy-pdf  timetravel.pdf README.pdf"
check_return_code $BINDIR/sqirvy-pdf -m gemini-2.0-flash-exp timetravel.pdf README.pdf
echo "-------------------------------"
echo "$BINDIR/sqirvy-pdf  timetravel.pdf README.pdf"
check_return_code $BINDIR/sqirvy-pdf -m gpt-4o timetravel.pdf README.pdf
echo "-------------------------------"
echo "$BINDIR/sqirvy-pdf xyz"
ignore_return_code $BINDIR/sqirvy-pdf xyz

rm -rf $BINDIR