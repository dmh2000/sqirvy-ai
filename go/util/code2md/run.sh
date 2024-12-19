#!/bin/bash

echo "CLI"
python main.py <../../cmd/cli/main.go >../../cmd/cli/cli.md
echo "WEB"
python main.py <../../web/main.go >../../web/web.md
echo "PWD"
python main.py <../../internal/pwd/pwd.go >../../internal/pwd/pwd.md