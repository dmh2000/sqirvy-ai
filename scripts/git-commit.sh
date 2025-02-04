
# output the diff to stdout and pipe it into sqirvy -f commit then optionally accept it
make -C ../cmd
BINDIR=../bin
message=$(mktemp)
echo $message
git diff | $BINDIR/sqirvy . -f commit -m gemini-2.0-flash-exp | tee $message
read -p "Do you want to commit this? [y/n] " -n 1 -r
git commit -m "$(cat $message)"
rm $message
