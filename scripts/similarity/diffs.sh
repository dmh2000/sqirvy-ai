#!/bin/bash

# this script does the following:
# - creates a directory called diffs
# - uses 'diff' to compare code files

export PATH=$PATH:../../bin
export CODE=./code

# remove old diffs
rm -rf diffs
mkdir  diffs

# create the new diffs
count=$(ls -1 $CODE | wc -l)
touch diffs/diffs.csv
for ((i=0; i<count; i++))
do
    for ((j=0; j<count; j++))
    do
        echo -n "$i, $j ," >> diffs/diffs.csv
        dname="diffs/$i-$j.diff"
        echo $dname
        diff  code/$i.html code/$j.html > $dname
        words=$(wc -w diffs/$i-$j.diff | awk '{printf "%s", $1}')
        echo "$words" >> diffs/diffs.csv
    done
done    
