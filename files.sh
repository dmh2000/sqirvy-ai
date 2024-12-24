#!/bin/sh
find . | egrep *.go$ > tmp
find . | egrep *.py$ >>tmp
cat tmp | sed 's/^..//' >files
cat files
rm tmp files

