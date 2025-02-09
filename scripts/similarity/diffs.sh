#!/bin/bash

# this script does the following:
# - creates a directory called diffs
# - uses 'diff' to compare code files

if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <code-directory>"
    exit 1
fi


export PATH=$PATH:../../bin
export CODE=$1
export DIFFS=$1-diffs
export CSV=$DIFFS/diffs.csv

# remove old diffs
rm -rf $DIFFS
mkdir  $DIFFS

# create the new diffs
count=$(ls -1 $CODE | wc -l)
touch $CSV

for ((i=0; i<count; i++))
do
    for ((j=0; j<count; j++))
    do
        echo -n "$i, $j ," >> $CSV
        dname="$DIFFS/$i-$j.diff"
        diff  $CODE/$i.html $CODE/$j.html > $dname
        words=$(wc -w $DIFFS/$i-$j.diff | awk '{printf "%s", $1}')
        echo "$words" >> $CSV
    done
done    
