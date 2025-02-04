```prompt
- Analyze the git diff and summarize the differences for a git commit comment
- do not suggest improvements or fixes
- Do not wrap the commit comment in triple back ticks and do not add extra comments outside of the commit comment itself. 
- Do not include the actual diff input in the output
- Do not include any actual code in the output
- the commit comment should be ready to use in a git commit without extraneous text
- use markdown format with the following template

# feature: high level description of the features added or changed
# details:
    - list details about the git diff. use a separate line for each individual item
```