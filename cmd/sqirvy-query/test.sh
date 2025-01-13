#!/bin/sh
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

make build 
echo "-------------------------------"
echo "sqiry-query"
ignore_return_code ../build/sqirvy-query
echo "-------------------------------"
echo "sqiry-query -h"
check_return_code ../build/sqirvy-query -h
echo "-------------------------------"
echo "../build/sqirvy-query < hello.txt"
check_return_code ../build/sqirvy-query < hello.txt
echo "sqiry-query file"
echo "-------------------------------"
echo "../build/sqirvy-query hello.txt"
check_return_code ../build/sqirvy-query hello.txt
echo "-------------------------------"
echo "../build/sqirvy-query goodbye.txt < hello.txt"
check_return_code ../build/sqirvy-query goodbye.txt < hello.txt
echo "-------------------------------"
echo "../build/sqirvy-query  hello.txt goodbye.txt"
check_return_code ../build/sqirvy-query  hello.txt goodbye.txt
echo "-------------------------------"
echo "../build/sqirvy-query  -m gemini-2.0-flash-exp < hello.txt goodbye.txt"
check_return_code ../build/sqirvy-query  -m gemini-2.0-flash-exp < hello.txt goodbye.txt
echo "-------------------------------"
echo "../build/sqirvy-query  -m gpt-4o-2024-11-20 < hello.txt goodbye.txt"
check_return_code ../build/sqirvy-query  -m gpt-4o-2024-11-20 < hello.txt goodbye.txt
echo "-------------------------------"
echo "../build/sqirvy-query  -m gpt-4o-2024-11-20 < hello.txt goodbye.txt"
ignore_return_code ../build/sqirvy-query  -m xyz < hello.txt 
echo "-------------------------------"
echo "../build/sqirvy-query  -m gpt-4o-2024-11-20 < hello.txt goodbye.txt"
ignore_return_code ../build/sqirvy-query  xyz
echo "-------------------------------"
