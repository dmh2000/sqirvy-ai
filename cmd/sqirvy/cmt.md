# feature: Enhance README to describe additional command line programs and remove deprecated scripts
# details:
    - Expanded README to include detailed descriptions of command line programs: sqirvy-query, sqirvy-review, sqirvy-scrape, focusing on their functionalities and use cases.
    - Added an example section showcasing how to generate a Git commit comment using a shell command with sqirvy.
    - Removed the 'run-me.sh' script as it contained outdated implementation details.
    - Deleted multiple pipeline scripts (d1.sh, p1.sh, p2.sh, p3.sh) and related utility scripts (extract_code.py, files.sh) as they are no longer needed.
    - Added new functionality to support processing commit messages via a new 'commit' function.
    - Introduced a new prompt for generating commit messages in the application logic.
