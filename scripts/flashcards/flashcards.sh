#!/bin/bash

# this script does the following:
# - creates a directory called react-app
# - uses sqirvy-query and gemini-1.5-flash to create a design for a web app
# - uses sqirvy-query and claude-3-5-sonnet-latest to generate code for the design
# - uses sqirvy-review and gpt-4o-mini to review the code
# - starts a web server to serve the generated code

plan="Design Specification: list the steps required to create a react app that allows users to create flashcards. \
    The app should have a homepage with a list of flashcards. \
    Users should be able to add new flashcards, edit existing flashcards, and delete flashcards. \
    Each flashcard should have a question and an answer. \
    The app should have a clean and modern design with a responsive layout.\
    The app should use React for the front-end and store the flashcards in memory."
code="Code Specification: create the react app according to the attached plan specification. \
    delimit the individual files you create as follows: '********<filename>' as suffix."


make -C ../../cmd
export BINDIR=$PWD/../../bin

export BUILD=build
rm -rf $BUILD && mkdir $BUILD

echo -n $plan         | $BINDIR/sqirvy -m gemini-2.0-flash-thinking-exp -f plan |\
tee $BUILD/plan.md  |\
cat - <(echo "$code") | $BINDIR/sqirvy -m claude-3-5-sonnet-latest -f code      |\
tee $BUILD/code.md |\
python extract_files.py $BUILD                                  

