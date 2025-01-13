#!/bin/bash
query="../../build/sqirvy-query"

$query -m gemini-1.5-flash system.md describe.md \
| $query -m claude-3-5-sonnet-latest system.md generate.md \

