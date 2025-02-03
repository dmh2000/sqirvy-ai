#!/bin/bash

models=(
	"claude-3-5-sonnet-latest"
	"claude-3-5-haiku-latest"
	"claude-3-opus-latest"
	"gemini-2.0-flash-exp"
	"gemini-1.5-flash"
	"gemini-1.5-pro"
	"gpt-4o"
	"gpt-4o-mini"
	"gpt-4-turbo"
	"o1-mini"
)


check_return_code() {
    echo ""
    local cmd="$1"
    $cmd $2 $3 $4 $5 $6 $7 $8 $9
    local return_code=$?
    
    if [ $return_code -eq 0 ]; then
        echo ""
        echo ""
        echo "$cmd $2 $3 executed successfully"
    else
        echo "$cmd $2 $3 failed with exit code $return_code"
        exit 1
    fi
    
    return $return_code
}



# use the latest build
# Build latest
export BINDIR=./bin
mkdir -p $BINDIR
make -s build

for item in "${models[@]}"
do
	echo "------ $item ------"
	echo "$BINDIR/sqirvy-query -m $item < hello.txt"
	check_return_code $BINDIR/sqirvy-query -m $item < hello.txt
done

rm -rf bin