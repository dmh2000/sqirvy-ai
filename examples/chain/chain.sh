#!/bin/bash
query="../../bin/sqirvy-query"

# this example uses the system.md prompt by default and 
# then uses a chain of prompts to generate the code.py file
# system.md   : a system prompt for software engineers
#             : this is the default system prompt used by sqirvy-query if it is in the local directory
# describe.md : a general description of well formed python code
#             : uses gemini-1.5-flash model for this query
# generate.md : a description of the specific code to generate
#             : uses claude-3-5-sonnet-latest model for this query


# create the prompt files then pipe them to the queries
$query -m gemini-1.5-flash describe.md |\
$query -m claude-3-5-sonnet-latest generate.md |\
tee code.py 
