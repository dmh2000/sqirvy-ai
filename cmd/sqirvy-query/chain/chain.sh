#!/bin/bash
query="../../build/sqirvy-query"

$query -m gemini-1.5-flash describe.md \
| $query -m claude-3-5-sonnet-latest generate.md \
| tee code.py
