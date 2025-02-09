#!/bin/sh

./diffs.sh gemini
python main.py gemini-2.0-flash-exp gemini-diffs/diffs.csv gemini

# ./diff.sh anthropic-diffs
# python main.py claude-3-5-sonnet-latest diffs/anthropic-diffs.csv anthropic
# python main.py o1_mini  diffs/diffs.csv openai