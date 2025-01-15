#!/bin/bash
query="../../bin/sqirvy-query"

# this example uses the system.md prompt by default 
prompt="review the following code for \
    bugs, style, clarity, security, performance and idiomatic code. \
    include filename and line number of any findings. \
    output in markdown format.  \
    code is well-designed. \
    code isn’t more complex than it needs to be. \
    no extraneous or unused code is present.\
    naming conventions are consistent and descriptive. \
    code is formatted properly. \
    Comments are clear and useful. \
    The code conforms to the appropriate style guide."

# use the smaller models for the review
echo $prompt | $query  -m claude-3-5-haiku-latest ../../cmd/sqirvy-query/*.go > review-anthropic.md
echo $prompt | $query  -m gemini-1.5-flash ../../cmd/sqirvy-query/*.go > review-gemini.md
echo $prompt | $query  -m gpt-4o-mini ../../cmd/sqirvy-query/*.go > review-openai.md

